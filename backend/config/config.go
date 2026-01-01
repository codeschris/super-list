package config

import (
    "log"
    "os"
    
    "github.com/joho/godotenv"
)

type Config struct {
    DatabaseURL      string
    SupabaseJWTSecret string
    Port             string
    GinMode          string
}

func Load() *Config {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
    
    cfg := &Config{
        DatabaseURL:      getEnv("DATABASE_URL", ""),
        SupabaseJWTSecret: getEnv("SUPABASE_JWT_SECRET", ""),
        Port:             getEnv("PORT", "8080"),
        GinMode:          getEnv("GIN_MODE", "debug"),
    }
    
    if cfg.DatabaseURL == "" {
        log.Fatal("DATABASE_URL is required")
    }
    if cfg.SupabaseJWTSecret == "" {
        log.Fatal("SUPABASE_JWT_SECRET is required")
    }
    
    return cfg
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}