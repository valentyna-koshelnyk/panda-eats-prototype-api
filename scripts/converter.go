package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	menu2 "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/menu"
	domain "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/restaurant"
	"os"
	"strconv"
)

const RestaurantPath = "data/restaurants.csv"
const MenuPath = "data/restaurant-menus.csv"

// ConverterRestaurant converts the CSV file to JSON, using mapping for Restaurant entity
// TODO: to create a generic converter
func ConverterRestaurant() {
	csvFile, err := os.Open(RestaurantPath)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}

	reader := csv.NewReader(csvFile)
	// No check for expected field per record
	reader.FieldsPerRecord = -1
	// Put quotes for unquoted fields
	reader.LazyQuotes = true

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	restaurant := domain.Restaurant{}
	var restaurants []domain.Restaurant
	for _, each := range csvData {
		if len(each) < 11 {
			fmt.Println("Encountered a row with insufficient fields")
			continue
		}
		restaurant.ID, _ = strconv.ParseInt(each[0], 10, 64)
		restaurant.Position = each[1]
		restaurant.Name = each[2]
		restaurant.Score, _ = strconv.ParseFloat(each[3], 64)
		restaurant.Ratings, _ = strconv.ParseInt(each[4], 10, 64)
		restaurant.Category = each[5]
		restaurant.PriceRange = each[6]
		restaurant.FullAddress = each[7]
		restaurant.ZipCode = each[8]
		restaurant.Lat = each[9]
		restaurant.Lng = each[10]
		restaurants = append(restaurants, restaurant)
	}

	jsonData, err := json.Marshal(restaurants)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	jsonFile, err := os.Create("data/restaurants.json")
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Write(jsonData)
	jsonFile.Close()
}

func ConverterMenu() {
	csvFile, err := os.Open(MenuPath)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}

	reader := csv.NewReader(csvFile)
	// No check for expected field per record
	reader.FieldsPerRecord = -1
	// Put quotes for unquoted fields
	reader.LazyQuotes = true

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	menu := menu2.Menu{}
	var menus []menu2.Menu
	for _, each := range csvData {
		if len(each) < 5 {
			fmt.Println("Encountered a row with insufficient fields")
			continue
		}
		menu.RestaurantID, _ = strconv.ParseInt(each[0], 10, 64)
		menu.Name = each[1]
		menu.Price = each[2]
		menu.Description = each[3]
		menu.Category = each[4]

		menus = append(menus, menu)
	}

	jsonData, err := json.Marshal(menus)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jsonFile, err := os.Create("data/menus.json")
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Write(jsonData)
	jsonFile.Close()
}

func main() {
	ConverterMenu()
	ConverterRestaurant()
}
