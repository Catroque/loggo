package main

import (
	"fmt"

	logger "github.com/Catroque/loggo/pkg/logger"
)

func main() {

	fmt.Println("== Begin")

	log := logger.New()
	log.Debug().Msg("Test simple debug level log.")

	fmt.Println("== End")
}

