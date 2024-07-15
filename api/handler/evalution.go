package handler

import (
	"encoding/json"
	"log"
	"net/http"

	pb "gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// float
// CreateEvaluation handles creating a new evaluation
// @Summary Create Evaluation
// @Description Create a new evaluation
// @Tags Evaluations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param evaluation body pb.EvaluationCreate true "Evaluation Data"
// @Success 200 {object} pb.Void "Evaluation Created Successfully"
// @Failure 401 {string} string "Error while creating evaluation"
// @Router /evaluations/create [post]
func (h *Handler) CreateEvaluation(ctx *gin.Context) {
	req := &pb.EvaluationCreate{}
	if err := ctx.BindJSON(&req); err != nil {
		panic(err)
	}

	arr, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	err = h.kafka.ProduceMessages("root", arr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err.Error())
		return
	}
	ctx.JSON(200, "SUCCESS!!!")
}

// GetEvaluation handles fetching an evaluation by ID
// @Summary Get Evaluation
// @Description Get an evaluation by ID
// @Tags Evaluations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string true "Evaluation ID"
// @Success 200 {object} pb.EvaluationUpdate "Evaluation Data"
// @Failure 401 {string} string "Error while fetching evaluation"
// @Router /evaluations/{id} [get]
func (h *Handler) GetEvaluation(ctx *gin.Context) {
	req := &pb.Byid{Id: ctx.Query("id")}
	res, err := h.srvs.Evaluation.Get(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// UpdateEvaluation handles updating an existing evaluation
// @Summary Update Evaluation
// @Description Update an existing evaluation
// @Tags Evaluations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param evaluation body pb.EvaluationUpdate true "Evaluation Data"
// @Success 200 {object} pb.Void "Evaluation Updated Successfully"
// @Failure 401 {string} string "Error while updating evaluation"
// @Router /evaluations [put]
func (h *Handler) UpdateEvaluation(ctx *gin.Context) {
	req := &pb.EvaluationUpdate{}
	if err := ctx.BindJSON(&req); err != nil {
		panic(err)
	}
	res, err := h.srvs.Evaluation.Update(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// DeleteEvaluation handles deleting an evaluation by ID
// @Summary Delete Evaluation
// @Description Delete an evaluation by ID
// @Tags Evaluations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string true "Evaluation ID"
// @Success 200 {object} pb.Void "Evaluation Deleted Successfully"
// @Failure 401 {string} string "Error while deleting evaluation"
// @Router /evaluations/{id} [delete]
func (h *Handler) DeleteEvaluation(ctx *gin.Context) {
	req := &pb.Byid{Id: ctx.Query("id")}
	res, err := h.srvs.Evaluation.Delete(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetAllEvaluations handles fetching all evaluations with optional filters
// @Summary Get All Evaluations
// @Description Get all evaluations with optional filters
// @Tags Evaluations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param evaluator_id query string false "Evaluator ID"
// @Param employee_id query string false "Employee ID"
// @Success 200 {object} pb.EvaluationGetAllRes "List of Evaluations"
// @Failure 401 {string} string "Error while fetching evaluations"
// @Router /evaluations [get]
func (h *Handler) GetAllEvaluations(ctx *gin.Context) {
	req := &pb.EvaluationGetAllReq{
		EvaluatorId: ctx.Query("evaluator_id"),
		EmployeeId:  ctx.Query("employee_id"),
	}
	res, err := h.srvs.Evaluation.GetAll(ctx, req)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}
