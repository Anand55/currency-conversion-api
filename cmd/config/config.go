package config

import "os"

var (
	REDIS_ADDR = os.Getenv("REDIS_ADDR")
	FIXER_KEY  = os.Getenv("FIXER_KEY")
)
