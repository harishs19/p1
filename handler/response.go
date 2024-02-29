package handler

import (
	"net/http"
	"registration/core/domain"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func newResponse(success bool, message string, data any) Response {
	return Response{
		Status:  success,
		Message: message,
		Data:    data,
	}

}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func newErrorResponse(message string) ErrorResponse {
	return ErrorResponse{

		Status:  false,
		Message: message,
	}
}

type Register struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func newRegisterResponse(r *domain.Reg) Register {
	return Register{
		Name:  r.Name,
		Email: r.Email,
	}
}

func HandleSuccess(r *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	r.JSON(http.StatusOK, rsp)
}

func HandleFailure(r *gin.Context, data any) {
	rsp := newResponse(false, "Failure", data)
	r.JSON(http.StatusOK, rsp)
}
