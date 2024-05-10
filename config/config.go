package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	db          *gorm.DB
	once        sync.Once
	viperConfig *ViperConfig
)

// ViperConfig configuration for viper package
type ViperConfig struct {
	viper *viper.Viper
}

// getViper() gets instance of Viper
func getViper() *ViperConfig {
	if viperConfig == nil {
		once.Do(func() {
			fmt.Println("Creating Viper Config")
			viperConfig = &ViperConfig{}
		})
	} else {
		fmt.Println("Viper Config created")
	}
	return viperConfig
}

// InitViperConfig initializes viper
func InitViperConfig() {
	viperConfig.viper.AddConfigPath("./config")
	viperConfig.viper.SetConfigName("config.dev")
	if err := viperConfig.viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

// GetDB gets instance of DB
func GetDB() *gorm.DB {
	once.Do(func() { db = InitDB() })
	return db
}

// InitDB initializes DB
func InitDB() *gorm.DB {
	// Use viper to read config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// Get database connection details from configuration
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.name")

	// Create the connection string
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	fmt.Println("DSN:", dsn)

	// Open the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
