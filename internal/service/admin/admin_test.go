package admin

// import (
// 	"context"
// 	"errors"
// 	"meawby/internal/model/modelUser"
// 	"testing"

// 	"github.com/google/uuid"
// )

// type AccountRepoMock struct {
// 	UpdateRoleFn        func(ctx context.Context, accountID uuid.UUID, role string) error
// 	GetAllFn            func(ctx context.Context) ([]modelUser.Account, error)
// 	DeleteAccountByIDFn func(ctx context.Context, accountID uuid.UUID) error
// }

// func (m *AccountRepoMock) UpdateRole(ctx context.Context, accountID uuid.UUID, role string) error {
// 	return m.UpdateRoleFn(ctx, accountID, role)
// }
// func (m *AccountRepoMock) GetAll(ctx context.Context) ([]modelUser.Account, error) {
// 	return m.GetAllFn(ctx)
// }

// func (m *AccountRepoMock) DeleteAccountByID(ctx context.Context, accountID uuid.UUID) error {
// 	return m.DeleteAccountByIDFn(ctx, accountID)
// }

// func TestUpdateRole(t *testing.T) {
// 	mockRepo := &AccountRepoMock{
// 		UpdateRoleFn: func(ctx context.Context, accountID uuid.UUID, role string) error {
// 			if role == "" {
// 				return errors.New("role is empty")
// 			}
// 			return nil
// 		},
// 	}
// 	service := &AdminService{account: mockRepo}
// 	err := service.UpdateRole(context.Background(), uuid.New(), "")
// 	if err == nil {

// 	}
// }
