package main

import (
	"dandelion_api/pkg/utils"
	"dandelion_api/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    if envErr := godotenv.Load(); envErr != nil {
        fmt.Println(envErr)
        return
    }

    fmt.Println("listening @", os.Getenv("PORT"))

    http.HandleFunc("/api/session", utils.CorsMiddleware(routes.SessionHandler, "GET, POST"))
    http.HandleFunc("/api/utils", utils.CorsMiddleware(routes.UtilsHandler, "GET"))
    http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}







