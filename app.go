package main

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query()) == 0 {
			data := map[string]string{
				"Region":   os.Getenv("FLY_REGION"),
			}
			if err := t.ExecuteTemplate(w, "index.html.tmpl", data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			dividend, divider, err := handleParams(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			quotient, err := divide(dividend, divider)
			if err != nil {
				http.Error(w, fmt.Errorf("error during dividing: %w", err).Error(), http.StatusBadRequest)
				return
			}
			data := map[string]string{
				"Region":   os.Getenv("FLY_REGION"),
				"Quotient": fmt.Sprintf("%d", quotient),
			}
			if err := t.ExecuteTemplate(w, "index.html.tmpl", data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})
	log.Println("listening on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("Error starting server:", err)
		os.Exit(1)
	}
}

func handleParams(r *http.Request) (int, int, error) {
	dividendStr := r.URL.Query().Get("dividend")
	if len(dividendStr) > 20 {
		return 0, 0, errors.New("we don't do dividends of such length here")
	}
	dividend, err := strconv.Atoi(dividendStr)
	if err != nil {
		return 0, 0, fmt.Errorf("dividend has to be a valid int: %w", err)
	}

	dividerStr := r.URL.Query().Get("divider")
	if len(dividerStr) > 20 {
		return 0, 0, errors.New("we don't do dividers of such length here")
	}
	divider, err := strconv.Atoi(dividerStr)
	if err != nil {
		return 0, 0, fmt.Errorf("divider has to be a valid int: %w", err)
	}
	return dividend, divider, nil
}

func divide(dividend, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return dividend / divider, nil
}
