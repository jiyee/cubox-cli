package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLinkCommand_Usage(t *testing.T) {
	command := LinkCommand{}

	command.Usage()
}

func TestLinkCommand_Run(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, `{"type":"url","content":"https://cubox.pro/"}`, string(requestBody[:]))
		rw.Write([]byte(`{"code":200,"message":""}`))
	}))

	command := LinkCommand{
		API:     server.URL,
		Content: "https://cubox.pro/",
	}

	if _, err := command.Run([]string{"Test"}); err != nil {
		log.Fatal(err)
	}
}
