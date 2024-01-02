package manifests

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/1898andCo/HAOS/pkg/config"
)

type Mock struct{}

func (Mock) Get(url string) (*http.Response, error) {
	stringReader := strings.NewReader("this is the body")
	stringReadCloser := io.NopCloser(stringReader)
	return &http.Response{
		StatusCode:    http.StatusOK,
		Body:          stringReadCloser,
		ContentLength: stringReader.Size(),
	}, nil
}

var mockData = []byte("")

func (Mock) WriteFile(path string, data []byte, perm os.FileMode) error {
	mockData = data
	return nil
}

type MockError struct{}

func (MockError) Get(url string) (*http.Response, error) {
	stringReader := strings.NewReader("")
	stringReadCloser := io.NopCloser(stringReader)
	return &http.Response{
		StatusCode:    http.StatusInternalServerError,
		Body:          stringReadCloser,
		ContentLength: stringReader.Size(),
	}, nil
}

func TestApplyBootManifests(t *testing.T) {
	manifest := config.Manifest{
		URL:    "localhost",
		SHA256: "",
	}
	impl = Mock{}
	cfg := &config.CloudConfig{
		BootManifests: []config.Manifest{manifest},
	}
	err := ApplyBootManifests(cfg)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if string(mockData) != "this is the body" {
		t.Errorf("expected %s, got %s", "this is the body", string(mockData))
	}
}
