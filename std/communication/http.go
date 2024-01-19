package communication

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	ProtocolHTTP  = "http"
	ProtocolHTTPS = "https"

	RequestTypeGet  = "GET"
	RequestTypePost = "POST"
	RequestTypePut  = "PUT"

	ContentTypeApplicationJSON = "application/json"

	defaultHTTPRequestTimeout = 10
	HTTPStatusOK              = 200
)

type HTTPRequest struct {
	TimeOut            int
	DownloadResponse   bool
	Body               []byte
	URL                *url.URL
	Proxy              *url.URL
	InsecureSkipVerify bool
	API                string
	RequestType        string
	ContentType        string
	DownloadLocation   string
	Headers            map[string]string
}

type HTTPResponse struct {
	Code int
	Body []byte
	Err  error
}

func formAPI(URL *url.URL) (string, error) {
	api := URL.String()

	parsedURL, err := url.ParseRequestURI(api)
	if err != nil {
		return "", err
	} else if len(parsedURL.Hostname()) == 0 {
		return "", fmt.Errorf("not a valid HTTP request uri (%s)", api)
	} else {
		return api, nil
	}

}

func (request *HTTPRequest) validateRequest() error {
	if len(request.RequestType) == 0 {
		return errors.New("request type not set")
	}

	if len(request.API) == 0 {
		var err error
		request.API, err = formAPI(request.URL)
		if err != nil {
			return err
		}
	}

	return nil
}

func (response *HTTPResponse) downloadResponse(resp *http.Response, request *HTTPRequest) {

	var err error

	out, err := os.Create(request.DownloadLocation)
	if err != nil {
		response.Err = fmt.Errorf("error creating file for package [%v]", err)
		return
	}

	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		response.Err = fmt.Errorf("error saving package [%v]", err)
		return
	}
}

func (request *HTTPRequest) Send() HTTPResponse {
	var response HTTPResponse
	var req *http.Request
	var err error

	if request.TimeOut == 0 {
		request.TimeOut = defaultHTTPRequestTimeout
	}

	request.validateRequest()
	if err != nil {
		response.Err = err
		return response
	}

	transport := http.Transport{
		Proxy:           http.ProxyURL(request.Proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: request.InsecureSkipVerify}}

	client := http.Client{
		Timeout:   time.Duration(request.TimeOut) * time.Second,
		Transport: &transport}

	req, err = http.NewRequest(request.RequestType, request.API, bytes.NewBuffer(request.Body))
	if err != nil {
		response.Err = err
		return response
	}

	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	httpResponse, err := client.Do(req)
	if err != nil {
		response.Err = err
		return response
	}

	defer httpResponse.Body.Close()

	response.Code = httpResponse.StatusCode

	if httpResponse.StatusCode == HTTPStatusOK {
		if request.DownloadResponse {
			response.downloadResponse(httpResponse, request)
		} else {
			response.Body, err = io.ReadAll(httpResponse.Body)
			if err != nil {
				response.Err = err
				return response
			}
		}
	} else {
		response.Err = fmt.Errorf("response code not OK (%d)", response.Code)
	}

	return response
}
