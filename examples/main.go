package main

import (
	"fmt"

	"github.com/Catroque/loggo/pkg/logger"
)

func main() {

	fmt.Println("== Begin")

	log := logger.New().SetLevel(logger.InfoLevel)
	log.Info().Stack().Msg("Test simple debug level log.")

	fmt.Println("== End")
}

