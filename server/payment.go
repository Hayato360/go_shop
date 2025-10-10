package server

import (
	"github.com/Hayato360/go_shop/modules/payment/paymentHandler"
	"github.com/Hayato360/go_shop/modules/payment/paymentRepository"
	"github.com/Hayato360/go_shop/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentrepository.NewPaymentRepository(s.db)
	usecase := paymentusecase.NewPaymentUsecase(repo)
	httpHandler := paymenthandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymenthandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	// Health check
	payment.GET("", s.healthCheckService)
}
