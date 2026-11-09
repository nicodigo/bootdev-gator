package main

import (
	"fmt"
	"os"

	"github.com/nicodigo/blog-aggregator/internal/config"
)

func main() {
	myConfig, err := config.Read()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	err = myConfig.SetUser("nico")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	myConfig, err = config.Read()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("db_url:", myConfig.DbURL)
	fmt.Println("current_user_name:", myConfig.CurrentUserName)

	os.Exit(0)
}
