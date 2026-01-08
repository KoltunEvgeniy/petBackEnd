package service

import (
	"context"

	"meawby/internal/model/modelAuth"
	"meawby/internal/model/modelPayment"
	"meawby/internal/model/modelSchedule"
	"meawby/internal/model/modelService"
	"meawby/internal/model/modelUser"
	"meawby/internal/repository"
	"meawby/internal/service/admin"
	"meawby/internal/service/auth"
	"meawby/internal/service/payment"
	"meawby/internal/service/schedule"
	"meawby/internal/service/service"
	"meawby/internal/service/user"
	"time"

	"github.com/google/uuid"
)

type AuthorizationService interface {
	SendSMS(ctx context.Context, phone string) error
	VerifySMS(ctx context.Context, phone, code string) (*modelAuth.VerifySMSResponse, error)
	Refresh(ctx context.Context, tokenStr string) (*modelAuth.RefreshTokenResponse, error)
	ParseJWT(acessToken string) (*modelAuth.JWTclaims, error)
}

type AdminService interface {
	UpdateRole(ctx context.Context, accountID uuid.UUID, role string) error
	GetAll(ctx context.Context) ([]modelUser.Account, error)
	DeleteAccountByID(ctx context.Context, accountID uuid.UUID) error
}

type ClientService interface {
	Create(ctx context.Context, accountID uuid.UUID, name string) (*modelUser.Client, error)
	GetByAccountID(ctx context.Context, accountID uuid.UUID) (*modelUser.Client, error)
}

type MasterService interface {
	Create(ctx context.Context, accountID uuid.UUID, name string) error
	GetByAccountID(ctx context.Context, accountID uuid.UUID) (*modelUser.Master, error)
	GetAll(ctx context.Context) ([]modelUser.Master, error)
}

type MasterServicesService interface {
	Add(ctx context.Context, masterID, serviceID uuid.UUID) error
	Remove(ctx context.Context, masterID, serviceID uuid.UUID) error
	GetProfle(ctx context.Context, accountID uuid.UUID) (*modelUser.Master, error)
	GetList(ctx context.Context, accountID uuid.UUID) (*modelUser.Master, []modelUser.MasterServiceBrief, error)
	GetListForClient(ctx context.Context, masterID uuid.UUID) (*modelUser.Master, []modelUser.MasterServiceBrief, error)
}

type ServicesService interface {
	Create(ctx context.Context, title string, duration int, price int) error
	GetAll(ctx context.Context) ([]modelService.Service, error)
}

type PaymentService interface {
	Create(ctx context.Context, p *modelPayment.Payment) error
	GetByAppointment(ctx context.Context, appointmentID uuid.UUID) ([]modelPayment.Payment, error)
}

type AppointmentService interface {
	Create(ctx context.Context, clientID uuid.UUID, req *modelSchedule.AppointmentRequest) (*modelSchedule.Appointment, error)
	GetByID(ctx context.Context, apptID uuid.UUID) (*modelSchedule.Appointment, error)
	ListByClient(ctx context.Context, clientID uuid.UUID) ([]modelSchedule.Appointment, error)
	ListByMaster(ctx context.Context, masterID uuid.UUID) ([]modelSchedule.Appointment, error)
}

type MasterScheduleServive interface {
	Set(ctx context.Context, masterID uuid.UUID, req []modelSchedule.ScheduleReq) error
	Get(ctx context.Context, masterID uuid.UUID) ([]modelSchedule.MasterScheduleResp, error)
}

type AvailiabilityService interface {
	SetDate(ctx context.Context, masterID uuid.UUID, dateStr string, avible bool) error
	GetRange(ctx context.Context, masterID uuid.UUID, from, to time.Time) ([]modelSchedule.AvailiabilityResp, error)
}

type SlotServive interface {
	GetFreeForClient(ctx context.Context, masterID uuid.UUID, date time.Time) (*modelSchedule.MasterDaySlotsResp, error)
	GetForMaster(ctx context.Context, masterID uuid.UUID, date time.Time) (*modelSchedule.MasterDaySlotsResp, error)
}

type Service struct {
	AuthorizationService
	AdminService
	ClientService
	MasterService
	MasterServicesService
	ServicesService
	PaymentService
	AppointmentService
	MasterScheduleServive
	AvailiabilityService
	SlotServive
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthorizationService:  auth.NewAuthService(repo),
		AdminService:          admin.NewAdminService(repo),
		ClientService:         user.NewClientServ(repo),
		MasterService:         user.NewMasterServ(repo),
		MasterServicesService: service.NewMasterServicesServ(repo),
		ServicesService:       service.NewServiceService(repo),
		PaymentService:        payment.NewPaymentServ(repo),
		AppointmentService:    schedule.NewAppointmentServ(repo),
		MasterScheduleServive: schedule.NewMasterScheduleServ(repo),
		AvailiabilityService:  schedule.NewAvailiabilityServ(repo),
		SlotServive:           schedule.NewSlotServ(repo)}
}
