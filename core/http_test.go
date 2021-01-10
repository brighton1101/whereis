package core

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBlockedRedirect(t *testing.T) {
	port := "8080"
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", port))
	if err != nil {
		t.Errorf("Could not listen on: %s", port)
	}
	baseuri := fmt.Sprintf("http://127.0.0.1:%s/", port)
	redirecteduri := baseuri + "hello"
	server := httptest.NewUnstartedServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.String() != "/" {
				t.Errorf("Expected %s got %s", "/", req.URL.String())
			}
			http.Redirect(rw, req, redirecteduri, http.StatusMovedPermanently)
		}))
	server.Listener.Close()
	server.Listener = listener
	server.Start()
	defer server.Close()

	resp, err := GetNoRedirect(baseuri)
	if err != nil {
		t.Errorf("Unknown error")
		return
	}
	if resp.RedirectedUri != redirecteduri {
		t.Errorf(
			"Expected %s got %s", redirecteduri, resp.RedirectedUri,
		)
	}
	if resp.StatusCode != 301 {
		t.Errorf(
			"Expected %d got %d", 301, resp.StatusCode,
		)
	}
}
