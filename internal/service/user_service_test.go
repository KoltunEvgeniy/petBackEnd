package service

// import (
// 	"database/sql"
// 	"errors"
// 	"meawby/internal/model"
// 	"testing"

// 	"github.com/google/uuid"
// )

// type MockUserRepository struct {
// 	Createfn     func(user *model.User) error
// 	GetByIdfn    func(id uuid.UUID) (*model.User, error)
// 	GetAllfn     func() ([]model.User, error)
// 	Updatefn     func(id uuid.UUID, user *model.UserUpdate) error
// 	Deletefn     func(id uuid.UUID) error
// 	GetByPhonefn func(phone string) (*model.User, error)
// }

// func (m *MockUserRepository) Create(user *model.User) error {
// 	if m.Createfn == nil {
// 		return nil
// 	}
// 	return m.Createfn(user)
// }
// func (m *MockUserRepository) GetById(id uuid.UUID) (*model.User, error) {
// 	if m.GetByIdfn == nil {
// 		return nil, nil
// 	}
// 	return m.GetByIdfn(id)
// }
// func (m *MockUserRepository) GetAll() ([]model.User, error) {
// 	if m.GetAllfn == nil {
// 		return nil, nil
// 	}
// 	return m.GetAll()
// }

// func (m *MockUserRepository) Update(id uuid.UUID, user *model.UserUpdate) error {
// 	if m.Updatefn == nil {
// 		return nil
// 	}
// 	return m.Updatefn(id, user)
// }

// func (m *MockUserRepository) Delete(id uuid.UUID) error {
// 	if m.Deletefn == nil {
// 		return nil
// 	}
// 	return m.Deletefn(id)
// }

// func (m *MockUserRepository) GetByPhone(phone string) (*model.User, error) {
// 	if m.Deletefn == nil {
// 		return nil, nil
// 	}
// 	return m.GetByPhonefn(phone)
// }

// func TestUserService_Create(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		user    *model.User
// 		mock    *MockUserRepository
// 		wantErr bool
// 	}{{name: "success",
// 		user: &model.User{
// 			ID:    uuid.New(),
// 			Name:  "Ludmilka",
// 			Phone: "+123",
// 			Email: nil,
// 			Role:  "admin",
// 		}, mock: &MockUserRepository{
// 			Createfn: func(user *model.User) error {
// 				return nil
// 			},
// 		}, wantErr: false},
// 		{name: "repo_error",
// 			user: &model.User{
// 				ID:    uuid.New(),
// 				Name:  "Ludmilka",
// 				Phone: "123",
// 				Role:  "admin",
// 			}, mock: &MockUserRepository{
// 				Createfn: func(user *model.User) error {
// 					return errors.New("db_error")
// 				},
// 			}, wantErr: true}}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			service := NewUserServ(tt.mock)

// 			_, err := service.CreateUser(&model.UserRequest{
// 				Name:  tt.user.Name,
// 				Phone: tt.user.Phone,
// 				Email: tt.user.Email,
// 				Role:  tt.user.Role,
// 			})

// 			if tt.wantErr && err == nil {
// 				t.Fatalf("expected err: %v", err)
// 			}
// 			if !tt.wantErr && err != nil {
// 				t.Fatalf("unexpected err: %v", err)
// 			}
// 		})
// 	}

// }

// func TestUserService_Create_Validation(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		user    *model.User
// 		wantErr error
// 	}{
// 		{name: "name_error",
// 			user: &model.User{
// 				ID:    uuid.New(),
// 				Name:  "",
// 				Phone: "1213213",
// 				Role:  "admin",
// 			}, wantErr: model.ErrName},
// 		{name: "phone_error",
// 			user: &model.User{
// 				ID:    uuid.New(),
// 				Name:  "Ludmilka",
// 				Phone: "",
// 				Role:  "admin",
// 			}, wantErr: model.ErrPhone},
// 		{name: "role_error",
// 			user: &model.User{
// 				ID:    uuid.New(),
// 				Name:  "Ludmilka",
// 				Phone: "134124",
// 				Role:  "",
// 			}, wantErr: model.ErrRole},
// 		{name: "all_error",
// 			user: &model.User{
// 				ID:    uuid.New(),
// 				Name:  "",
// 				Phone: "",
// 				Role:  "",
// 			}, wantErr: model.ErrAllEmpty},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			repo := &MockUserRepository{
// 				Createfn: func(user *model.User) error {

// 					return nil
// 				},
// 			}
// 			service := NewUserServ(repo)
// 			_, err := service.CreateUser(&model.UserRequest{
// 				Name:  tt.user.Name,
// 				Phone: tt.user.Phone,
// 				Role:  tt.user.Role,
// 			})
// 			if !errors.Is(err, tt.wantErr) {
// 				t.Fatalf("expected %v, got %v", tt.wantErr, err)
// 			}
// 		})
// 	}
// }

// func TestUserService_GetById(t *testing.T) {
// 	userID := uuid.New()
// 	tests := []struct {
// 		name    string
// 		id      uuid.UUID
// 		mock    *MockUserRepository
// 		wantErr bool
// 	}{{name: "success",
// 		id: userID, mock: &MockUserRepository{
// 			GetByIdfn: func(id uuid.UUID) (*model.User, error) {
// 				return &model.User{
// 					ID:    userID,
// 					Name:  "Ludmilka",
// 					Phone: "123",
// 					Email: nil,
// 					Role:  "admin",
// 				}, nil
// 			},
// 		}, wantErr: false},
// 		{name: "not_found",
// 			id: userID, mock: &MockUserRepository{
// 				GetByIdfn: func(id uuid.UUID) (*model.User, error) {
// 					return nil, sql.ErrNoRows
// 				},
// 			}, wantErr: true}}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			service := NewUserServ(tt.mock)

// 			_, err := service.GetById(tt.id)

// 			if tt.wantErr && err == nil {
// 				t.Fatalf("expected err: got nil")
// 			}
// 			if !tt.wantErr && err != nil {
// 				t.Fatalf("unexpected err: %v", err)
// 			}
// 		})
// 	}

// }

// func TestUserService_Update(t *testing.T) {
// 	userID := uuid.New()
// 	name := "newName"
// 	tests := []struct {
// 		name    string
// 		update  *model.UserUpdate
// 		mock    *MockUserRepository
// 		wantErr bool
// 	}{{name: "success",
// 		update: &model.UserUpdate{
// 			Name: &name,
// 		}, mock: &MockUserRepository{
// 			Updatefn: func(id uuid.UUID, user *model.UserUpdate) error {
// 				return nil
// 			},
// 		}, wantErr: false},
// 		{name: "empty_update",
// 			update: &model.UserUpdate{},
// 			mock: &MockUserRepository{
// 				Updatefn: func(id uuid.UUID, user *model.UserUpdate) error {

// 					return nil
// 				},
// 			}, wantErr: true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			service := NewUserServ(tt.mock)

// 			err := service.Update(userID, tt.update)

// 			if tt.wantErr && err == nil {
// 				t.Fatalf("expected err %v", err)
// 			}
// 			if !tt.wantErr && err != nil {
// 				t.Fatalf("unexpected err: %v", err)
// 			}
// 		})
// 	}

// }

// func TestUserService_Update_Validation(t *testing.T) {
// 	userId := uuid.New()
// 	nameEmpty := ""
// 	phoneEmpty := ""
// 	roleEmpty := ""
// 	emailEmpty := ""
// 	tests := []struct {
// 		name    string
// 		user    *model.UserUpdate
// 		wantErr error
// 	}{
// 		{name: "name_error",
// 			user: &model.UserUpdate{
// 				Name: &nameEmpty,
// 			}, wantErr: model.ErrName},
// 		{name: "phone_error",
// 			user: &model.UserUpdate{
// 				Phone: &phoneEmpty,
// 			}, wantErr: model.ErrPhone},
// 		{name: "role_error",
// 			user: &model.UserUpdate{
// 				Role: &roleEmpty,
// 			}, wantErr: model.ErrRole},
// 		{name: "email_error",
// 			user: &model.UserUpdate{
// 				Email: &emailEmpty,
// 			}, wantErr: model.ErrEmail},
// 		{name: "all_error",
// 			user: &model.UserUpdate{
// 				Name:  &nameEmpty,
// 				Phone: &phoneEmpty,
// 				Email: &emailEmpty,
// 				Role:  &roleEmpty,
// 			}, wantErr: model.ErrAllEmpty},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			repo := &MockUserRepository{
// 				Updatefn: func(id uuid.UUID, user *model.UserUpdate) error {
// 					return nil
// 				},
// 			}
// 			service := NewUserServ(repo)
// 			err := service.Update(userId, &model.UserUpdate{
// 				Name:  tt.user.Name,
// 				Phone: tt.user.Phone,
// 				Email: tt.user.Email,
// 				Role:  tt.user.Role,
// 			})
// 			if !errors.Is(err, tt.wantErr) {
// 				t.Fatalf("expected %v, got %v", tt.wantErr, err)
// 			}
// 		})
// 	}
// }

// func TestUserService_Delete(t *testing.T) {
// 	userId := uuid.New()
// 	test := []struct {
// 		name    string
// 		id      uuid.UUID
// 		mock    *MockUserRepository
// 		wantErr bool
// 	}{
// 		{
// 			name: "success",
// 			id:   userId,
// 			mock: &MockUserRepository{
// 				Deletefn: func(id uuid.UUID) error {
// 					return nil
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "not_found",
// 			id:   userId,
// 			mock: &MockUserRepository{
// 				Deletefn: func(id uuid.UUID) error {
// 					return sql.ErrNoRows
// 				},
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range test {
// 		t.Run(tt.name, func(t *testing.T) {
// 			service := NewUserServ(tt.mock)
// 			err := service.Delete(userId)

// 			if tt.wantErr && err == nil {
// 				t.Fatalf("expected err %v", err)
// 			}
// 			if !tt.wantErr && err != nil {
// 				t.Fatalf("unexpected err: %v", err)
// 			}
// 		})
// 	}
// }
