package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var mappedArgs = map[string][]string{}
var mappedHeaders = map[string][]string{}

func ensureThatValueNextToOptionIsValid(nextValue string, optionName string) {
	if !strings.Contains(nextValue[0:2], "-") {
		return
	}

	panic(fmt.Sprintf("%v must be proceeded by a valid value", optionName))
}

// GetArgs ex:
//   - gurl --request GET https://httpbin.org
//   - gurl --request POST https://httpbin.org/post --header 'Content-Type: application/json' --data '{"name": "Almir", "job": "Developer"}'
//   - gurl --request POST https://httpbin.org/post --header 'Content-Type: multipart/form-data' --data name='Almir' --data job='Developer'
//   - gurl --request POST https://httpbin.org/post --header 'Content-Type: multipart/form-data' --data name='Almir' --file coolFile='./file.txt'
func ensureArgs() {
	if len(mappedArgs) > 0 {
		return
	}

	for i, arg := range os.Args {
		switch arg {
		case "-r":
		case "--request":
			method := strings.ToUpper(os.Args[i+1])
			ensureThatValueNextToOptionIsValid(method, arg)
			url := os.Args[i+2]
			ensureThatValueNextToOptionIsValid(url, arg)
			mappedArgs["request"] = []string{method, url}
		case "-h":
		case "--header":
			header := strings.ToLower(os.Args[i+1])
			ensureThatValueNextToOptionIsValid(header, arg)
			mappedArgs["header"] = append(mappedArgs["header"], header)
		case "-d":
		case "--data":
			data := os.Args[i+1]
			ensureThatValueNextToOptionIsValid(data, arg)
			mappedArgs["data"] = append(mappedArgs["data"], data)
		case "-f":
		case "--file":
			file := os.Args[i+1]
			ensureThatValueNextToOptionIsValid(file, arg)
			mappedArgs["file"] = append(mappedArgs["file"], file)
		case "-v":
		case "--verbose":
			mappedArgs["verbose"] = []string{""}
		}
	}
}

func GetRequest() (method string, url string) {
	ensureArgs()
	method = mappedArgs["request"][0]
	url = mappedArgs["request"][1]
	return method, url
}

func splitHeader(header string) (headerKey string, headerValue string) {
	headerData := strings.Split(header, ":")
	return headerData[0], strings.TrimSpace(headerData[1])
}

// --header 'Content-Type: application/json'
// --header 'Authorization: bearer 123'
func GetHeaders() map[string][]string {
	if len(mappedHeaders) > 0 {
		return mappedHeaders
	}

	ensureArgs()
	headers := mappedArgs["header"]
	for _, header := range headers {
		headerKey, headerValue := splitHeader(header)
		mappedHeaders[headerKey] = append(mappedHeaders[headerKey], headerValue)
	}

	mappedHeaders["accept"] = []string{"*/*"}
	mappedHeaders["user-agent"] = []string{"gurl/0.1.0"}
	return mappedHeaders
}

func splitOption(option string) (fieldName string, fieldValue string) {
	optionData := strings.Split(option, "=")
	return optionData[0], strings.TrimSpace(optionData[1])
}

// --data '{"name": "Almir", "job": "Developer"}'
// --data name='Almir' --data job='Developer'
// --data name='Almir' --file coolFile='./file.txt'
// --file coolFile='./file.txt'
func GetBodyData() (io.Reader, string) {
	ensureArgs()
	dataOptions := mappedArgs["data"]
	headers := GetHeaders()

	if slices.Contains(headers["content-type"], "application/json") {
		return strings.NewReader(dataOptions[0]), ""
	}

	if slices.Contains(headers["content-type"], "multipart/form-data") {
		body := bytes.Buffer{}
		multipartWriter := multipart.NewWriter(&body)
		defer func() {
			err := multipartWriter.Close()
			if err != nil {
				panic(err)
			}
		}()
		for _, dataOption := range dataOptions {
			fieldName, fieldValue := splitOption(dataOption)
			err := multipartWriter.WriteField(fieldName, fieldValue)
			if err != nil {
				panic(err)
			}
		}

		fileOptions := mappedArgs["file"]
		for _, fileOption := range fileOptions {
			fieldName, filePath := splitOption(fileOption)
			fileWriter, err := multipartWriter.CreateFormFile(fieldName, filepath.Base(filePath))
			if err != nil {
				panic(err)
			}

			file, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(fileWriter, file)
			if err != nil {
				panic(err)
			}
		}

		return &body, multipartWriter.FormDataContentType()
	}

	return nil, ""
}

func IsVerbose() bool {
	ensureArgs()
	return len(mappedArgs["verbose"]) > 0
}
