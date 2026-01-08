package auth

import (
	"context"
	"fmt"
	"math/rand/v2"
	"meawby/internal/config"
	"meawby/internal/model/modelAuth"
	"meawby/internal/model/modelErrors"
	"meawby/internal/model/modelUser"

	"meawby/internal/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthService struct {
	account      repository.AccountRepository
	smsCode      repository.SmsCodeRepository
	refreshToken repository.RefreshToken
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{account: repo.AccountRepository, smsCode: repo.SmsCodeRepository, refreshToken: repo.RefreshToken}
}

var cfg = &config.AuthConfig{
	JWTsecret:    "kjwefiuwv",
	AccessToken:  15 * time.Minute,
	RefreshToken: 7 * 24 * time.Hour,
}

func (s *AuthService) SendSMS(ctx context.Context, phone string) error {
	account, err := s.account.GetByPhone(ctx, phone)
	if err != nil {
		return err
	}
	if account == nil {
		account = &modelUser.Account{
			ID:    uuid.New(),
			Phone: phone,
		}
		if err := s.account.Create(ctx, account); err != nil {
			return err
		}
	}
	code := &modelAuth.SMSCode{
		ID:        uuid.New(),
		AccountID: account.ID,
		Code:      generateCode(),
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	return s.smsCode.Create(ctx, code)
}

func (s *AuthService) VerifySMS(ctx context.Context, phone, code string) (*modelAuth.VerifySMSResponse, error) {
	account, err := s.account.GetByPhone(ctx, phone)
	if err != nil || account == nil {
		return nil, modelErrors.ErrAccountNotFound
	}
	sms, err := s.smsCode.GetValidCode(ctx, account.ID, code)
	if err != nil {
		return nil, modelErrors.ErrSmsCode
	}
	if err := s.smsCode.Delete(ctx, sms.ID); err != nil {
		return nil, err
	}
	accessToken, tt := generateJWT(account.ID, account.Role, cfg.JWTsecret, cfg.AccessToken)
	refreshToken, _ := generateRefresh(cfg.JWTsecret, cfg.RefreshToken)
	refresh := &modelAuth.RefreshToken{
		ID:        uuid.New(),
		AccountID: account.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(cfg.RefreshToken),
	}
	if err := s.refreshToken.Create(ctx, refresh); err != nil {
		return nil, err
	}
	return &modelAuth.VerifySMSResponse{
		AccessToken:  accessToken,
		RefreshToken: refresh.Token,
		ExpiresAt:    tt,
	}, nil

}

func (s *AuthService) Refresh(ctx context.Context, tokenStr string) (*modelAuth.RefreshTokenResponse, error) {
	rt, err := s.refreshToken.GetByToken(ctx, tokenStr)
	if err != nil {
		return nil, modelErrors.ErrRefreshToken
	}
	if time.Now().After(rt.ExpiresAt) {
		return nil, modelErrors.ErrTokenExp
	}
	account, err := s.account.GetById(ctx, rt.AccountID)
	if err != nil {
		return nil, modelErrors.ErrAccountNotFound
	}
	accessToken, accessExp := generateJWT(account.ID, account.Role, cfg.JWTsecret, cfg.AccessToken)
	return &modelAuth.RefreshTokenResponse{
		AccessToken:      accessToken,
		AccessExpiresAt:  accessExp,
		RefreshToken:     rt.Token,
		RefreshExpiresAt: rt.ExpiresAt.Unix(),
	}, nil
}

func generateCode() string {
	return fmt.Sprintf("%06d", rand.IntN(100000))
}

func (s *AuthService) ParseJWT(acessToken string) (*modelAuth.JWTclaims, error) {
	token, err := jwt.ParseWithClaims(acessToken, &modelAuth.JWTclaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTsecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*modelAuth.JWTclaims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}
	return claims, nil
}
