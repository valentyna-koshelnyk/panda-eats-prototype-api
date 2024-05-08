package menu

// Service  layer for menu
type Service struct {
	repository Repository
}

// NewService is a constructor for service layer of menu
func NewService(r Repository) Service {
	return Service{repository: r}
}

// GetMenu retrieves menu of the specific restaurant
func (s *Service) GetMenu(id int64) (*[]Menu, error) {
	return s.repository.GetMenu(id)
}
