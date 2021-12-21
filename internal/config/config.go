package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Postgres   PostgresConfig
		Redis      RedisConfig
		GRPC       GRPCConfig
		ClickHouse ClickHouseConfig
		Kafka      KafkaConfig
	}

	PostgresConfig struct {
		Host     string
		Port     string
		Username string
		DBName   string
		SSLMode  string
		Password string
	}

	ClickHouseConfig struct {
		Host  string
		Port  string
		Debug string
	}

	RedisConfig struct {
		Addr string
	}

	GRPCConfig struct {
		Server string
		Client string
	}

	KafkaConfig struct {
		Net       string
		Addr      string
		Topic     string
		Partition int
	}
)

func Init(configsDir string) (*Config, error) {

	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	cfg.Postgres.Password = viper.GetString("postgres_pass")
	return nil
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("redis", &cfg.Redis); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("grpc", &cfg.GRPC); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("clickHouse", &cfg.ClickHouse); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("kafka", &cfg.Kafka); err != nil {
		return err
	}
	return nil
}
