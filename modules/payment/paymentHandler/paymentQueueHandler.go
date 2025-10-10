package paymenthandler

import (
	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		cfg            *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentusecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
