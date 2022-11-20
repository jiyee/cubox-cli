package client

import (
	"encoding/json"
	"github.com/gojektech/heimdall/v6/httpclient"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Memo struct {
	Type    string
	Content string
	Tags    []string
	Folder  string
	API     string
}

type Link struct {
	Type        string
	Content     string
	Title       string
	Description string
	Tags        []string
	Folder      string
	API         string
}

type Payload struct {
	Type        string   `json:"type"`
	Content     string   `json:"content"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Folder      string   `json:"folder,omitempty"`
}

func (m *Memo) Submit(verbose bool) (*string, error) {
	content := strings.TrimSpace(m.Content)

	if m.API == "" || content == "" {
		return nil, errors.New("lack of necessary arguments")
	}

	timeout := 3000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	payloadJSON, _ := json.Marshal(Payload{
		m.Type,
		content,
		"",
		"",
		m.Tags,
		m.Folder,
	})
	body := ioutil.NopCloser(strings.NewReader(string(payloadJSON)))
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	if verbose {
		log.Printf("Raw content: %s", content)
		log.Printf("Payload JSON: %s", payloadJSON)
	}

	response, err := client.Post(m.API, body, headers)

	if err != nil {
		return nil, errors.Wrap(err, "failed to make a request to server")
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var responseData map[string]interface{}

	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		return nil, err
	}

	if verbose {
		log.Printf("Response Body: %v", responseData)
	}

	statusCode := response.StatusCode

	if statusCode >= 200 && statusCode < 400 {
		message := responseData["message"].(string)

		return &message, nil
	} else if statusCode >= 400 && statusCode < 500 {
		return nil, &ResponseError{
			Err:        errors.New("request is not valid"),
			StatusCode: statusCode,
		}
	} else {
		return nil, &ResponseError{
			Err:        errors.New("response error"),
			StatusCode: statusCode,
		}
	}
}

func (l *Link) Submit(verbose bool) (*string, error) {
	content := strings.TrimSpace(l.Content)

	if l.API == "" || content == "" {
		return nil, errors.New("lack of necessary arguments")
	}

	timeout := 3000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	payloadJSON, _ := json.Marshal(Payload{
		l.Type,
		content,
		l.Title,
		l.Description,
		l.Tags,
		l.Folder,
	})
	body := ioutil.NopCloser(strings.NewReader(string(payloadJSON)))
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	if verbose {
		log.Printf("Raw content: %s", content)
		log.Printf("Payload JSON: %s", payloadJSON)
	}

	response, err := client.Post(l.API, body, headers)

	if err != nil {
		return nil, errors.Wrap(err, "failed to make a request to server")
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var responseData map[string]interface{}

	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		return nil, err
	}

	if verbose {
		log.Printf("Response Body: %v", responseData)
	}

	statusCode := response.StatusCode

	if statusCode >= 200 && statusCode < 400 {
		message := responseData["message"].(string)

		return &message, nil
	} else if statusCode >= 400 && statusCode < 500 {
		return nil, &ResponseError{
			Err:        errors.New("request is not valid"),
			StatusCode: statusCode,
		}
	} else {
		return nil, &ResponseError{
			Err:        errors.New("response error"),
			StatusCode: statusCode,
		}
	}
}
