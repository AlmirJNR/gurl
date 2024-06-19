# gurl (wip)

`curl` but written in `go`

## http requests

`gurl` haves many options to send your http data to servers, you can use the conventional methods:
POST, GET, PUT, PATCH, DELETE, HEAD and OPTIONS

## flags

- --request [POST | GET | PUT | PATCH | DELETE | HEAD | OPTIONS] string
    - base identifier of http method request, proceeded by the url
- --header string
    - header values to be passed
- --data [string | key=value]
    - data values or key/values to be passed,
      generally values if json and key/values if form data
- --file key=value
    - file field and value to be passed with the form data content type
- --verbose
    - tells the application that it should print on the stdout with all content from the response, not only the body

## usage examples

```shell
gurl --request GET {URL}
```

```shell
gurl --request GET {URL} \
--header 'Authorization: bearer {TOKEN}'
```

```shell
gurl --request POST {URL} \
--header 'Content-Type: application/json' \
--header 'Authorization: bearer {TOKEN}' \
--data '{"name": "Almir", "job": "Developer"}'
```

```shell
gurl --request POST {URL} \
--header 'Content-Type: multipart/form-data' \
--header 'Authorization: bearer {TOKEN}' \
--data name=Almir \
--data job=Developer \
--data age=23 \
--file genericFile={FILE_PATH} \
--file genericFile2={FILE_PATH_2}
```
