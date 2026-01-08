package repository

import (
	"context"

	"meawby/internal/model/modelAuth"
	"meawby/internal/model/modelPayment"
	"meawby/internal/model/modelSchedule"
	"meawby/internal/model/modelService"
	"meawby/internal/model/modelUser"
	"meawby/internal/repository/auth"
	"meawby/internal/repository/payment"
	"meawby/internal/repository/schedule"
	"meawby/internal/repository/service"
	"meawby/internal/repository/user"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AccountRepository interface {
	GetByPhone(ctx context.Context, phone string) (*modelUser.Account, error)
	Create(ctx context.Context, acc *modelUser.Account) error
	GetById(ctx context.Context, accountID uuid.UUID) (*modelUser.Account, error)
	UpdateRole(ctx context.Context, accountID uuid.UUID, role string) error
	GetAll(ctx context.Context) ([]modelUser.Account, error)
	DeleteAccountByID(ctx context.Context, accountID uuid.UUID) error
}

type SmsCodeRepository interface {
	Create(ctx context.Context, code *modelAuth.SMSCode) error
	GetValidCode(ctx context.Context, accountID uuid.UUID, code string) (*modelAuth.SMSCode, error)
	Delete(ctx context.Context, smsmID uuid.UUID) error
}

type RefreshToken interface {
	Create(ctx context.Context, token *modelAuth.RefreshToken) error
	GetByToken(ctx context.Context, token string) (*modelAuth.RefreshToken, error)
}

type ClientRepository interface {
	Create(ctx context.Context, client *modelUser.Client) error
	GetByAccountID(ctx context.Context, accountID uuid.UUID) (*modelUser.Client, error)
}

type MasterRepository interface {
	Create(ctx context.Context, master *modelUser.Master) error
	GetByAccountId(ctx context.Context, accountID uuid.UUID) (*modelUser.Master, error)
	GetAllMasters(ctx context.Context) ([]modelUser.Master, error)
}

type MasterServiceRepository interface {
	Add(ctx context.Context, ms *modelUser.MasterService) error
	Remove(ctx context.Context, masterID, serviceID uuid.UUID) error
	ListByMaster(ctx context.Context, masterID uuid.UUID) ([]uuid.UUID, error)
	ListByClient(ctx context.Context, masterID uuid.UUID) ([]uuid.UUID, error)
	GetServicesByIDs(ctx context.Context, ids []uuid.UUID) ([]modelUser.MasterServiceBrief, error)
}

type ServiceRepository interface {
	Create(ctx context.Context, s *modelService.Service) error
	GetAll(ctx context.Context) ([]modelService.Service, error)
	GetById(ctx context.Context, serviceID uuid.UUID) (*modelService.Service, error)
}

type AppointmentRepository interface {
	Create(ctx context.Context, appt *modelSchedule.Appointment) error
	GetByID(ctx context.Context, apptID uuid.UUID) (*modelSchedule.Appointment, error)
	ListByClient(ctx context.Context, clientID uuid.UUID) ([]modelSchedule.Appointment, error)
	ListByMaster(ctx context.Context, masterID uuid.UUID) ([]modelSchedule.Appointment, error)
	Exist(ctx context.Context, masterID uuid.UUID, startAt time.Time) (bool, error)
}

type PaymentRepository interface {
	Create(ctx context.Context, p *modelPayment.Payment) error
	GetByAppointmentID(ctx context.Context, appointmentID uuid.UUID) ([]modelPayment.Payment, error)
}

// type MasterDayOfWeekRepository interface {
// 	// ListByMaster(ctx context.Context, masterID uuid.UUID) ([]modelSchedule.MasterDayOff, error)
// 	// Create(ctx context.Context, mdo *modelSchedule.MasterDayOff) error
// }

type MasterScheduleRepository interface {
	Insert(ctx context.Context, s modelSchedule.Schedule) error
	DeleteByMaster(ctx context.Context, masterID uuid.UUID) error
	ListByMaster(ctx context.Context, masterID uuid.UUID) ([]modelSchedule.Schedule, error)
	Save(ctx context.Context, masterID uuid.UUID, s []modelSchedule.Schedule) error
	GetByWeekDay(ctx context.Context, masterID uuid.UUID, weekday int) (*modelSchedule.Schedule, error)
}

type AvailiabilityRepository interface {
	Upset(ctx context.Context, a modelSchedule.Availiability) error
	GetRange(ctx context.Context, masterID uuid.UUID, from, to time.Time) ([]modelSchedule.Availiability, error)
	GetByDay(ctx context.Context, masterID uuid.UUID, date time.Time) (*modelSchedule.Availiability, error)
	Update(ctx context.Context, masterID uuid.UUID, date time.Time, avaible bool) error
}

type Repository struct {
	AccountRepository
	SmsCodeRepository
	RefreshToken
	ClientRepository
	MasterRepository
	MasterServiceRepository
	ServiceRepository
	// MasterDayOfWeekRepository
	MasterScheduleRepository
	PaymentRepository
	AppointmentRepository
	AvailiabilityRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AccountRepository:        user.NewAccountRepo(db),
		SmsCodeRepository:        auth.NewSmsCodeRepo(db),
		RefreshToken:             user.NewRefreshTokenRepo(db),
		ClientRepository:         user.NewClientRepo(db),
		MasterRepository:         user.NewMasterRepo(db),
		MasterServiceRepository:  service.NewMasterServicesRepo(db),
		ServiceRepository:        service.NewServiceRepo(db),
		AppointmentRepository:    schedule.NewAppointmentRepo(db),
		PaymentRepository:        payment.NewPaymentRepo(db),
		MasterScheduleRepository: schedule.NewMasterScheduleRepo(db),
		// MasterDayOfWeekRepository: schedule.NewMasterDayOfWeekRepo(db),
		AvailiabilityRepository: schedule.NewAvailiabilityRepo(db),
	}
}
