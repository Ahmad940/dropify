package config

import (
	"os"
	"strconv"
)

type envVar struct {
	PORT         string
	JWT_SECRET   string
	JWT_DURATION int64
	// DATABASE_URL string
}

func GetEnv() *envVar {
	config := &envVar{
		// DATABASE_URL: os.Getenv("DATABASE_URL"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

	// parsing jwt_duration
	jwt_duration, err := strconv.Atoi(os.Getenv("JWT_SECRET"))
	if err != nil {
		jwt_duration = 0
	}
	config.JWT_DURATION = int64(jwt_duration)

	var port string = os.Getenv("PORT")

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
