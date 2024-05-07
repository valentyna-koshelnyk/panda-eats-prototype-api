package restaurant

import (
	"errors"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/config"
	"gorm.io/gorm"
	"regexp"
)

// A Restaurant represents information about the restaurant entity
type Restaurant struct {
	gorm.Model
	ID          int64   `json:"ID"`
	Position    int64   `json:"position"`
	Name        string  `json:"name"`
	Score       float64 `json:"score"`
	Ratings     int64   `json:"ratings"`
	Category    string  `json:"category"`
	PriceRange  string  `json:"price_range"`
	FullAddress string  `json:"full_address"`
	//TODO: Add additional validation for the zip range
	ZipCode string `json:"zip_code"`
	// JS might not handle very large integers or high-precision floating numbers accurately
	// TODO: add validation
	Lat string `json:"lat"`
	// TODO: add validation
	Lng string `json:"lng"`
}

// Regular expressions validating input, e.g. if price range consists just of '$' sign or category doesn't have any digits
var (
	regexContainsOnlySymbol = regexp.MustCompile(`^\$+$`)
	regexContainsOnlyLetter = regexp.MustCompile(`^[a-zA-Z]+$`)
)

// ValidateScore Validates if score for the restaurant is in a valid range
func ValidateScore(score float64) error {
	if score < 0 || score > 5.0 {
		return errors.New("invalid rating")
	}
	return nil
}

// ValidateRestaurant Validates such restaurant fields as ID, Score, Category and Price Range
func ValidateRestaurant(restaurant *Restaurant) []error {
	errs := make([]error, 0)
	//Validate restaurant.
	if restaurant.ID < 1 {
		errs = append(errs, errors.New("invalid id"))
	}
	if err := ValidateScore(restaurant.Score); err != nil {
		errs = append(errs, errors.New("invalid score"))
	}
	if !regexContainsOnlyLetter.MatchString(restaurant.Category) {
		errs = append(errs, errors.New("invalid category"))
	}
	if !regexContainsOnlySymbol.MatchString(restaurant.PriceRange) {
		errs = append(errs, errors.New("invalid price_range"))
	}
	return errs
}

func init() {
	db := config.GetDB()
	err := db.AutoMigrate(&Restaurant{})
	if err != nil {
		return
	}
}
