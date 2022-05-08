package setting

import "os"

/*
Service configuration
*/

type Config struct {
	Port    string
	Address string
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
	Port:    getEnvStr("APP_PORT", "8080"),
	Address: getEnvStr("APP_ADDRESS", "0.0.0.0"),
}
