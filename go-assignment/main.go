package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", usageGuide)
	http.HandleFunc("/convert/", convertHandler)
	http.HandleFunc("/health", healthCheck)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func usageGuide(w http.ResponseWriter, r *http.Request) {
	guide := `
Usage Guide:
GET /convert/<value>/<input-format>/<output-format>

<value>: Any alphanumeric value in the format specified by <input-format>
<input-format> and <output-format>:
  dec: Decimal (base-10) format
  bin: Binary (base-2) format
  hex: Hexadecimal (base-16) format

Example: /convert/1010/bin/dec
`
	fmt.Fprint(w, guide)
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 5 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	value := parts[2]
	inputFormat := parts[3]
	outputFormat := parts[4]

	result, err := convert(value, inputFormat, outputFormat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, result)
}

func convert(value, inputFormat, outputFormat string) (string, error) {
	var num int64
	var err error

	switch inputFormat {
	case "dec":
		num, err = strconv.ParseInt(value, 10, 64)
	case "bin":
		num, err = strconv.ParseInt(value, 2, 64)
	case "hex":
		num, err = strconv.ParseInt(value, 16, 64)
	default:
		return "", fmt.Errorf("invalid input format: %s", inputFormat)
	}

	if err != nil {
		return "", fmt.Errorf("invalid input value: %s", value)
	}

	switch outputFormat {
	case "dec":
		return strconv.FormatInt(num, 10), nil
	case "bin":
		return strconv.FormatInt(num, 2), nil
	case "hex":
		return strconv.FormatInt(num, 16), nil
	default:
		return "", fmt.Errorf("invalid output format: %s", outputFormat)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}