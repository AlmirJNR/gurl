package main

import (
	"fmt"
	"gurl/internal/utils"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		utils.ShowHelpMessage()
		os.Exit(0)
	}

	method, url := utils.GetRequest()
	bodyData, contentType := utils.GetBodyData()
	request, err := http.NewRequest(method, url, bodyData)
	if err != nil {
		panic(err)
	}

	headers := utils.GetHeaders()
	request.Header = headers
	if contentType != "" {
		request.Header.Set("content-type", contentType)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	if utils.IsVerbose() {
		err = response.Write(os.Stdout)
		if err != nil {
			panic(err)
		}
		return
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(responseBody))
}
