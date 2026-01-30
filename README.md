# HTTP Request Analyzer

Go package for analyzing HTTP requests.

## Installation

Write the next text in concole 
```
go get github.com/aqvilov/http-request-analyzer
```


## USAGE
```import "github.com/aqvilov/http-request-analyzer"```

#### SOURCE CODE IN `ANALYZER` DIRECTORY


## What does the function output from the package:

1. Method
2. URL
3. URL Path
4. HOST
5. Raw query
6. RemoteAddr
7.Status code
8.User-Agent
9.Headers 



### Example usage
```
func main() {
    // your code before ( with all handlefunc and added r *http.Request )
    analyzer.PrintRequest(r)
}
```


An example of usage is shown in example.go

Aqvilov.


