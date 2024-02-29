package handler

import (
	"fmt"
	"registration/core/domain"

	repo "registration/repo"

	"github.com/gin-gonic/gin"
)

type RegHandler struct {
	rr repo.RegRepository
}

func NewRegHandler(rr repo.RegRepository) *RegHandler {
	return &RegHandler{
		rr,
	}
}

type CreateReg struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *RegHandler) CreateReg(c *gin.Context) {
	var r2 CreateReg

	err := c.ShouldBindJSON(&r2)
	if err != nil {
		fmt.Println("error in handler")
	}

	register := domain.Reg{
		Name:  r2.Name,
		Email: r2.Email,
	}
	_, _ = r.rr.CreateReg(c, &register)

	// if err != nil {
	// 	fmt.Println("error while creating")
	// 	return
	// }
	rsp := newRegisterResponse(&register)
	HandleSuccess(c, rsp)
}
