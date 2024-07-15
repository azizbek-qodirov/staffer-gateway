package handler

import (
	pb "gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateGuide handles creating a new guide
// @Summary Create Guide
// @Description Create a new guide
// @Tags Guides
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param guide body pb.CreateGuideRequest true "Guide Data"
// @Success 200 {object} pb.Void "Guide Created Successfully"
// @Failure 401 {string} string "Error while creating guide"
// @Router /guides [post]
func (h *Handler) CreateGuide(ctx *gin.Context) {
	req := &pb.CreateGuideRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		panic(err)
	}
	res, err := h.srvs.Guide.CreateGuide(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetGuide handles fetching a guide by ID
// @Summary Get Guide
// @Description Get a guide by ID
// @Tags Guides
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string true "Guide ID"
// @Success 200 {object} pb.GuideResponse "Guide Data"
// @Failure 401 {string} string "Error while fetching guide"
// @Router /guides/{id} [get]
func (h *Handler) GetGuide(ctx *gin.Context) {
	req := &pb.GetGuideRequest{Id: ctx.Query("id")}
	res, err := h.srvs.Guide.GetGuide(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// UpdateGuide handles updating an existing guide
// @Summary Update Guide
// @Description Update an existing guide
// @Tags Guides
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param guide body pb.UpdateGuideRequest true "Guide Data"
// @Success 200 {object} pb.Void "Guide Updated Successfully"
// @Failure 401 {string} string "Error while updating guide"
// @Router /guides [put]
func (h *Handler) UpdateGuide(ctx *gin.Context) {
	req := &pb.UpdateGuideRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		panic(err)
	}
	res, err := h.srvs.Guide.UpdateGuide(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// DeleteGuide handles deleting a guide by ID
// @Summary Delete Guide
// @Description Delete a guide by ID
// @Tags Guides
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string true "Guide ID"
// @Success 200 {object} pb.Void "Guide Deleted Successfully"
// @Failure 401 {string} string "Error while deleting guide"
// @Router /guides/{id} [delete]
func (h *Handler) DeleteGuide(ctx *gin.Context) {
	req := &pb.DeleteGuideRequest{Id: ctx.Query("id")}
	res, err := h.srvs.Guide.DeleteGuide(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// ListAllGuides handles listing all guides with optional filters
// @Summary List All Guides
// @Description List all guides with optional filters
// @Tags Guides
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param title query string false "Guide Title"
// @Param content query string false "Guide Content"
// @Success 200 {object} pb.ListAllGuidesResponse "List of Guides"
// @Failure 401 {string} string "Error while listing guides"
// @Router /guides [get]
func (h *Handler) ListAllGuides(ctx *gin.Context) {
	req := &pb.ListAllGuidesRequest{
		Title:   ctx.Query("title"),
		Content: ctx.Query("content"),
	}
	res, err := h.srvs.Guide.ListAllGuides(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// SearchGuides handles searching for guides by title or content
// @Summary Search Guides
// @Description Search for guides by title or content
// @Tags Guides
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param title query string false "Guide Title"
// @Param content query string false "Guide Content"
// @Success 200 {object} pb.ListAllGuidesResponse "List of Found Guides"
// @Failure 401 {string} string "Error while searching guides"
// @Router /guides/search [get]
func (h *Handler) SearchGuides(ctx *gin.Context) {
	req := &pb.SearchGuidesRequest{
		Title:   ctx.Query("title"),
		Content: ctx.Query("content"),
	}
	res, err := h.srvs.Guide.SearchGuides(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}
