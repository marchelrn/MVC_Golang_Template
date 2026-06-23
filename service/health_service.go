package service

import "github.com/your-org/your-app/contract"

type HealthService struct {
	repo contract.HealthRepository
}

func ImplHealthService(repo contract.HealthRepository) contract.HealthService {
	return &HealthService{repo: repo}
}

func (s *HealthService) GetStatus() string {
	return s.repo.GetStatus()
}
