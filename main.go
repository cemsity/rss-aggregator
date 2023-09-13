package main

import (
	"fmt"
	"log"
	"os"
	genv "github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
)

func main(){
	fmt.Println("HELLO WORLD")
	fmt.Println("TESTING!!!")
	genv.Load(".env")
	
	port := os.Getenv("PORT")

	if port == ""{
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println(port)

	router := chi.NewRouter()
	
}