package main

import (
	"net/http"
	"testing"
)

type rr string

func (self rr) Header() http.Header {
	return nil
}

func (self rr) WriteHeader(int) {
}

func (self rr) Write(b []byte) (int, error) {
	return len(b), nil
}

func Test_Writer(t *testing.T) {

}
