package user

import (
	"context"
	"meawby/internal/model/modelUser"
	"meawby/internal/repository"
	"time"

	"github.com/google/uuid"
)

// import (
// 	"meawby/internal/model"
// 	"meawby/internal/repository"
// 	"time"

// 	"github.com/google/uuid"
// )

type ClientServ struct {
	client repository.ClientRepository
}

func NewClientServ(repo *repository.Repository) *ClientServ {
	return &ClientServ{client: repo.ClientRepository}
}

func (s *ClientServ) Create(ctx context.Context, accountID uuid.UUID, name string) (*modelUser.Client, error) {
	clientNew := modelUser.Client{
		ID:        uuid.New(),
		AccountID: accountID,
		Name:      name,
		CreatedAt: time.Now(),
	}
	if err := s.client.Create(ctx, &clientNew); err != nil {
		return nil, err
	}
	return &clientNew, nil
}

func (s *ClientServ) GetByAccountID(ctx context.Context, accountID uuid.UUID) (*modelUser.Client, error) {
	return s.client.GetByAccountID(ctx, accountID)
}
