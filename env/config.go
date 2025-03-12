package env

// NOTE: Only the Config is modified to include the new environment variables
// Config represents the configuration for environment variables

var EnvConfig = map[string]EnvVariable{
	"PORT": {
		Type:     "int",
		Required: true,
		Default:  8080,
	},
	"NODE_ENV": {
		Type:     "string",
		Required: false,
		Default:  "development",
	},
}

// LoadEnv loads the environment variables
// NOTE: This function must be called before using the environment variables
// DO NOT MODIFY BELOW THIS LINE
func LoadEnv() {
	VerifyEnv(EnvConfig)
}

var Env = UTILS_INTERNAL_ENV
