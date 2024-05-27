package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/repository/mocks"
	"testing"
)

func TestMenuService_GetMenu(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// Arrange
		menu := []entity.Menu{
			{RestaurantID: 1, Category: "Extra Large Pizza", Name: "Extra Large Supreme Slice", Description: "Slice.", Price: "3.99 USD"},
		}

		mockMenuRepository := new(mocks.MenuRepository)
		mockMenuRepository.On("GetMenu", int64(1)).Return(&menu, nil)
		ms := menuService{
			repository: mockMenuRepository,
		}

		// Act
		actualMenu, _ := ms.GetMenu(1)

		// Assert
		assert.Equal(t, &menu, actualMenu)
	})

	t.Run("Error", func(t *testing.T) {
		// Arrange
		mockMenuRepository := new(mocks.MenuRepository)
		mockMenuRepository.On("GetMenu", int64(1)).Return(nil, errors.New("Unexpected error"))
		ms := menuService{
			repository: mockMenuRepository,
		}

		// Act
		actualMenu, err := ms.GetMenu(1)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, actualMenu)

		mockMenuRepository.AssertExpectations(t)

	})
}
