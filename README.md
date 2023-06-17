# Description

## Objectives
Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, ascii-art.

### Your webpage must allow the use of the different banners:
    shadow
    standard
    thinkertoy

### Implement the following HTTP endpoints:
    GET /: Sends HTML response, the main page.
    1.1. GET Tip: go templates to receive and display data from the server.
    POST /ascii-art: that sends data to Go server (text and a banner)
    2.1. POST Tip: use form and other types of tags to make the post request.\

### Displaying the result from the POST. 
    What we recommend are one of the following:
    Display the result in the route /ascii-art after the POST is completed. So going from the home page to another page.
    Or display the result of the POST in the home page. This way appending the results in the home page.

### The main page must have:
    text input
    radio buttons, select object or anything else to switch between banners
    button, which sends a POST request to '/ascii-art' and outputs the result on the page.

## HTTP status code
### Your endpoints must return appropriate HTTP status codes.
    OK (200), if everything went without errors.
    Not Found, if nothing is found, for example templates or banners.
    Bad Request, for incorrect requests.
    Internal Server Error, for unhandled errors.
### In the root project directory create a README.MD file with the following sections and contents:
    Description
    Authors
    Usage: how to run
    Implementation details: algorithm

## Instructions
    HTTP server must be written in Go.
    HTML templates must be in the project root directory templates.
    The code must respect the good practices.

## Allowed packages
    Only the standard go packages are allowed

## Usage
Here's an example: http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20

## Your group
You need at least to be 2 members to start the project, but you can be up to 3. Minimum audit ratio of 0.5 is required for every member of the group.



# Authors
Aram Simonyan



# Usage: how to run.
    1. clone the repository: https://learn.01founders.co/git/asimonya/ascii-art-web.git
    2. Start the application: go run main.go
    3. Open your web browser and visit http://localhost:8080 to access the application.
    4. Enter your desired text in the input field, select a banner style from the dropdown menu, and click the "ASCII art" button to generate the corresponding ASCII art.
    5. The generated ASCII art will be displayed on the result page.



# Implementation details: algorithm
The ASCII-Art-Web is implemented using Go programming language. The application utilizes the Go standard library's net/http package for handling HTTP requests and responses. Here's a brief overview of the implementation details:

    The application follows a client-server architecture, where the client is a web browser and the server is implemented using Go's net/http package.
    The server serves static files (HTML templates and CSS) from the templates directory using the http.FileServer and http.StripPrefix handlers.
    The server defines two HTTP request handlers: homeHandler and asciiArtHandler.
        The homeHandler handles the GET request for the main page, where users can input their text and select the banner style. It renders the index.html template.
        The asciiArtHandler handles the POST request when the user submits the form. It generates the ASCII art based on the user input and selected banner style, and renders the result in the result.html template.
    The ASCII art generation is performed by the generateASCIIArt function, which converts the user input into ASCII art using the selected banner style.
    The application handles various HTTP status codes to provide appropriate error responses:
        HTTP 200 (OK): Returned when everything goes without errors.
        HTTP 400 (Bad Request): Returned for incorrect requests or failed form parsing.
        HTTP 404 (Not Found): Returned when templates or banners are not found.
        HTTP 500 (Internal Server Error): Returned for unhandled errors during ASCII art generation.