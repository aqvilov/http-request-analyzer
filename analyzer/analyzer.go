package analyzer

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Skip(r *http.Request) bool {
	if r.URL.Path == "/favicon.ico" {
		return true
	}
	if strings.HasSuffix(r.URL.Path, ".png") ||
		strings.HasSuffix(r.URL.Path, ".js") ||
		strings.HasSuffix(r.URL.Path, ".css") ||
		strings.HasSuffix(r.URL.Path, ".jpg") {
		return true
	}
	return false
}

func PrintRequest(r *http.Request) {
	if Skip(r) {
		return
	} else {
		fmt.Println("────────────────────────────────────────")

		fmt.Println("Time: ", time.Now().Format("Mon Jan 2 15:04:05 2006 Year"))

		fmt.Println("Method: ", r.Method)
		fmt.Println("URL: ", r.URL.String())
		fmt.Println("Url Path: ", r.URL.Path)
		fmt.Println("Host: ", r.Host)
		fmt.Println("Raw query: ", r.URL.RawQuery)
		fmt.Println("RemoteAddr: ", r.RemoteAddr)
		fmt.Println("Status Code: ", r.Response.StatusCode)

		fmt.Println("Header: ")
		userAgent := r.Header.Get("User-Agent")
		if userAgent != "" {
			fmt.Printf("User-Agent:  %s\n", userAgent)
		}
		contentType := r.Header.Get("Content-Type")
		if contentType != "" {
			fmt.Printf("Content-Type: %s\n", contentType)
		}
		accept := r.Header.Get("Accept")
		if accept != "" {
			fmt.Printf("Accept: %s\n", accept)
		}

		fmt.Println("Other Headers: ")
		for name, values := range r.Header {
			if name != "User-Agent" && name != "Content-Type" && name != "Accept" {
				fmt.Printf("   %s: %v\n", name, values)
			}
		}

		if len(r.URL.Query()) > 0 {
			fmt.Println("Query parameters: ")
			for keyName, values := range r.URL.Query() { // проход по query-параметрам ( ключ значение (оба string))
				fmt.Printf("%s:  %v\n", keyName, values)
			}
		} else {
			fmt.Println("Doesn`t have Query parametrs")
		}

		fmt.Println("────────────────────────────────────────")
	}

}
