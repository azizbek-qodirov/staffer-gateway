package handler

import (
	pb "gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateNotification handles creating a new notification
// @Summary Create Notification
// @Description Create a new notification
// @Tags Notification
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param notification body pb.CreateNotification true "Notification data"
// @Success 200 {object} pb.Void "Notification created successfully"
// @Failure 401 {string} string "Error while creating notification"
// @Router /notifications [post]
func (h *Handler) CreateNotification(ctx *gin.Context) {
	req := &pb.CreateNotification{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	res, err := h.srvs.Notification.Create(ctx, req)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// GetNotificationsByUser handles fetching notifications by user ID
// @Summary Get Notifications
// @Description Get notifications by user ID
// @Tags Notification
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string true "User ID"
// @Param notification_id query string false "Notification ID"
// @Success 200 {object} pb.GetAllResponse "Notifications fetched successfully"
// @Failure 401 {string} string "Error while fetching notifications"
// @Router /notifications [get]
func (h *Handler) GetNotificationsByUser(ctx *gin.Context) {
	req := &pb.GetAllRequest{
		UserId: ctx.Query("user_id"),
	}
	res, err := h.srvs.Notification.GetByUserId(ctx, req)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// ReadAllNotifications handles marking all notifications as read for a user
// @Summary Read All Notifications
// @Description Mark all notifications as read for a user
// @Tags Notification
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id query string true "User ID"
// @Success 200 {object} pb.Void "All notifications marked as read"
// @Failure 401 {string} string "Error while marking notifications as read"
// @Router /notifications/read_all [post]
func (h *Handler) ReadAllNotifications(ctx *gin.Context) {
	req := &pb.ReadeAllRequest{UserId: ctx.Query("user_id")}
	res, err := h.srvs.Notification.ReadeAll(ctx, req)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// SendNotificationToAll handles sending a notification to all users in a company
// @Summary Send Notification To All
// @Description Send a notification to all users in a company
// @Tags Notification
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param company_id query string true "Company ID"
// @Param notification body pb.SendByCompanyidToUsers true "Notification data"
// @Success 200 {object} pb.Void "Notification sent to all users"
// @Failure 401 {string} string "Error while sending notification"
// @Router /notifications/send_all [post]
func (h *Handler) SendNotificationToAll(ctx *gin.Context) {
	req := &pb.SendByCompanyidToUsers{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	req.CompanyId = ctx.Query("company_id")
	res, err := h.srvs.Notification.SendAll(ctx, req)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}

// SendNotificationToAllUsers handles sending a notification to all users
// @Summary Send Notification To All Users
// @Description Send a notification to all users
// @Tags Notification
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param notification body pb.SendByCompanyidToUsers true "Notification data"
// @Success 200 {object} pb.Void "Notification sent to all users"
// @Failure 401 {string} string "Error while sending notification"
// @Router /notifications/send_all_users [post]
func (h *Handler) SendNotificationToAllUsers(ctx *gin.Context) {
	req := &pb.SendByCompanyidToUsers{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	res, err := h.srvs.Notification.SendAll(ctx, req)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
}
