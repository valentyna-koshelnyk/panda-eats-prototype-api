package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Initialize viper and load configuration
func InitViperConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config.dev")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

// Initialize and return the database instance
func GetDB() *gorm.DB {
	once.Do(func() {
		db = initDB()
	})
	return db
}

// Private function to initialize the database
func initDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = db.AutoMigrate(&entity.Menu{}, &entity.Restaurant{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
