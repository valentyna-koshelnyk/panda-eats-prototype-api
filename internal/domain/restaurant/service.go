package restaurant

// Service defines an API for restaurant service to be used by presentation layer
type Service interface {
	GetAllRestaurants() ([]Restaurant, error)
	FilterRestaurants(category string, zip string, priceRange string) ([]Restaurant, error)
	CreateRestaurant(restaurant Restaurant) error
	UpdateRestaurant(restaurant Restaurant) error
	DeleteRestaurant(id int64) error
}

// service gets repository
type service struct {
	repository Repository
}

// NewRestaurantService is a constructor with pointer to service struct which returned as instance of the interface
func NewRestaurantService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAllRestaurants() ([]Restaurant, error) {
	return s.repository.GetAll()
}

func (s *service) FilterRestaurants(category string, zip string, priceRange string) ([]Restaurant, error) {
	return s.repository.FilterRestaurants(category, zip, priceRange)
}

func (s *service) CreateRestaurant(restaurant Restaurant) error {
	err := s.repository.Create(restaurant)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateRestaurant(restaurant Restaurant) error {
	err := s.repository.Update(restaurant)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteRestaurant(id int64) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
