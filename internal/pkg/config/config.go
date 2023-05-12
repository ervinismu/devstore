package config

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver                string        `mapstructure:"DB_DRIVER"`
	DBConnection            string        `mapstructure:"DB_CONNECTION"`
	ServerPort              string        `mapstructure:"SERVER_PORT"`
	LogLevel                string        `mapstructure:"LOG_LEVEL"`
	AccessTokenKey          string        `mapstructure:"ACCESS_TOKEN_KEY"`
	RefreshTokenKey         string        `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration    time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	CloudinaryCloudName     string        `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	CloudinaryApiKey        string        `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinaryApiSecret     string        `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudinaryUploadFolder  string        `mapstructure:"CLOUDINARY_UPLOAD_FOLDER"`
	PaginateDefaultPage     int           `mapstructure:"PAGINATE_DEFAULT_PAGE"`
	PaginateDefaultPageSize int           `mapstructure:"PAGINATE_DEFAULT_PAGE_SIZE"`
	MidtransServerKey       string        `mapstructure:"MIDTRANS_SERVER_KEY"`
	MidtransClientKey       string        `mapstructure:"MIDTRANS_CLIENT_KEY"`
	MidtransMerchantID      string        `mapstructure:"MIDTRANS_MERCHANT_ID"`
	MidtransBaseURL         string        `mapstructure:"MIDTRANS_BASE_URL"`
}

func LoadConfig(fileConfigPath string) (Config, error) {
	var config Config

	viper.AddConfigPath(fileConfigPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Error(fmt.Errorf("error LoadConfig : %w", err))
	}
	return config, nil
}
