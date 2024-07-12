package handler

import (
	pb "github.com/Azizbek-Qodirov/hr_platform_api_service/genprotos"
)

type Handler struct {
	Auth              pb.AuthServiceClient
	EvaluationStorage pb.EvaulationServiceClient
	GuideStorage      pb.GuideServiceClient
    NotificationStorage pb.ServiceNotificationClient
}

func NewHandler(auth pb.AuthServiceClient,
	EvaluationStorage pb.EvaulationServiceClient,
	GuideStorage pb.GuideServiceClient,
    NotificationStorage pb.ServiceNotificationClient) *Handler {
	return &Handler{
		Auth:              auth,
		EvaluationStorage: EvaluationStorage,
		GuideStorage:      GuideStorage,
        NotificationStorage: NotificationStorage,
	}
}
