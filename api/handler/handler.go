package handler

import (
	pb "github.com/Azizbek-Qodirov/hr_platform_api_service/genprotos"
	"github.com/Azizbek-Qodirov/hr_platform_api_service/kafka"
)

type Handler struct {
	Auth              pb.AuthServiceClient
	EvaluationStorage pb.EvaulationServiceClient
	GuideStorage      pb.GuideServiceClient
    NotificationStorage pb.ServiceNotificationClient
	kafka 				kafka.KafkaProducer
}

func NewHandler(auth pb.AuthServiceClient,
	EvaluationStorage pb.EvaulationServiceClient,
	GuideStorage pb.GuideServiceClient,
    NotificationStorage pb.ServiceNotificationClient,
	Kafka kafka.KafkaProducer) *Handler {
	return &Handler{
		Auth:              auth,
		EvaluationStorage: EvaluationStorage,
		GuideStorage:      GuideStorage,
        NotificationStorage: NotificationStorage,
		kafka: Kafka,

	}
}
