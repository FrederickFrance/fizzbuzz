package main

import (
	"test.com/fizzbuzz/logger"
	"test.com/fizzbuzz/router"
)

func main() {
	logger.Logger.Info("Launch")

	router.FillRouter()
}
