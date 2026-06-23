package contract

type Repository struct {
	HealthRepo HealthRepository
}

type HealthRepository interface {
	GetStatus() string
}
