package config

import (
	"os"

	"feldrise.com/animal-api/database"
	"feldrise.com/animal-api/database/dbmodel"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func doesEnvExists(name string) bool {
	_, exists := os.LookupEnv(name)

	return exists
}

type Constants struct {
	// Constants
	Port           string `yaml:"port"`
	JWTSecret      string `yaml:"jwtSecret"`
	DataPath       string `yaml:"dataPath"`
	BaseURL        string `yaml:"baseURL"`
	ApplicationURL string `yaml:"applicationURL"`

	// Database
	ConnectionString string `yaml:"connectionString"`
}

type Config struct {
	Constants

	// Repositories
	CatsRepository   dbmodel.CatsRepository
	UserRepository   dbmodel.UserRepository
	VisitsRepository dbmodel.VisitsRepository
}

func initViper(configName string) (Constants, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok && err != nil {
		return Constants{}, err
	}

	// At this point the only error would be a missing file
	if err != nil {
		err = initViperEnv()

		if err != nil {
			return Constants{}, err
		}
	}

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}

func initViperEnv() error {
	if !doesEnvExists("JWT_SECRET") ||
		!doesEnvExists("DATA_PATH") ||
		!doesEnvExists("BASE_URL") ||
		!doesEnvExists("APPLICATION_URL") ||
		!doesEnvExists("CONNECTION_STRING") {
		return &MissingEnvVariableError{}
	}

	viper.SetDefault("Port", "8080")
	viper.SetDefault("JWTSecret", os.Getenv("JWT_SECRET"))
	viper.SetDefault("DataPath", os.Getenv("DATA_PATH"))
	viper.SetDefault("BaseURL", os.Getenv("BASE_URL"))
	viper.SetDefault("ApplicationURL", os.Getenv("APPLICATION_URL"))
	viper.SetDefault("ConnectionString", os.Getenv("CONNECTION_STRING"))

	return nil
}

func New() (*Config, error) {
	config := Config{}

	// Constants
	constants, err := initViper("config")

	config.Constants = constants
	if err != nil {
		return &config, err
	}

	// Database
	databaseSession, err := gorm.Open(postgres.Open(config.Constants.ConnectionString), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)
	database.ApplySeeds(databaseSession)

	config.CatsRepository = dbmodel.NewCatsRepository(databaseSession)
	config.UserRepository = dbmodel.NewUserRepository(databaseSession)
	config.VisitsRepository = dbmodel.NewVisitsRepository(databaseSession)

	return &config, nil
}
