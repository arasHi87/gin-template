package setting

import "os"

/*
Service configuration
*/

type Config struct {
	Port       string
	Address    string
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
	DBTimezone string
}

/*
Get environment key as string.

Default value(defaultVal) will be applied
if the environment key's value is missing or can't convert to desired type.
*/
func getEnvStr(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

var CONFIG = Config{
	Port:       getEnvStr("APP_PORT", "8080"),
	Address:    getEnvStr("APP_ADDRESS", "0.0.0.0"),
	DBHost:     getEnvStr("DB_HOST", "localhost"),
	DBPort:     getEnvStr("DB_PORT", "5432"),
	DBName:     getEnvStr("DB_NAME", "gintemplate"),
	DBUsername: getEnvStr("DB_USERNAME", "m3ow87"),
	DBPassword: getEnvStr("DB_PASSWORD", "m3ow87"),
	DBTimezone: getEnvStr("DB_TIMEZONE", "Asia/Taipei"),
}
