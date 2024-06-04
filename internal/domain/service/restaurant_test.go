package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository/mocks"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
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

	restaurantUpdated = entity.Restaurant{
		ID:          2,
		Position:    6,
		Name:        "Philly Cheesesteaks (541-B Graymont Ave)",
		Score:       0,
		Ratings:     0,
		Category:    "Pasta",
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
	emptyRestaurant = entity.Restaurant{}
)

func TestRestaurantService_FilterRestaurants(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Arrange

		mockedRepository := new(mocks.RestaurantRepository)
		mockedRepository.On("FilterRestaurants", "Pizza", "23204", "$").Return(restaurants, nil)
		rs := restaurantService{
			repository: mockedRepository,
		}

		var items []utils.Item
		for _, m := range restaurants {
			items = append(items, m)
		}

		expectedResponse := &utils.PaginatedResponse{
			APIVersion: "1.0",
			Data: utils.Data{
				StartIndex:   1,
				ItemsCount:   len(items),
				ItemsPerPage: len(items),
				Items:        items,
			},
		}

		// Act
		response, _ := rs.FilterRestaurants("Pizza", "23204", "$")

		// Assert
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("Error", func(t *testing.T) {
		// Arrange
		mockedRepository := new(mocks.RestaurantRepository)
		rs := restaurantService{
			repository: mockedRepository,
		}
		mockedRepository.On("FilterRestaurants", "Pasta", "12345", "$").Return(nil, errors.New("fail"))

		//Act
		_, err := rs.FilterRestaurants("Pasta", "12345", "$")

		// Assert
		assert.Error(t, err, "fail")
	})
}

func TestRestaurantService_CreateRestaurant(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Arrange
		mockedRepository := new(mocks.RestaurantRepository)
		rs := restaurantService{
			repository: mockedRepository,
		}

		// Act
		mockedRepository.On("Create", restaurant).Return(nil)
		err := rs.CreateRestaurant(restaurant)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		// Arrange
		mockedRepository := new(mocks.RestaurantRepository)
		rs := restaurantService{
			repository: mockedRepository,
		}
		mockedRepository.On("Create", emptyRestaurant).Return(errors.New("failed to create restaurant"))

		// Act
		err := rs.CreateRestaurant(emptyRestaurant)

		// Assert
		assert.Equal(t, "failed to create restaurant", err.Error())
	})
}
func TestRestaurantService_UpdateRestaurant(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Arrange
		mockedRepository := new(mocks.RestaurantRepository)
		rs := restaurantService{
			repository: mockedRepository,
		}
		mockedRepository.On("Update", restaurant).Return(nil)

		// Act
		err := rs.UpdateRestaurant(restaurant)

		// Assert
		assert.NoError(t, err)

	})
	t.Run("Error", func(t *testing.T) {
		// Arrange
		mockedRepository := new(mocks.RestaurantRepository)
		rs := restaurantService{
			repository: mockedRepository,
		}
		mockedRepository.On("Update", emptyRestaurant).Return(errors.New("failed to update restaurant"))

		//Act
		err := rs.UpdateRestaurant(emptyRestaurant)

		// Assert
		assert.Equal(t, "failed to update restaurant", err.Error())
	})
}

func TestRestaurantService_DeleteRestaurant(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Arrange
		mockedRepository := new(mocks.RestaurantRepository)
		rs := restaurantService{
			repository: mockedRepository,
		}
		mockedRepository.On("Delete", int64(restaurant.ID)).Return(nil)

		// Act
		err := rs.DeleteRestaurant(restaurant.ID)

		// Assert
		assert.NoError(t, err)
	})
	t.Run("Error", func(t *testing.T) {
		mockedRepository := new(mocks.RestaurantRepository)
		rs := restaurantService{
			repository: mockedRepository,
		}
		mockedRepository.On("Delete", emptyRestaurant.ID).Return(errors.New("failed to delete restaurant"))

		// Act
		err := rs.DeleteRestaurant(emptyRestaurant.ID)

		// Assert
		assert.Equal(t, "failed to delete restaurant", err.Error())
	})

}
