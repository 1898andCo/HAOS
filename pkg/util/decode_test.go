package util_test

import (
	"os"
	"testing"

	"github.com/1898andCo/HAOS/pkg/util"
)

const (
	base64Content = "SXQgd2FzIHRoZSBiZXN0IG9mIHRpbWVzLCBpdCB3YXMgdGhlIHdvcnN0IG9mIHRpbWVz"
)

func TestDecodeContent(t *testing.T) {
	_, err := util.DecodeContent("foo", "bar")
	if err == nil {
		t.Errorf("DecodeContent() expected unsupported encoding error")
	}

	result, err := util.DecodeContent(base64Content, "base64")
	if err != nil {
		t.Errorf("DecodeContent() error = %v", err)
	}
	expected := "It was the best of times, it was the worst of times"
	if string(result) != expected {
		t.Errorf("DecodeContent() = %v, want %v", string(result), expected)
	}
	content, _ := os.ReadFile("gzip_test.gz")
	example := string(content)
	result, err = util.DecodeContent(example, "gzip")
	if err != nil {
		t.Errorf("DecodeContent() error = %v", err)
	}
	expected = "foo\n"
	if string(result) != expected {
		t.Errorf("DecodeContent() = %v, want %v", string(result), expected)
	}

}

func TestDecodeBase64Content(t *testing.T) {
	result, err := util.DecodeBase64Content(base64Content)
	if err != nil {
		t.Errorf("DecodeBase64Content() error = %v", err)
	}
	expected := "It was the best of times, it was the worst of times"
	if string(result) != expected {
		t.Errorf("DecodeBase64Content() = %v, want %v", string(result), expected)
	}
}

func TestDecodeGzipContent(t *testing.T) {
	content, _ := os.ReadFile("gzip_test.gz")
	example := string(content)
	result, err := util.DecodeGzipContent(example)
	if err != nil {
		t.Errorf("DecodeGzipContent() error = %v", err)
	}
	expected := "foo\n"
	if string(result) != expected {
		t.Errorf("DecodeGzipContent() = %v, want %v", string(result), expected)
	}

}
