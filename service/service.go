package service

import "github.com/your-org/your-app/contract"

func New(repo *contract.Repository) (*contract.Service, error) {
	return &contract.Service{
		HealthService: ImplHealthService(repo.HealthRepo),
	}, nil
}
