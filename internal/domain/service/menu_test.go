package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository/mocks"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/utils"
)

func TestMenuService_GetMenu(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Arrange
		pagedResponse := &utils.Pagination{
			Limit:      1,
			Page:       1,
			TotalRows:  1,
			TotalPages: 1,
			Rows: []entity.Menu{
				{RestaurantID: 1, Category: "Extra Large Pizza", Name: "Extra Large Supreme Slice", Description: "Slice.", Price: "3.99 USD"},
			},
		}

		mockMenuRepository := new(mocks.MenuRepository)
		mockMenuRepository.On("GetMenu", 1, mock.Anything).Return(pagedResponse, nil)
		ms := menuService{
			repository: mockMenuRepository,
		}

		// Act
		actualMenu, _ := ms.GetMenu(1, 1, 1)

		// Assert
		assert.Equal(t, pagedResponse, actualMenu)
	})

	t.Run("Error", func(t *testing.T) {
		// Arrange
		mockMenuRepository := new(mocks.MenuRepository)
		mockMenuRepository.On("GetMenu", 1, mock.Anything).Return(nil, errors.New("Unexpected error"))
		ms := menuService{
			repository: mockMenuRepository,
		}

		// Act
		actualMenu, err := ms.GetMenu(1, 1, 1)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, actualMenu)

		mockMenuRepository.AssertExpectations(t)

	})
}
