package main

import (
	"fmt"

	"github.com/ademirfteo/gologwrp/pkg/logger"
)

func main() {

	fmt.Println("== Begin")

	log := logger.New()
	log.Debug().Msg("Test simple debug level log.")

	fmt.Println("== End")
}

