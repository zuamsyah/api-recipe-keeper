package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var Default = map[string]Config{
	"APP_URL":               "https://kurir.zahironline.com",
	"APP_ENV":               "production",
	"APP_PORT":              "4000",
	"DB_DRIVER":             "mysql",
	"DB_HOST":               "127.0.0.1",
	"DB_PORT":               "3306",
	"DB_NAME":               "kurir",
	"DB_USER":               "root",
	"DB_PASSWORD":           "",
	"DB_MAX_OPEN_CONNS":     "100",
	"DB_MAX_IDLE_CONNS":     "2",
	"DB_CONN_MAX_LIFETIME":  "0ms",
	"DB_IS_DEBUG":           "false",
	"REDIS_HOST":            "127.0.0.1",
	"REDIS_PORT":            "6379",
	"REDIS_PASSWORD":        "",
	"REDIS_DB":              "0",
	"API_IS_DEBUG":          "false",
	"ZAHIR_ID_URL":          "https://zahir-id.zahironline.com",
	"ZAHIR_ID_IS_SKIP_AUTH": "false",
	"SLACK_ERROR_URL":       "https://hooks.slack.com/services/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",
	"GOSEND_URL":            "https://integration-kilat-api.gojekapi.com",
	"GOSEND_CLIENT_ID":      "zahir-tokoku-engine",
	"GOSEND_PASS_KEY":       "9a1a0fd5ab32474011d1f5d26c339b21b580829a6ed9046efc567afd0603cb84",
	"RAJAONGKIR_URL":        "https://pro.rajaongkir.com",
	"RAJAONGKIR_KEY":        "481c946ef3707d9e5776d7b5e60b747a",
	"LODI_URL":              "https://dev-api-v1.lodi.id",
	"LODI_KEY":              "65d86451998bf564877d9ba79673b298e53e42cc46e5caf7b00e7bdbcb6d1cf03d3137c6615a46957f73cb07376fd3a9757949cff2b518ba0efcb8bb2ed67064OyryegQ+AZ86XTBjQRA41seOEx9SjnKUnGbpoGU+pbM3IK8Owqrp7SeXrVE8uJX6XSxMbtAi4z45ts0Ig1itgg==",
	"LODI_TRANSPORT_URL":    "https://transport-api-dev.lodi.id",
	"LODI_USERNAME":         "zahir",
	"LODI_PASSWORD":         "12345678",
}

type Config string

func init() {
	godotenv.Load()
}

func Get(key string) Config {
	value := Config(os.Getenv(key))
	if value == "" {
		value = Default[key]
	}
	return value
}

func (c Config) String() string {
	return string(c)
}

func (c Config) Int() int {
	v, err := strconv.Atoi(c.String())
	if err != nil {
		return 0
	}
	return v
}

func (c Config) Bool() bool {
	if strings.ToLower(c.String()) == "true" {
		return true
	}
	return false
}

func (c Config) Duration() time.Duration {
	v, err := time.ParseDuration(c.String())
	if err != nil {
		return 0
	}
	return v
}
