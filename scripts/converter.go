package main

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/config"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"gorm.io/gorm/clause"
	"os"
	"strconv"
)

// ParseRestaurantCSV parses data from csv dataset to database
func ParseRestaurantCSV() {
	x := viper.GetString("paths.restaurant_csv")
	csvFile, err := os.Open(x)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	var restaurants []entity.Restaurant
	for _, each := range csvData {
		if len(each) < 11 {
			log.Printf("Encountered a row with insufficient fields: %v", each)
			continue
		}

		score, _ := strconv.ParseFloat(each[3], 64)
		ratings, _ := strconv.ParseInt(each[4], 10, 64)
		position, _ := strconv.ParseInt(each[5], 10, 64)
		restaurant := entity.Restaurant{
			Position:    position,
			Name:        each[2],
			Score:       score,
			Ratings:     ratings,
			Category:    each[5],
			PriceRange:  each[6],
			FullAddress: each[7],
			ZipCode:     each[8],
			Lat:         each[9],
			Lng:         each[10],
		}

		restaurants = append(restaurants, restaurant)
	}

	db := config.GetDB()

	result := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(&restaurants, 100)
	if result.Error != nil {
		log.Fatalf("Error inserting data into database: %v", result.Error)
	} else {
		log.Printf("Rows affected (inserted or updated): %d", result.RowsAffected)
	}
}

// ParseMenuCSV parses data from csv dataset to database
func ParseMenuCSV() {
	x := viper.GetString("paths.menu_csv")
	csvFile, err := os.Open(x)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	var menus []entity.Menu
	for _, each := range csvData {
		if len(each) < 5 {
			log.Printf("Encountered a row with insufficient fields: %v", each)
			continue
		}

		restaurantID, _ := strconv.ParseInt(each[0], 10, 64)
		m := entity.Menu{
			RestaurantID: restaurantID,
			Category:     each[1],
			Name:         each[2],
			Description:  each[3],
			Price:        each[4],
		}

		menus = append(menus, m)
	}

	db := config.GetDB()

	result := db.CreateInBatches(menus, 50)
	if result.Error != nil {
		log.Fatalf("Error inserting data into database: %v", result.Error)
	} else {
		log.Printf("Rows affected (inserted or updated): %d", result.RowsAffected)
	}
}

func main() {
	ParseRestaurantCSV()
	ParseMenuCSV()
}
