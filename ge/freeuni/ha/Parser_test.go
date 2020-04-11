package main

import (
	"testing"
)

func Test_should_returnIsEmptyString_when_givenEmptyString(t *testing.T) {
	var s = ""
	var isNotEmptyString = isNotString(s)
	assertFalse(t, isNotEmptyString, "string [" + s + "] is not empty line")
}

func Test_should_returnIsNotEmptyString_when_givenNonEmptyString(t *testing.T) {
	var s = "abc"
	var isNotEmptyString = isNotString(s)
	assertTrue(t, isNotEmptyString, "string [" + s + "] is empty line")
}

func assertTrue(t *testing.T, value bool, message string) {
	if !value {
		t.Fatal(message)
	}
}

func assertFalse(t *testing.T, value bool, message string) {
	if value {
		t.Fatal(message)
	}
}
