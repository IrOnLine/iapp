package config

import "os"

// Load config from environment variables

var (
  JWTSecret = os.Getenv("JWT_SECRET")
  DBUser = os.Getenv("DB_USER")
  DBPass = os.Getenv("DB_PASSWORD")
  DBHost = os.Getenv("DB_HOST")
  DBName = os.Getenv("DB_NAME")
)