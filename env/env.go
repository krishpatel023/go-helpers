package env

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

// Variable represents a configuration for an environment variable
type EnvVariable struct {
	Type     string // "string", "int", "bool", etc.
	Required bool
	Default  any
}

// Global configuration variables
var UTILS_INTERNAL_ENV = make(map[string]any)

// Convert the env variable to the correct type
func convertEnv(key string, env string, Type string) any {
	switch Type {
	case "string":
		return env
	case "int":
		i, err := strconv.Atoi(env)
		if err != nil {
			fmt.Printf("Error: %s is not an integer\n", key)
			return 0
		}
		return i
	case "bool":
		b, err := strconv.ParseBool(env)
		if err != nil {
			fmt.Printf("Error: %s is not a boolean\n", key)
			return false
		}
		return b
	case "float":
		f, err := strconv.ParseFloat(env, 64)
		if err != nil {
			fmt.Printf("Error: %s is not a float\n", key)
			return 0.0
		}
		return f
	case "url":
		u, err := url.Parse(env)
		if err != nil {
			fmt.Printf("Error: %s is not a URL\n", key)
			return ""
		}
		return u.String()
	default:
		return env
	}
}

// Check if the environment variable is of the correct type
func checkEnvType(key string, env any, Type string) bool {
	if env == "" {
		return false
	}
	EnvType := reflect.TypeOf(env)

	switch Type {
	case "string":
		if EnvType != reflect.TypeOf("") {
			fmt.Printf("Error: %s is not a string\n", key)
			return false
		}
	case "int":
		if EnvType != reflect.TypeOf(0) {
			fmt.Printf("Error: %s is not an integer\n", key)
			return false
		}
	case "bool":
		if EnvType != reflect.TypeOf(true) {
			fmt.Printf("Error: %s is not a boolean\n", key)
			return false
		}
	case "float":
		if EnvType != reflect.TypeOf(0.0) {
			fmt.Printf("Error: %s is not a float\n", key)
			return false
		}
	case "url":
		if EnvType != reflect.TypeOf("") {
			fmt.Printf("Error: %s is not a URL\n", key)
			return false
		} else {
			_, err := url.Parse(env.(string))
			if err != nil {
				fmt.Printf("Error: %s is not a URL\n", key)
				return false
			}
		}
	default:
		fmt.Printf("Error: %s is not a valid type\n", Type)
		return false
	}

	return true
}

// SetEnv sets the environment variable
func SetEnv(key string, value any) {
	UTILS_INTERNAL_ENV[key] = value
}

// VerifyEnv checks if the required environment variables are set and valid
func VerifyEnv(EnvConfig map[string]EnvVariable) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	isError := false

	for key, variable := range EnvConfig {
		// Get the environment variable
		env := os.Getenv(key)

		// Check if required environment variables are set
		if variable.Required && env == "" {
			fmt.Printf("Error: %s is required but not set in .env\n", key)
			isError = true
			continue
		}

		if env != "" {

			// Convert the environment variable to the correct type
			convertedEnv := convertEnv(key, env, variable.Type)

			// Check if the environment variable is of the correct type
			res := checkEnvType(key, convertedEnv, variable.Type)
			if !res {
				isError = true
				continue
			}

			// Set the environment variable
			SetEnv(key, convertedEnv)
		} else {
			// Set the default value if the environment variable is not set
			SetEnv(key, variable.Default)
		}
	}

	if isError {
		fmt.Println("Error: Environment variables are not set correctly")
		os.Exit(1)
	}
}
