package handler

import (
	"gateway/grpc/clients"
	"gateway/kafka"
)

type Handler struct {
	srvs  *clients.GrpcClients
	kafka kafka.KafkaProducer
}

func NewHandler(srvs *clients.GrpcClients, Kafka kafka.KafkaProducer) *Handler {
	return &Handler{srvs: srvs, kafka: Kafka}
}
