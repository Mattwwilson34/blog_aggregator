package main

import (
	"fmt"

	"github.com/Mattwwilson34/blog_aggregator/internal/config"
)

func main() {
	cfg := config.Config{}
	_, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	err = cfg.SetUser("mattwilson")
	if err != nil {
		fmt.Println(err)
	}
	_, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
}
