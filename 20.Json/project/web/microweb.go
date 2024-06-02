package web

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func RunWeb(address string) {

	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(address, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	// str := "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n<meta charset=\"UTF-8\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n<meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">\n<title>HTML 5 Boilerplate</title>\n<link rel=\"stylesheet\" href=\"style.css\">\n</head>\n<body>\n<script src=\"index.js\"></script>\n</body>\n</html>"
	// fmt.Fprint(w, str)
	// fmt.Fprintf(w, "Welcome,Your Ip is:%s\n", r.RemoteAddr)
	//-------------------------------------------------------------------------------------------
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Fprint(w, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
