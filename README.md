# ascii-art-web
A web application in written in Go that converts text inputs into ASCII art using predefined patterns from external files. Users can customize the output style by selecting from multiple ASCII art templates/fonts. This functionality is accessible through a user-friendly web interface, allowing for dynamic and engaging text-to-art transformations.

## Usage: how to run.
    1. clone the repository: https://github.com/aramsimonyan1/ascii-art-web.git
    2. Start the application: go run main.go
    3. Open your web browser and navigate to http://localhost:8080 to access the application.
    4. Enter your desired text in the input field, select a banner/font style from the dropdown menu, and click the "ASCII art" button to generate the corresponding ASCII graphic reprezentation.
    5. The generated ASCII art will be displayed on the result page.

## Implementation details
###
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