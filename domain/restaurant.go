package domain

import (
	"errors"
	"regexp"
)

type Restaurant struct {
	// use int instead of uint since JSON doesn't distinguish and to avoid casting
	ID          int64   `json:"id"`
	Position    string  `json:"position"`
	Name        string  `json:"name"`
	Score       float64 `json:"score"`
	Ratings     int64   `json:"ratings"`
	Category    string  `json:"category"`
	PriceRange  string  `json:"price_range"`
	FullAddress string  `json:"full_address"`
	// use int instead of uint since JSON doesn't distinguish and to avoid casting
	ZipCode int64 `json:"zip_code"`
	// JS might not handle very large integers or high-precision floating numbers accurately
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

var (
	regexContainsOnlySymbol = regexp.MustCompile(`^\$+$`)
	regexContainsOnlyLetter = regexp.MustCompile(`^[a-zA-Z]+$`)
)

func ValidateScore(score float64) error {
	if score < 0 || score > 5.0 {
		return errors.New("invalid rating")
	}
	return nil
}

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
