package repository

import (
	"registration/core/domain"

	"github.com/gin-gonic/gin"
)

type RegRepository struct {
	db *DB
}

func NewRegRepository(db *DB) *RegRepository {
	return &RegRepository{
		db,
	}
}

func (r *RegRepository) CreateReg(c *gin.Context, reg *domain.Reg) (*domain.Reg, error) {
	query := psql.Insert("register").
		Columns("name", "email").
		Values(reg.Name, reg.Email)
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(c, sql, args...).Scan(
		&reg.Name,
		&reg.Email,
	)
	if err != nil {
		return nil, err
	}
	return reg, nil
}
