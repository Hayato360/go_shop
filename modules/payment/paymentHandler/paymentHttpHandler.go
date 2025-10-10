package paymenthandler

import (
	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentusecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
