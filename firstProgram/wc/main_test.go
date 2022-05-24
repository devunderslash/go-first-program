package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("This is a test word1 word2 word3\n")

	exp := 7
	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("This is a test word1 word2 word3\nline2\nline3\n word5")

	exp := 4
	res := count(b, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead", exp, res)
	}
}
