package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository/mocks"
	"testing"
)

var (
	restaurant = entity.Restaurant{
		ID:          2,
		Position:    5,
		Name:        "Philly Fresh Cheesesteaks (541-B Graymont Ave)",
		Score:       0,
		Ratings:     0,
		Category:    "Pizza",
		PriceRange:  "$",
		FullAddress: "541-B Graymont Ave, Birmingham, AL, 23204",
		ZipCode:     "23204",
		Lat:         "33.43098",
		Lng:         "-86.8565464",
	}

	restaurants = []entity.Restaurant{
		{ID: 2,
			Position:    5,
			Name:        "Philly Fresh Cheesesteaks (541-B Graymont Ave)",
			Score:       0,
			Ratings:     0,
			Category:    "Pizza",
			PriceRange:  "$",
			FullAddress: "541-B Graymont Ave, Birmingham, AL, 23204",
			ZipCode:     "23204",
			Lat:         "33.43098",
			Lng:         "-86.8565464"},
	}
)

func TestRestaurantService_FilterRestaurants(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Arrange

		mockedRepository := new(mocks.RestaurantRepository)
		mockedRepository.On("FilterRestaurants", "Pizza", "23204", "$").Return(restaurants, nil)
		rs := restaurantService{
			repository: mockedRepository,
		}

		// Act
		actualRestaurants, err := rs.FilterRestaurants("Pizza", "23204", "$")
		assert.NoError(t, err)
		assert.Equal(t, restaurants, actualRestaurants)
	})

}
