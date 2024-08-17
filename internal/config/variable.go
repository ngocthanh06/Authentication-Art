package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

var TokenType = "Bearer"
var Env = ".env"

type envKey struct {
	JwtKey []byte
	TTL    time.Duration
}

var EnvKey *envKey

func InitEnvKey() {
	ttl, err := strconv.ParseInt(os.Getenv("TTL"), 10, 64)

	if err != nil {
		log.Fatal(err)

		return
	}

	EnvKey = &envKey{
		JwtKey: []byte(os.Getenv("JWT_KEY")),
		TTL:    time.Duration(ttl),
	}
}
