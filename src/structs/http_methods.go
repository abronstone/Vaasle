package structs

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func DecodeResponseBody[T interface{}](res *http.Response, responseStructPointer T) error {
	/*
		Takes in a pointer to a struct, and decodes an HTTP response body into the parametized struct in place

		@param:

			http.Response
			pointer to struct (interface{})

		@return:

			error
	*/
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, responseStructPointer)
	if err != nil {
		return err
	}

	return nil
}

func MakePutRequest[T any](endpoint string, requestStruct interface{}, responseStructPointer ...*T) (*http.Response, error) {
	/*
		Calls a PUT request with the specified endpoint, using the 'requestStruct' as the request body, and decodes the response body into 'responseStructPointer', if any

		@param:

			endpoint (string)
			request body struct (interface{})
			pointer to generic type (*interface{})

		@return:

			http response
			error
	*/

	// Step 1: Marshal body bytes from request structure
	var bodyBuffer *bytes.Buffer
	if requestStruct != nil {
		bodyBytes, err := json.Marshal(requestStruct)
		if err != nil {
			return nil, err
		}
		bodyBuffer = bytes.NewBuffer(bodyBytes)
	}

	// Step 2: Make request
	req, err := http.NewRequest(http.MethodPut, endpoint, bodyBuffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return res, errors.New("PUT request returned status code " + strconv.Itoa(res.StatusCode))
	}
	// Step 3: Check to see if responseStructPointer variadic parameter has a single value. If so, decode the response into it
	if len(responseStructPointer) == 1 {
		err = DecodeResponseBody(res, responseStructPointer[0])
		if err != nil {
			return nil, err
		}
	} else if len(responseStructPointer) > 1 {
		return nil, errors.New("more than one response struct provided")
	}

	return res, nil
}

func MakePostRequest[T any](endpoint string, requestStruct interface{}, responseStructPointer ...*T) (*http.Response, error) {
	/*
		Calls a POST request with the specified endpoint, using the 'requestStruct' as the request body, and decodes the response body into 'responseStructPointer', if any

		@param:

			endpoint (string)
			request body struct (interface{})
			pointer to generic type (*interface{})

		@return:

			http response
			error
	*/

	// Step 1: Marshal body bytes from request structure
	var bodyBuffer *bytes.Buffer
	if requestStruct != nil {
		bodyBytes, err := json.Marshal(requestStruct)
		if err != nil {
			return nil, err
		}
		bodyBuffer = bytes.NewBuffer(bodyBytes)
	}

	// Step 2: Make request
	res, err := http.Post(endpoint, "application/json", bodyBuffer)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return res, errors.New("POST request returned status code " + strconv.Itoa(res.StatusCode))
	}

	// Step 3: Check to see if responseStructPointer variadic parameter has a single value. If so, decode the response into it
	if len(responseStructPointer) == 1 {
		err = DecodeResponseBody(res, responseStructPointer[0])
		if err != nil {
			return nil, err
		}
	} else if len(responseStructPointer) > 1 {
		return nil, errors.New("more than one response struct provided")
	}

	return res, nil
}

func MakeGetRequest[T any](endpoint string, responseStructPointer ...*T) (*http.Response, error) {
	/*
		Calls a GET request with the specified endpoint, and decodes the response body into 'responseStructPointer', if any

		@param:

			endpoint (string)
			pointer to generic type (*interface{})

		@return:

			http response
			error
	*/
	// Step 1: Make request
	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	// Step 2: Check to see if responseStructPointer variadic parameter has a single value. If so, decode the response into it
	if len(responseStructPointer) == 1 {
		err = DecodeResponseBody(res, responseStructPointer[0])
		if err != nil {
			return nil, err
		}
	} else if len(responseStructPointer) > 1 {
		return nil, errors.New("more than one response struct provided")
	}
	return res, nil
}
