package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	ENV_KEY_PORT         string = "PORT"
	ENV_KEY_JWT_SECRET   string = "JWT_SECRET"
	ENV_KEY_JWT_DURATION string = "JWT_DURATION"
)

type envVar struct {
	PORT         string
	JWT_SECRET   string
	JWT_DURATION int64
}

func GetEnv() *envVar {
	config := &envVar{
		JWT_SECRET: os.Getenv(ENV_KEY_JWT_SECRET),
	}

	// parsing jwt_duration
	jwt_duration, err := strconv.Atoi(os.Getenv(ENV_KEY_JWT_DURATION))
	if err != nil {
		jwt_duration = 0
	}
	config.JWT_DURATION = int64(jwt_duration)

	var port string = os.Getenv(ENV_KEY_PORT)

	// setting port to 5000 if env PORT not provided
	if port == "" {
		port = ":5000"
	} else {
		port = ":" + port
	}

	// setting port value
	config.PORT = port

	return config
}

func UpdateEnv(key, value string) {
	// Specify the path to your .env file
	envFilePath := ".env"

	// Read the contents of the .env file
	envFileContents, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		fmt.Println("Error reading .env file:", err)
		return
	}

	// Split the file contents into lines
	lines := strings.Split(string(envFileContents), "\n")

	// Create a map to hold the key-value pairs
	envVars := make(map[string]string)

	// Parse the existing key-value pairs
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			envVars[parts[0]] = parts[1]
		}
	}

	// Update the variable with the new value
	envVars[key] = value

	// Serialize the updated key-value pairs back to a string
	var updatedEnvFileContents string
	for key, value := range envVars {
		updatedEnvFileContents += key + "=" + value + "\n"
	}

	// Write the updated contents back to the .env file
	if err := ioutil.WriteFile(envFilePath, []byte(updatedEnvFileContents), 0644); err != nil {
		fmt.Println("Error writing .env file:", err)
		return
	}

	fmt.Printf("Updated %s to %s in .env\n", key, value)

}
