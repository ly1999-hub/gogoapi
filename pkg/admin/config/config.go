package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var env ENV

// Init read ENV config
func Init() {
	var (
		ctx = context.Background()
	)
	envFile := ".env"
	if err := godotenv.Load(envFile); err != nil {
		fmt.Println("Load env file err: ", err)
	}
	if err := envconfig.Process(ctx, &env); err != nil {
		fmt.Println("Assign env err: ", err)
	}
}

// GetENV ...
func GetENV() *ENV {
	return &env
}
