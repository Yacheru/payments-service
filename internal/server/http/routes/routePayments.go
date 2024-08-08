package routes

import (
	"github.com/gin-gonic/gin"
	"payments-service/internal/kafka/producer"
	"payments-service/internal/server/http/handlers"
)

type RoutePunishments struct {
	handler *handlers.PaymentsHandler
	router  *gin.RouterGroup
}

func NewPaymentsRoute(router *gin.RouterGroup, producer *producer.KafkaProducer) *RoutePunishments {
	handler := handlers.NewPaymentsHandler(producer)

	return &RoutePunishments{
		handler: handler,
		router:  router,
	}
}

func (r *RoutePunishments) Routes() {
	{
		r.router.POST("/create", r.handler.CreatePayment)
		r.router.GET("/accept", r.handler.Accept)
	}
}
