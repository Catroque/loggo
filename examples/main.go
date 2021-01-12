package main

import (
	"fmt"

	"github.com/Catroque/loggo/pkg/logger"
)

func main() {

	fmt.Println("== Begin")

	log := logger.New()
	log.Debug().Stack().Msg("Test simple debug level log.")

	fmt.Println("== End")
}

