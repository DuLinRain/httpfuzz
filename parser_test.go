package httpfuzz

import "testing"

func TestHTTPRequestInvalidFileReturnsError(t *testing.T) {
	req, err := RequestFromFile("notfound.request")
	if err == nil {
		t.Fatalf("expected error")
	}

	if req != nil {
		t.Fatalf("request returned when expected nil: %+v", req)
	}
}

func TestHTTPRequestParsedCorrectlyFromFile(t *testing.T) {
	req, err := RequestFromFile("./testdata/validGET.request")
	if err != nil {
		t.Fatalf("expected err to be nil, got %v", err)
	}

	if req.Method != "GET" {
		t.Fatalf("expected GET, got %v", req.Method)
	}

	if req.Host != "detectportal.firefox.com" {
		t.Fatalf("expected URL 'detectportal.firefox.com', got %v", req.Host)
	}

	if req.UserAgent() != "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:78.0) Gecko/20100101 Firefox/78.0" {
		t.Fatalf("got unexpected User-Agent %v", req.UserAgent())
	}

	if req.Header.Get("Cache-Control") != "no-cache" {
		t.Fatalf("got unexpected cache-control header")
	}
}
