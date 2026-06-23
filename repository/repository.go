package repository

import (
	"github.com/your-org/your-app/contract"
)

func New() *contract.Repository {
	return &contract.Repository{
		HealthRepo: ImplHealthRepository(),
	}
}
