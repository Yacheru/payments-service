package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"payments-service/init/config"
	"payments-service/internal/entities"
	"payments-service/internal/kafka/producer"
	"payments-service/internal/server/http/client"
	"payments-service/pkg/util/setups"
)

type PaymentsHandler struct {
	producer *producer.KafkaProducer
}

func NewPaymentsHandler(producer *producer.KafkaProducer) *PaymentsHandler {
	return &PaymentsHandler{
		producer: producer,
	}
}

func (h *PaymentsHandler) CreatePayment(ctx *gin.Context) {
	var service = new(entities.PaymentService)

	if err := ctx.ShouldBindJSON(service); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	yooClient := client.NewClient(config.ServerConfig.YKassaID, config.ServerConfig.YKassaPass)

	pH := client.NewPaymentHandler(yooClient)

	payment, err := setups.SetupPayment(pH, service)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Failed creating payment")

		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "success", payment.Confirmation)
}

func (h *PaymentsHandler) Accept(ctx *gin.Context) {
	var paid = &entities.Paid{}
	if err := ctx.ShouldBindJSON(paid); err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, "Invalid request body")

		return
	}

	message, err := json.Marshal(&entities.KafkaMcMessage{
		Nickname: paid.Object.Metadata.Nickname,
		Duration: paid.Object.Metadata.Duration,
		Service:  paid.Object.Metadata.Service,
	})

	if err != nil {
		NewErrorResponse(ctx, http.StatusForbidden, "Invalid request body")

		return
	}

	err = h.producer.PrepareMessage(message, &config.ServerConfig)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to prepare message")

		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "success", paid.Object.Metadata)
}
