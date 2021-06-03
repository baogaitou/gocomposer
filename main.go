package main

import (
	"log"
	"net/url"
	"os"

	"github.com/alexsasharegan/dotenv"
)

// "github.com/davecgh/go-spew/spew"

func main() {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 检查env文件中的参数
	u, err := url.ParseRequestURI(os.Getenv("domain"))
	if err != nil {
		log.Fatal("Invalid domain in .env: ", os.Getenv("domain"))
	}

	log.Println("Start service @ ", u)
	log.Println("Origin site:", os.Getenv("mirror"))

	// 启动服务器
	router := InitRouter()
	ginError := router.Run(":2048")
	if ginError != nil {
		log.Fatal("Start web service fail.")
	}
}
