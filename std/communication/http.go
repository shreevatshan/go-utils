package communication

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const (
	ProtocolHTTP  = "http://"
	ProtocolHTTPS = "https://"

	RequestTypeGET  = http.MethodGet
	RequestTypePOST = http.MethodPost
	RequestTypePUT  = http.MethodPut

	ContentTypeApplicationJSON = "application/json"

	ContentEncodingGzip = "gzip"

	HeaderContentType     = "Content-Type"
	HeaderContentEncoding = "Content-Encoding"

	defaultHTTPRequestTimeout = 10
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
	DownloadLocation   string
	Headers            map[string]string
}

type HTTPResponse struct {
	Code     int
	Body     []byte
	Duration time.Duration
	Err      error
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

	err = os.MkdirAll(filepath.Dir(request.DownloadLocation), 0755)
	if err != nil {
		response.Err = fmt.Errorf("error creating directory for package [%v]", err)
		return
	}

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

func Gzip(payload []byte) ([]byte, error) {

	if payload == nil {
		return nil, nil
	}

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	if _, err := gz.Write(payload); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func ExtractGzip(gzipped []byte) ([]byte, error) {

	if gzipped == nil {
		return nil, nil
	}

	r, err := gzip.NewReader(bytes.NewReader(gzipped))
	if err != nil {
		return nil, err
	}

	defer r.Close()

	unzipped, err := io.ReadAll(r)

	return unzipped, err
}

func (request *HTTPRequest) Send() HTTPResponse {
	var response HTTPResponse
	var req *http.Request
	var err error

	if request.TimeOut == 0 {
		request.TimeOut = defaultHTTPRequestTimeout
	}

	err = request.validateRequest()
	if err != nil {
		response.Err = err
		return response
	}

	transport := http.Transport{
		Proxy:             http.ProxyURL(request.Proxy),
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: request.InsecureSkipVerify},
		DisableKeepAlives: true, // disable keep-alive to prevent connection leaks as we are not reusing connections
	}

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

	timeStart := time.Now()
	httpResponse, err := client.Do(req)
	response.Duration = time.Since(timeStart)
	if err != nil {
		response.Err = err
		return response
	}
	defer httpResponse.Body.Close()

	response.Code = httpResponse.StatusCode

	if httpResponse.StatusCode != http.StatusOK {
		response.Err = fmt.Errorf("response code not OK (%d)", response.Code) // set a default error message
	}

	if request.DownloadResponse {
		response.downloadResponse(httpResponse, request)
	} else {
		response.Body, err = io.ReadAll(httpResponse.Body)
		if err != nil {
			response.Err = err
		}
	}

	return response
}
