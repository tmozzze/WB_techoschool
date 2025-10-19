package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// PostgreSQL
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Kafka
	KafkaHost          string
	KafkaPort          string
	KafkaTopic         string
	KafkaConsumerGroup string
	KafkaPartitions    int
	KafkaReplication   int

	// Server
	ServerPort string

	// Cache
	CacheCapacity int
}

func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println(".env not found. Skip")
	}

	KafkaPartitions, err := strconv.Atoi(os.Getenv("KAFKA_TOPIC_PARTITIONS"))
	if err != nil {
		err := errors.New("KAFKA_TOPIC_PARTITIONS incorrect")
		return nil, err
	}
	KafkaReplication, err := strconv.Atoi(os.Getenv("KAFKA_TOPIC_REPLICATION"))
	if err != nil {
		err := errors.New("KAFKA_TOPIC_REPLICATION incorrect")
		return nil, err
	}
	CacheCapacity, err := strconv.Atoi(os.Getenv("CACHE_CAPACITY"))
	if err != nil {
		err := errors.New("CACHE_CAPACITY incorrect")
		return nil, err
	}

	cfg := &Config{
		// DB
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		DBName:     os.Getenv("POSTGRES_DB"),

		// KAFKA
		KafkaHost:          os.Getenv("KAFKA_ADVERTISED_HOST"),
		KafkaPort:          os.Getenv("KAFKA_PORT"),
		KafkaTopic:         os.Getenv("KAFKA_TOPIC"),
		KafkaConsumerGroup: os.Getenv("KAFKA_CONSUMER_GROUP"),
		KafkaPartitions:    KafkaPartitions,
		KafkaReplication:   KafkaReplication,

		// CACHE
		CacheCapacity: CacheCapacity,

		// SERVER
		ServerPort: os.Getenv("SERVER_PORT"),
	}

	if cfg.DBUser == "" || cfg.DBPassword == "" {
		err := errors.New("DB_USER or DB_PASSWORD is empty")
		return nil, err
	}

	if cfg.KafkaTopic == "" || cfg.KafkaConsumerGroup == "" {
		err := errors.New("KAFKA_TOPIC or KAFKA_CONSUMER_GROUP is empty")
		return nil, err
	}

	if cfg.ServerPort == "" {
		err := errors.New("SERVER_PORT is empty")
		return nil, err
	}

	return cfg, nil
}
