package api

import (
	"gateway/api/handler"
	"gateway/api/middleware"

	// "gateway/api/middleware"
	"gateway/grpc/clients"
	"gateway/kafka"
	"log"
	"os"

	_ "gateway/api/docs"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(connC *clients.GrpcClients) *gin.Engine {
	if _, err := os.Stat("config/model.conf"); os.IsNotExist(err) {
		log.Fatal("config/model.conf file not found")
	}
	if _, err := os.Stat("config/policy.csv"); os.IsNotExist(err) {
		log.Fatal("config/policy.csv file not found")
	}

	ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		log.Fatal("Casbin enforcer initialization error: ", err)
	}

	err = ca.LoadPolicy()
	if err != nil {
		log.Fatal("Casbin error loading policy: ", err)
	}

	kaf, err := kafka.NewKafkaProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatal("Error while connection kafka: ", err.Error())
	}

	h := handler.NewHandler(connC, kaf)

	router := gin.Default()
	r := router.Group("/")
	r.Use(middleware.CasbinMiddleware(ca))

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	company := r.Group("/companies")
	{
		company.POST("/", h.CreateCompanyHandler)
		company.GET("/:id", h.GetByIdCompanyHandler)
		company.PUT("/:id", h.UpdateCompanyHandler)
		company.DELETE("/:id", h.DeleteCompanyHandler)
		company.GET("/all", h.GetAllCompaniesHandler)
	}

	jobApplications := r.Group("/applications")
	{
		jobApplications.POST("/", h.CreateJobAppHandler)
		jobApplications.POST("/:id/steps", h.CreateJobStepHandler)

		jobApplications.GET("/:id", h.GetByIdJobAppHandler)
		jobApplications.GET("/:id/steps", h.GetJobStepHandler)
		jobApplications.GET("/all", h.GetAllJobAppsHandler)

		jobApplications.PUT("/:applicationId/steps/:stepId", h.UpdateJobStepHandler)
		jobApplications.PUT("/jobApp/:id/confirm", h.UpdateJobAppHandler)

		jobApplications.DELETE("/:applicationId/steps/:stepId", h.DeletedJobStepHandler)
		jobApplications.DELETE("/jobApp/:id", h.DeleteJobAppHandler)
	}

	jobOffer := r.Group("/offers")
	{
		jobOffer.POST("/", h.CreateJobOfferHandler)
		jobOffer.GET("/:id", h.GetByIdJobOfferHandler)
		jobOffer.GET("/all", h.GetAllJobOffersHandler)
		jobOffer.PUT("/:id/confirm", h.UpdateJobOfferHandler)
		jobOffer.DELETE("/:id", h.DeleteJobOfferHandler)
	}

	vacancy := r.Group("/vacancies")
	{
		vacancy.POST("/", h.CreateVacancyHandler)
		vacancy.GET("/:id", h.GetByIdVacancyHandler)
		vacancy.GET("/all", h.GetAllVacancyHandler)
		vacancy.PUT("/:id", h.UpdateVacancyHandler)
		vacancy.DELETE("/:id", h.DeleteVacancyHandler)
		vacancy.GET("/Vacancy/:id/applications", h.VacancyApplicatHundler)
		vacancy.GET("/offers/:id", h.VacanyOffersHundler)
	}

	// company := router.Group("/company")
	// {
	// 	company.POST("/", h.CompanyCreate)
	// 	company.GET("/:id", h.CompanyGetById)
	// 	company.GET("/all", h.CompanyGetAll)
	// 	company.PUT("/:id", h.CompanyUpdate)
	// 	company.DELETE("/:id", h.CompanyDelete)
	// }

	department := r.Group("/department")
	{
		department.POST("/", h.DepartmentCreate)
		department.GET("/:id", h.DepartmentGetById)
		department.GET("/all", h.DepartmentGetAll)
		department.PUT("/:id", h.DepartmentUpdate)
		department.DELETE("/:id", h.DepartmentDelete)
		department.GET("/:id/position", h.PositionGetByDepartment)
	}

	position := r.Group("/position")
	{
		position.POST("/", h.PositionCreate)
		position.GET("/:id", h.PositionGetById)
		position.GET("/all", h.PositionGetAll)
		position.PUT("/:id", h.PositionUpdate)
		position.DELETE("/:id", h.PositionDelete)
	}

	resume := r.Group("/resume")
	{
		resume.POST("/", h.ResumeCreate)
		resume.GET("/:id", h.ResumeGetById)
		resume.GET("/all", h.ResumeGetAll)
		resume.PUT("/:id", h.ResumeUpdate)
		resume.DELETE("/:id", h.ResumeDelete)
	}

	evaluationRoutes := r.Group("/evaluations")
	{
		evaluationRoutes.POST("/create", h.CreateEvaluation)
		evaluationRoutes.GET("/:id", h.GetEvaluation)
		evaluationRoutes.PUT("", h.UpdateEvaluation)
		evaluationRoutes.DELETE("/:id", h.DeleteEvaluation)
		evaluationRoutes.GET("", h.GetAllEvaluations)
	}

	guideRoutes := r.Group("/guides")
	{
		guideRoutes.POST("", h.CreateGuide)
		guideRoutes.GET("/:id", h.GetGuide)
		guideRoutes.PUT("", h.UpdateGuide)
		guideRoutes.DELETE("/:id", h.DeleteGuide)
		guideRoutes.GET("", h.ListAllGuides)
		guideRoutes.GET("/search", h.SearchGuides)
	}

	notificationRoutes := r.Group("/notifications")
	{
		notificationRoutes.POST("", h.CreateNotification)
		notificationRoutes.GET("", h.GetNotificationsByUser)
		notificationRoutes.POST("/read_all", h.ReadAllNotifications)
		notificationRoutes.POST("/send_all", h.SendNotificationToAll)
		notificationRoutes.POST("/send_all_users", h.SendNotificationToAllUsers)
	}

	return router
}
