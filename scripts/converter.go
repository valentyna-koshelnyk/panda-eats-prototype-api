package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/config"
	menu2 "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/menu"
	domain "github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/restaurant"
	"gorm.io/gorm/clause"
	"os"
	"strconv"
)

func initConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config.dev")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

// ConverterRestaurantCSVtoJSON converts the CSV file to JSON, using mapping for Restaurant entity
func ConverterRestaurantCSVtoJSON() {
	x := viper.GetString("paths.restaurant_csv")
	csvFile, err := os.Open(x)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}

	reader := csv.NewReader(csvFile)
	// No check for expected field per record
	reader.FieldsPerRecord = -1
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
	prJSON := viper.GetString("paths.restaurants")
	jsonFile, err := os.Create(prJSON)
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Write(jsonData)
	jsonFile.Close()
}

// ConverterMenuCSVtoJSON converts menu csv file to json
func ConverterMenuCSVtoJSON() {
	x := viper.GetString("paths.menu_csv")
	csvFile, err := os.Open(x)
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
		menu.Category = each[1]
		menu.Name = each[2]
		menu.Description = each[3]
		menu.Price = each[4]

		menus = append(menus, menu)
	}

	jsonData, err := json.Marshal(menus)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pmJSON := viper.GetString("paths.menu")
	jsonFile, err := os.Create(pmJSON)
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Write(jsonData)
	jsonFile.Close()
}

// ConverterRestaurantCSVtoDB parses data from csv dataset to database
func ConverterRestaurantCSVtoDB() {
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

	var restaurants []domain.Restaurant
	for _, each := range csvData {
		if len(each) < 11 {
			log.Printf("Encountered a row with insufficient fields: %v", each)
			continue
		}

		score, _ := strconv.ParseFloat(each[3], 64)
		ratings, _ := strconv.ParseInt(each[4], 10, 64)
		restaurant := domain.Restaurant{
			Position:    each[1],
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

	db := config.InitDB()
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	result := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(&restaurants, 100)
	if result.Error != nil {
		log.Fatalf("Error inserting data into database: %v", result.Error)
	} else {
		log.Printf("Rows affected (inserted or updated): %d", result.RowsAffected)
	}
}

func main() {
	ConverterMenuCSVtoJSON()
	ConverterRestaurantCSVtoJSON()
	initConfig()
	ConverterRestaurantCSVtoDB()
}
