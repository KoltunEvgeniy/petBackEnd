package handler

import (
	"meawby/internal/handler/adminHandler"
	"meawby/internal/handler/authHandler"
	"meawby/internal/handler/middleware"
	paymenthandler "meawby/internal/handler/paymentHandler"
	scheduleHandler "meawby/internal/handler/schedule"
	"meawby/internal/handler/serviceHandler"
	"meawby/internal/handler/userHandler"
	"meawby/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Auth                  *authHandler.AuthHandl
	Admin                 *adminHandler.AdminHandl
	Middleware            *middleware.Middleware
	ServicesHandler       *serviceHandler.ServiceHandl
	MasterHandler         *userHandler.MasterHandler
	AppointmentHandler    *scheduleHandler.AppointmentHandl
	AvailiabilityHandler  *scheduleHandler.AvailiabilityHandl
	MasterScheduleHandler *scheduleHandler.MasterScheduleHandl
	SlotHandler           *scheduleHandler.SlotHandl
	ClientHandler         *userHandler.ClientHandler
	PaymentHandler        *paymenthandler.PaymentHandl
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Auth:                  authHandler.NewAuthHandler(service),
		Admin:                 adminHandler.NewAdminHandler(service),
		Middleware:            middleware.NewMiddleware(service),
		ServicesHandler:       serviceHandler.NewServiceHandl(service),
		MasterHandler:         userHandler.NewMasterHandler(service),
		AppointmentHandler:    scheduleHandler.NewAppointmentHandler(service),
		AvailiabilityHandler:  scheduleHandler.NewAvailiabilityHandl(service),
		MasterScheduleHandler: scheduleHandler.NewMasterScheduleHandl(service),
		SlotHandler:           scheduleHandler.NewSlotHandl(service),
		ClientHandler:         userHandler.NewClientHandler(service),
		PaymentHandler:        paymenthandler.NewPaymentHandl(service),
	}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Auth.SendSMS)
		auth.POST("/verify", h.Auth.Verify)
		auth.POST("/refresh", h.Auth.Refresh)
	}
	admin := router.Group("/admin", h.Middleware.UserIndentity, h.Middleware.RequereRole("admin"))
	{
		client := admin.Group("/client")
		{
			client.GET("/", h.Admin.GetAll)
			client.PATCH("/:id", h.Admin.UpdateRole("master"))
		}
		services := admin.Group("/services")
		{
			services.POST("/", h.ServicesHandler.CreateService)
			services.GET("/", h.ServicesHandler.GetServices)
		}
	}
	master := router.Group("/master", h.Middleware.UserIndentity, h.Middleware.RequereRole("master"))
	{

		me := master.Group("/me")
		{
			me.POST("/", h.MasterHandler.CreateMasterProfile)
			me.GET("/", h.MasterHandler.GetMasterProfile)
			services := me.Group("/services")
			{
				services.GET("/", h.ServicesHandler.GetServices)
				myServices := services.Group("/my")
				{
					myServices.POST("/", h.MasterHandler.AddService)
					myServices.GET("/", h.MasterHandler.ListService)
					myServices.DELETE("/:id", h.MasterHandler.RemoveService)
				}
			}
			appointments := me.Group("/appointments")
			{
				appointments.GET("/", h.AppointmentHandler.ListAppointmentMaster)
			}
			availability := me.Group("/availability")
			{
				availability.POST("/", h.AvailiabilityHandler.SetAvailiability)
				availability.GET("/", h.AvailiabilityHandler.GetAvailiabilityByMaster)
			}
			schedule := me.Group("/schedule")
			{
				schedule.POST("/", h.MasterScheduleHandler.AddMasterSchedule)
				schedule.GET("/", h.MasterScheduleHandler.GetMasterSchedule)
			}
			slots := me.Group("/slots")
			{
				slots.GET("/", h.SlotHandler.GetMasterSlotsM)
			}
		}
	}
	client := router.Group("/client", h.Middleware.UserIndentity, h.Middleware.RequereRole("client"))
	{
		master := client.Group("/master")
		{
			master.GET("/", h.ClientHandler.GetAllActiveMasters)
			master.GET("/:id/availability", h.AvailiabilityHandler.GetAvailiabilityByClient)
			master.GET("/:id/slots", h.SlotHandler.GetMasterSlotsC)
		}
		me := client.Group("/me")
		{
			me.POST("/", h.ClientHandler.CreateClientProfile)
			me.GET("/", h.ClientHandler.GetClientProfile)
			appointments := me.Group("/appointments")
			{
				appointments.POST("/", h.AppointmentHandler.CreateAppointment)
				appointments.GET("/", h.AppointmentHandler.ListAppointmentClient)

				payments := appointments.Group("/:id/payments")
				{
					payments.POST("/", h.PaymentHandler.CreatePayment)
					payments.GET("/", h.PaymentHandler.ListPayment)
				}
			}

		}
		client.POST("/admn", h.Admin.UpdateRoleNow("admin"))
		client.POST("/mstr", h.Admin.UpdateRoleNow("master"))

	}
	return router
}
