package config

import (
	"doollm/logging"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var EnvConfig = envConfigSchema{}

func (s *envConfigSchema) GetDSN() string {
	return dsn
}

var dsn string

func init() {
	envInit()
	dsn = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", EnvConfig.MYSQL_USERNAME, EnvConfig.MYSQL_PASSWORD, EnvConfig.MYSQL_HOST, EnvConfig.MYSQL_PORT, EnvConfig.MYSQL_DB_NAME)
}

var defaultConfig = envConfigSchema{
	ENV: "dev",

	MYSQL_HOST:     "127.0.0.1",
	MYSQL_PORT:     "50497",
	MYSQL_USERNAME: "dootask",
	MYSQL_PASSWORD: "123456",
	MYSQL_DB_NAME:  "dootask",
	// MYSQL_HOST:            "127.0.0.1",
	// MYSQL_PORT:            "18888",
	// MYSQL_USERNAME:        "devlop",
	// MYSQL_PASSWORD:        "123456",
	// MYSQL_DB_NAME:         "devlop",
	MAX_REQUEST_BODY_SIZE: 200 * 1024 * 1024,

	PUBLIC_PATH: "/Users/mac-47/Desktop/zeniein/devlop/dootask/public",

	LLM_SERVER_URL: "http://localhost:3001/api",
	LLM_TOKEN:      "xxxxxxxxxxxxx",
}

type envConfigSchema struct {
	ENV string `env:"ENV,DREAM_ENV"`

	MYSQL_HOST     string
	MYSQL_PORT     string
	MYSQL_USERNAME string
	MYSQL_PASSWORD string
	MYSQL_DB_NAME  string

	MAX_REQUEST_BODY_SIZE int

	PUBLIC_PATH string

	LLM_SERVER_URL string
	LLM_TOKEN      string
}

func (s *envConfigSchema) IsDev() bool {
	return s.ENV == "dev" || s.ENV == "TESTING"
}

// envInit Reads .env as environment variables and fill corresponding fields into EnvConfig.
// To use a value from EnvConfig , simply call EnvConfig.FIELD like EnvConfig.ENV
// Note: Please keep Env as the first field of envConfigSchema for better logging.
func envInit() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file, ignored")
	}
	v := reflect.ValueOf(defaultConfig)
	typeOfV := v.Type()

	for i := 0; i < v.NumField(); i++ {
		envNameAlt := make([]string, 0)
		fieldName := typeOfV.Field(i).Name
		fieldType := typeOfV.Field(i).Type
		fieldValue := v.Field(i).Interface()

		envNameAlt = append(envNameAlt, fieldName)
		if fieldTag, ok := typeOfV.Field(i).Tag.Lookup("env"); ok && len(fieldTag) > 0 {
			tags := strings.Split(fieldTag, ",")
			envNameAlt = append(envNameAlt, tags...)
		}

		switch fieldType {
		case reflect.TypeOf(0):
			{
				configDefaultValue, ok := fieldValue.(int)
				if !ok {
					logging.Logger.WithFields(logrus.Fields{
						"field": fieldName,
						"type":  "int",
						"value": fieldValue,
						"env":   envNameAlt,
					}).Warningf("Failed to parse default value")
					continue
				}
				envValue := resolveEnv(envNameAlt, fmt.Sprintf("%d", configDefaultValue))
				if EnvConfig.IsDev() {
					fmt.Printf("Reading field[ %s ] default: %v env: %s\n", fieldName, configDefaultValue, envValue)
				}
				if len(envValue) > 0 {
					envValueInteger, err := strconv.ParseInt(envValue, 10, 64)
					if err != nil {
						logging.Logger.WithFields(logrus.Fields{
							"field": fieldName,
							"type":  "int",
							"value": fieldValue,
							"env":   envNameAlt,
						}).Warningf("Failed to parse env value, ignored")
						continue
					}
					reflect.ValueOf(&EnvConfig).Elem().Field(i).SetInt(envValueInteger)
				}
				continue
			}
		case reflect.TypeOf(""):
			{
				configDefaultValue, ok := fieldValue.(string)
				if !ok {
					logging.Logger.WithFields(logrus.Fields{
						"field": fieldName,
						"type":  "int",
						"value": fieldValue,
						"env":   envNameAlt,
					}).Warningf("Failed to parse default value")
					continue
				}
				envValue := resolveEnv(envNameAlt, configDefaultValue)

				if EnvConfig.IsDev() {
					fmt.Printf("Reading field[ %s ] default: %v env: %s\n", fieldName, configDefaultValue, envValue)
				}
				if len(envValue) > 0 {
					reflect.ValueOf(&EnvConfig).Elem().Field(i).SetString(envValue)
				}
			}
		}

	}
}

func resolveEnv(configKeys []string, defaultValue string) string {
	for _, item := range configKeys {
		envValue := os.Getenv(item)
		if envValue != "" {
			return envValue
		}
	}
	return defaultValue
}
