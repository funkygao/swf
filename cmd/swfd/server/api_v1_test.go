package server

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func setupServer() *Server {
	return &Server{}
}

func TestHandleApiV1(t *testing.T) {
	ts := httptest.NewServer(http.HandleFunc("/v1", handler))
	defer ts.Close()

	rw := httptest.NewRecorder()
	s.apiServer.handleApiV1(rw, req(t, "v"), params)
}

func TestHandleHi_TestServer_Parallel(t *testing.T) {
	s := setupServer()
	defer s.stop()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

		}()

	}
	wg.Wait()
}

func req(t *testing.T, v string) *http.Request {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(v)))
	if err != nil {
		t.Fatal(err)
	}
	return req
}
