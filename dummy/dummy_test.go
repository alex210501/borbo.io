package dummy

import (
	"bytes"
	"testing"
)

func TestSayHellWord(t *testing.T) {
	out = new(bytes.Buffer)

	SayHellWorld()
	got := out.(*bytes.Buffer).String()

	if got != "hello world" {
		t.Errorf("%q != \"hello world\"", got)
	}
}

func TestStringToWrite(t *testing.T) {
	out = new(bytes.Buffer)

	StringToWrite()

	got := out.(*bytes.Buffer).String()

	if got != stringToWrite {
		t.Errorf("%q != %q", got, stringToWrite)
	}
}
