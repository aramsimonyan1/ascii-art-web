package main

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	// Register a custom 404 handler for unknown URLs
	http.HandleFunc("/404", notFoundHandler)
	http.HandleFunc("/favicon.ico", notFoundHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusNotFound)
			return
		}
		tmpl.Execute(w, nil)
	} else {
		notFoundHandler(w, r)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		notFoundHandler(w, r)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		inputText := r.Form.Get("inputText")
		banner := r.Form.Get("banners")

		if inputText == "" {
			http.Error(w, "Input text is required", http.StatusBadRequest)
			return
		}

		// new
		if !isASCII(inputText) {
			http.Error(w, "Invalid input text: Non-ASCII characters not allowed", http.StatusBadRequest)
			return
		}

		art, err := generateASCIIArt(inputText, banner)
		if err != nil {
			http.Error(w, "Failed to generate ASCII art", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/result.html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusNotFound)
			return
		}
		tmpl.Execute(w, art)
	} else {
		notFoundHandler(w, r)
	}
}

func isASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}

func generateASCIIArt(text, banner string) (string, error) {
	var filename string
	switch banner {
	case "shadow":
		filename = "shadow.txt"
	case "standard":
		filename = "standard.txt"
	case "thinkertoy":
		filename = "thinkertoy.txt"
	default:
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	lines := strings.Split(text, "\n")

	var generatedArt strings.Builder

	for _, word := range lines {
		if len(word) == 0 {
			generatedArt.WriteString("\n")
			continue
		}

		runes := []rune(word)
		for k := 1; k < 9; k++ {
			for _, ch := range runes {
				m := rune(k)
				asciiFetch := ((ch - 32) * 9) + m
				if int(asciiFetch) >= 0 && int(asciiFetch) < len(fileLines) {
					generatedArt.WriteString(fileLines[int(asciiFetch)])
				}
			}
			generatedArt.WriteString("\n")
		}
	}
	return generatedArt.String(), nil
}
