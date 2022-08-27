package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mh testHandler
	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is http.Handler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var mh testHandler
	h := SessionLoad(&mh)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("type is http.Handler, but is %T", v))
	}
}
