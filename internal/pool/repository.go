package pool

import (
	"gorm.io/gorm"

	"go-template/internal/repository"
)

type Repository struct {
	*repository.Temp
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		Temp: repository.NewTemp(db),
	}
}
