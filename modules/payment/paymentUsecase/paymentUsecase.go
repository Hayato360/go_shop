package paymentusecase

import "github.com/Hayato360/go_shop/modules/payment/paymentRepository"

type (
	PaymentUsecaseService interface{}

	paymentUsecase struct {
		paymentRepository paymentrepository.PaymentRepositoryService
	}
)

func NewPaymentUsecase(paymentRepository paymentrepository.PaymentRepositoryService) PaymentUsecaseService{
	return &paymentUsecase{
		paymentRepository: paymentRepository,
	}
}