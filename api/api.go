package api

import (
	"log"
	"os"

	"github.com/Azizbek-Qodirov/hr_platform_api_service/api/handler"
	"github.com/Azizbek-Qodirov/hr_platform_api_service/api/middleware"
	_ "github.com/Azizbek-Qodirov/hr_platform_api_service/docs"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Voting service
// @version 1.0
// @description Voting service
// @host localhost:8082
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization
func NewGin(h *handler.Handler) *gin.Engine {
	if _, err := os.Stat("config/model.conf"); os.IsNotExist(err) {
		log.Fatal("config/model.conf file not found")
	}
	if _, err := os.Stat("config/policy.csv"); os.IsNotExist(err) {
		log.Fatal("config/policy.csv file not found")
	}

	// Initialize Casbin enforcer
	ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		log.Fatal("Casbin enforcer initialization error: ", err)
	}

	err = ca.LoadPolicy()
	if err != nil {
		log.Fatal("Casbin error loading policy: ", err)
	}

	r := gin.Default()
	g := r.Group("/")
	g.Use(middleware.CasbinMiddleware(ca))

	evaluationRoutes := g.Group("/evaluations")
	{
		evaluationRoutes.POST("/create", h.CreateEvaluation)
		evaluationRoutes.GET("/:id", h.GetEvaluation)
		evaluationRoutes.PUT("", h.UpdateEvaluation)
		evaluationRoutes.DELETE("/:id", h.DeleteEvaluation)
		evaluationRoutes.GET("", h.GetAllEvaluations)
	}

	guideRoutes := g.Group("/guides")
	{
		guideRoutes.POST("", h.CreateGuide)
		guideRoutes.GET("/:id", h.GetGuide)
		guideRoutes.PUT("", h.UpdateGuide)
		guideRoutes.DELETE("/:id", h.DeleteGuide)
		guideRoutes.GET("", h.ListAllGuides)
		guideRoutes.GET("/search", h.SearchGuides)
	}

	notificationRoutes := g.Group("/notifications")
	{
		notificationRoutes.POST("", h.CreateNotification)
		notificationRoutes.GET("", h.GetNotificationsByUser)
		notificationRoutes.POST("/read_all", h.ReadAllNotifications)
		notificationRoutes.POST("/send_all", h.SendNotificationToAll)
		notificationRoutes.POST("/send_all_users", h.SendNotificationToAllUsers)
	}
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))
	return r
}
