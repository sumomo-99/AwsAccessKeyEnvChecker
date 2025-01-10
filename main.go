package main

import (
	"fmt"
	"os"
)

func isEnvExist(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return "not set"
	}
	return value
}

func main() {
	var envs = []string{"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY"}

	for _, env := range envs {
		value := isEnvExist(env)
		fmt.Printf("%s: %s\n", env, value)
	}
}
