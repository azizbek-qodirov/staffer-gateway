package clients

import (
	"gateway/config"
	cp "gateway/genprotos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Company          cp.CompanyServiceClient
	VacancyService   cp.VacancyServiceClient
	JobsService      cp.JobServiceClient
	JobsAppService   cp.RecruitmentServiceClient
	JobsOfferService cp.JobOfferServiceClient
	Department       cp.DepartmentServiceClient
	Position         cp.PositionServiceClient
	Resume           cp.ResumeServiceClient

	Auth         cp.AuthServiceClient
	Evaluation   cp.EvaulationServiceClient
	Guide        cp.GuideServiceClient
	Notification cp.ServiceNotificationClient
}

func NewGrpcClients(cfg *config.Config) (*GrpcClients, error) {
	connC, err := grpc.NewClient("localhost"+cfg.RECRUITMENT_SERVICE_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer connC.Close()

	return &GrpcClients{
		Company:          cp.NewCompanyServiceClient(connC),
		VacancyService:   cp.NewVacancyServiceClient(connC),
		JobsService:      cp.NewJobServiceClient(connC),
		JobsAppService:   cp.NewRecruitmentServiceClient(connC),
		JobsOfferService: cp.NewJobOfferServiceClient(connC),
		Department:       cp.NewDepartmentServiceClient(connC),
		Position:         cp.NewPositionServiceClient(connC),
		Resume:           cp.NewResumeServiceClient(connC),

		Auth:         cp.NewAuthServiceClient(connC),
		Evaluation:   cp.NewEvaulationServiceClient(connC),
		Guide:        cp.NewGuideServiceClient(connC),
		Notification: cp.NewServiceNotificationClient(connC),
	}, nil
}
