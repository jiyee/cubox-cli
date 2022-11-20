package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMemoCommand_Usage(t *testing.T) {
	command := MemoCommand{}

	command.Usage()
}

func TestMemoCommand_Run(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		requestBody, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, `{"type":"memo","content":"Test"}`, string(requestBody[:]))
		rw.Write([]byte(`{"code":200,"message":""}`))
	}))

	command := MemoCommand{
		API: server.URL,
	}

	if _, err := command.Run([]string{"Test"}); err != nil {
		log.Fatal(err)
	}
}
