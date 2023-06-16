package main

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		inputText := r.Form.Get("inputText")
		banner := r.Form.Get("banners")

		art, err := generateASCIIArt(inputText, banner)
		if err != nil {
			http.Error(w, "Failed to generate ASCII art", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("result.html"))
		tmpl.Execute(w, art)
	}
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

	// looking for "\n" and turn it into "n3wL1ne" so string.Split can find it
	preLine := []rune(text)
	for m := 0; m < len(preLine); m++ {
		arrayMiddle := "n3wL!Ne"
		if preLine[m] == 92 && preLine[m+1] == 'n' {
			array1 := preLine[0:m]
			array2 := preLine[m+2:]
			s1 := string([]rune(array1))
			s2 := string([]rune(array2))
			text = s1 + arrayMiddle + s2
			preLine = ([]rune(text))
		}
	}

	// split the text into lines if required
	lines := strings.Split(string(preLine), "n3wL!Ne")

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
				// reduce each character value by 32 in ascii table,
				// multiply by the 9 rows each character uses in standard.txt,
				// add the row number
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
