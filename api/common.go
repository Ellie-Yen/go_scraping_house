package api

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Request struct {
	Url     string
	Method  string
	Headers [][]string
	Cookies [][]string
	Body    io.Reader
}

func newRequest(r *Request) *http.Request {
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return nil
	}
	// set headers
	for _, header := range r.Headers {
		req.Header.Set(header[0], header[1])
	}

	// set cookies
	for _, cookie := range r.Cookies {
		req.AddCookie(&http.Cookie{Name: cookie[0], Value: cookie[1]})
	}

	return req
}

func Do(r *Request) ([]byte, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Send request
	req := newRequest(r)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %v\n", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body) // Read response body even on error

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
