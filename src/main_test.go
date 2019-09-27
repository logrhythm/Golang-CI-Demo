package main

import "testing"

func TestHelloWorld(t *testing.T) {
	var msg = helloWorld()
	if msg != "Hello World.. This CI works" {
		t.Errorf("helloWorld did not work, expected Hello World! but got %s", msg)
	}
}
