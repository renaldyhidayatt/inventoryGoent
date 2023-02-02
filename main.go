package main

import (
	"context"
	"log"

	"github.com/renaldyhidayatt/inventorygoent/utils"
)

func main() {
	ctx := context.Background()

	err := utils.Viper()

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = utils.Database(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}
}
