package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"log"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/valentyna-koshelnyk/panda-eats-prototype-api/internal/domain/entity"
)

var (
	db       *gorm.DB
	once     sync.Once
	dynamoDB *dynamodb.Client
)

type Resolver struct {
	URL *url.URL
}

// InitViperConfig Initialize viper and load configuration
func InitViperConfig() {
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName("config.dev")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}

// GetDB Initialize and return the database instance
func GetDB() *gorm.DB {
	once.Do(func() {
		db = initDB()
	})
	return db
}

// initDB private function to initialize the database
func initDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = db.AutoMigrate(&entity.Menu{}, &entity.Restaurant{}, &entity.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

func InitDynamoSession() dynamo.Table {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Endpoint: aws.String("http://localhost:4566")},
	}))
	db := dynamo.New(sess, &aws.Config{Region: aws.String("eu-central-1")})

	table := db.Table("cart")
	return table
}
