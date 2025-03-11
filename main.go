package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lionpuro/portfolio/views"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", indexHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	views.Base().Render(r.Context(), w)
}
