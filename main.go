package main

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// Handlers check the requested URL path before processing the request. If the URL path does not match the expected path
// ("/" for the home handler and "/ascii-art" for the ASCII art handler), we redirect to the notFoundHandler.
func main() {
	fs := http.FileServer(http.Dir("templates"))              // Create a file server handler for serving static files
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // Register the file server handler for the "/static/" route

	// Registers the homeHandler function as the handler for the root URL ("/") of the server. When a request is made to the root URL, the homeHandler function will be executed.
	http.HandleFunc("/", homeHandler)

	// Registers the asciiArtHandler function as the handler for the "/ascii-art" URL of the server. When a request is made to the "/ascii-art" URL, the asciiArtHandler function will be executed.
	http.HandleFunc("/ascii-art", asciiArtHandler)

	// Register a custom 404 handler for unknown URLs
	http.HandleFunc("/404", notFoundHandler)

	/* Registering a route handler for the /favicon.ico path with the notFoundHandler function.
	The /favicon.ico path is a commonly used path to request the favicon (short for "favorite icon") for a website.
	The favicon is typically a small image file that is displayed in the browser's tab or bookmark bar to represent the website.
	In this case, it means that whenever a request is made for the favicon, the notFoundHandler function will be called.
	This ensures that if a request is made for the /favicon.ico path, the server will respond with a "Page not found" error rather than serving a specific favicon file.
	This approach is commonly used when a website doesn't have a favicon or when you want to handle the request for the favicon in a custom way,
	such as returning an error response instead of serving an actual favicon file. */
	http.HandleFunc("/favicon.ico", notFoundHandler)

	/* This line starts the HTTP server and listens for incoming requests on port 8080. The the 1st argument is the address and
	port to listen on (in this case, ":8080" means to listen on all available network interfaces on port 8080), and the 2nd argument is
	the handler to use for incoming requests (in this case, nil means to use the default handler, which is DefaultServeMux). */
	err := http.ListenAndServe(":8080", nil)
	if err != nil { // Checks if there was an error starting the server. If an error occurs, such as if the port is already in use,
		panic(err) // it will cause the program to panic, meaning it will immediately stop execution and print the error message.
	}
}

/*
Function handles requests to the root path ("/") of the server.
For GET requests, it parses and executes the "templates/index.html" template and sends it as the response.
For any other HTTP method or non-root paths, it calls the notFoundHandler function to return a "Page not found" error response.
*/
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// checks if the requested URL path is not equal to "/". If it's not the root path, the notFoundHandler function is called, which returns
	// a "Page not found" error response. If the requested URL path is indeed "/", the function continues to the next if statement.
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}

	// If the HTTP request method is "GET", the function proceeds to execute the following code block.
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("templates/index.html") // It attempts to parse the template file "templates/index.html" using template.ParseFiles.
		if err != nil {                                          // If there's an error during the parsing,
			http.Error(w, "Template not found", http.StatusNotFound) // it returns a "Template not found" error response with the status code http.StatusNotFound.
			return
		} // If the template parsing is successful, it executes the template using tmpl.Execute. In this case,
		tmpl.Execute(w, nil) // it passes nil as the data parameter, meaning no specific data is provided for the template. The executed template is then written to the response writer, which sends it back as the server's response.
	} else { // If the HTTP request method is not "GET" (e.g., POST, PUT, DELETE),
		notFoundHandler(w, r) // the function calls the notFoundHandler function, which returns a "Page not found" error response.
	}
}

// The asciiArtHandler func handles the POST request when the user submits the form, generates the ASCII art endpoint, and renders the result in the result.html template.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		notFoundHandler(w, r)
		return
	}

	// the code inside the block is executed only when a POST request is received.
	if r.Method == "POST" {
		err := r.ParseForm() // parses the form data from the request. It extracts and populates the r.Form map with the key-value pairs sent in the form data. It also handles the encoding of the form data, whether it's URL-encoded or multipart form data.
		if err != nil {      // If an error occurred while parsing the form data, it means that the form data could not be properly processed, and an HTTP Bad Request (400) error is sent back to the client using http.Error(w, "Failed to parse form", http.StatusBadRequest).
			http.Error(w, "Failed to parse form", http.StatusBadRequest) // The error message "..." is displayed, indicating that there was an issue with the submitted form data.
			return
		}

		// Retrieves the value of the "inputText" field from the parsed form data. The r.Form.Get function looks for the specified key in the r.Form map and returns its corresponding value. In this case, it retrieves the user-provided input text.
		inputText := r.Form.Get("inputText")
		// Retrieves the value of the "banners" field from the parsed form data. Similar to the previous line, it retrieves the selected banner value.
		banner := r.Form.Get("banners")

		// validation check for the input before generating the ASCII art (clicking the button 'ASCII art' without any input)
		if inputText == "" {
			http.Error(w, "Input text is required", http.StatusBadRequest)
			return
		}

		// Sends inputText to isASCII func as an argument, if that function returns true (meaning inputText contains only ASCII characters),
		// then the output of condition "if !isASCII" is an StatusBadRequest error with given message displayed on the webpage.
		if !isASCII(inputText) {
			http.Error(w, "Invalid input text: Non-ASCII characters not allowed", http.StatusBadRequest)
			return
		}

		// when the ASCII art generation encounters an unhandled errors
		art, err := generateASCIIArt(inputText, banner)
		if err != nil {
			http.Error(w, "Failed to generate ASCII art", http.StatusInternalServerError)
			return
		}

		// It uses the template.ParseFiles function to load and parse the "result.html" template file. If an error occurs during parsing, the err variable will capture it.
		tmpl, err := template.ParseFiles("templates/result.html")
		if err != nil { // if there was an error while parsing the template file (the template file could not be found or there was an issue with its format)
			http.Error(w, "Template not found", http.StatusNotFound) // an HTTP Not Found (404) error is sent back to the client and error message is discplaied.
			return
		} // Assuming the template file was successfully parsed, this line executes the template and writes the generated output to the http.ResponseWriter interface w.
		tmpl.Execute(w, art) // The art variable, which contains the generated ASCII art, is passed as the data for rendering the template.
	} else { // If a request method other than POST is received (Since the /ascii-art endpoint is intended for POST requests)
		notFoundHandler(w, r) // this block calls the notFoundHandler function, which returns an HTTP Not Found (404) error to indicate that the requested resource was not found.
	}
}

// to verify if the input text contains only ASCII characters.
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

// Function takes a text string (user input) and a banner string (user choice) as input and returns the generated ASCII art as a string.
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

	// Attempts to open the file specified by the filename variable. If successful, it returns a file descriptor (file) and a nil error.
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close() // schedules the Close() function of the file to be called when the generateASCIIArt completes.

	fileScanner := bufio.NewScanner(file) // creates a scanner (fileScanner) to read the contents of the opened file.
	fileScanner.Split(bufio.ScanLines)    // configures the scanner to split the file content into lines based on newline characters ('\n').
	var fileLines []string                // declares an empty slice (fileLines) to store the lines read from the file.
	for fileScanner.Scan() {              // enters a loop that iterates until the scanner reaches the end of the file. In each iteration,
		fileLines = append(fileLines, fileScanner.Text()) // it reads the next line from the file using fileScanner.Scan() and adds the line to the fileLines slice using fileScanner.Text()
	}

	lines := strings.Split(text, "\n") // Splits the text string into multiple lines based on the newline character ("\n"). Each line of the text will be processed separately.

	var generatedArt strings.Builder // Variable of strings.Builder type is used to efficiently build strings by concatenating various pieces together.

	for _, word := range lines { // The code then enters a loop that iterates over each word in the lines slice.
		if len(word) == 0 { // If the current word is an empty string,
			generatedArt.WriteString("\n") // a newline character is written to the generatedArt builder using generatedArt.WriteString("\n"),
			continue                       // and the loop continues to the next iteration.
		}
		runes := []rune(word)    // Converts the word string into a slice of runes (integer values representing a Unicode code point)
		for k := 1; k < 9; k++ { // Nested loop that iterates over each ch (rune) in the runes slice.
			for _, ch := range runes { // This loop iterates from 1 to 8, representing the rows of each character in the ASCII art .txt files.
				m := rune(k)
				// Calculate the index in the fileLines slice to retrieve the corresponding ASCII art line.
				// Reduce the ASCII value of the character (ch) by 32 (to adjust for the starting ASCII value of the characters),
				asciiFetch := ((ch - 32) * 9) + m                             // multiply it by 9 (the number of rows each character occupies in the ASCII art file), add the row number (m)
				if int(asciiFetch) >= 0 && int(asciiFetch) < len(fileLines) { // checks if the calculated asciiFetch index is within the valid range of the fileLines slice.
					generatedArt.WriteString(fileLines[int(asciiFetch)]) // writes the corresponding ASCII art line from the fileLines slice to the generatedArt builder.
				}
			} // After the nested loops finish,
			generatedArt.WriteString("\n") // a newline character is written to the generatedArt builder, indicating the end of the current line of the ASCII art.
		} // The process repeats for each line of the input text, building the complete ASCII art in the generatedArt builder.
	}
	return generatedArt.String(), nil // The function returns the generated ASCII art as a string using generatedArt.String(), along with a nil error value.
}
