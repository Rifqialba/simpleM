package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
	Env  string
}

type DatabaseConfig struct {
	Host string
	Port string
	User string

	Password string
	Name string
	SSLMode string
}

type RedisConfig struct {
	Host string
	Port string
}

type Config struct {
	App AppConfig

	Database DatabaseConfig

	Redis RedisConfig

	WebhookSecret string

	JWTSecret string
}

func Load() *Config {

	_ = godotenv.Load()

	return &Config{
		App: AppConfig{
			Port: getEnv("APP_PORT", "8080"),

			Env: getEnv(
				"APP_ENV",
				"development",
			),
		},

		Database: DatabaseConfig{
			Host: getEnv(
				"POSTGRES_HOST",
				"localhost",
			),

			Port: getEnv(
				"POSTGRES_PORT",
				"5433",
			),

			User: getEnv(
				"POSTGRES_USER",
				"simplem",
			),

			Password: getEnv(
				"POSTGRES_PASSWORD",
				"simplem_password",
			),

			Name: getEnv(
				"POSTGRES_DB",
				"simplem_db",
			),

			SSLMode: getEnv(
				"POSTGRES_SSLMODE",
				"disable",
			),
		},

		Redis: RedisConfig{
			Host: getEnv(
				"REDIS_HOST",
				"localhost",
			),

			Port: getEnv(
				"REDIS_PORT",
				"6380",
			),
		},

		WebhookSecret: getEnv(
			"WEBHOOK_SECRET",
			"",
		),

		JWTSecret: getEnv(
			"JWT_SECRET",
			"",
		),
	}
}

func (c *Config) DatabaseURL() string {

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

func getEnv(
	key string,
	fallback string,
) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}