package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/yceman/store-go/internal/config"
)

func main() {
	default_config := &config.Config{}
	if file_config := os.Getenv("STORE_FILE_CONFIG"); file_config != "" {
		file, err := os.ReadFile(file_config)

		if err != nil {
			log.Panicln(err.Error())
		}

		err = json.Unmarshal(file, default_config)
		if err != nil {
			log.Panicln(err.Error())
		}
	}
	conf := config.NewConfig(default_config)
	fmt.Println(conf)
}
