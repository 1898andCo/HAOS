package questions

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

type mockBool struct{}

func (mockBool) Stdin() *bufio.Reader {
	return bufio.NewReader(bytes.NewBufferString("y\n"))
}

type mockBlank struct{}

func (mockBlank) Stdin() *bufio.Reader {
	return bufio.NewReader(bytes.NewBufferString(" \n"))
}

func TestPrompt(t *testing.T) {
	impl = mockBool{}
	ans, err := Prompt("Enter y: ", "")
	if err != nil {
		t.Fatal(err)
	}
	if ans != "y" {
		t.Fatalf("expected y, got %s", ans)
	}
	fmt.Println(ans)

}

func TestPromptBool(t *testing.T) {
	impl = mockBool{}
	ans, err := PromptBool("Enter y: ", false)
	if err != nil {
		t.Fatal(err)
	}
	if !ans {
		t.Fatalf("expected true, got %t", ans)
	}
	fmt.Println(ans)

}

func TestPromptOptional(t *testing.T) {
	impl = mockBlank{}
	ans, err := PromptOptional("Enter something: ", "default")
	if err != nil {
		t.Fatal(err)
	}
	if ans != "default" {
		t.Fatalf("expected 'default', got %s", ans)
	}
	fmt.Println(ans)

}
