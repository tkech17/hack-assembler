package main

import "testing"

func AssertTrue(t *testing.T, value bool, message string) {
	if !value {
		t.Fatal(message)
	}
}

func AssertFalse(t *testing.T, value bool, message string) {
	if value {
		t.Fatal(message)
	}
}

func AssertEqualsInt(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Fatalf("expected [%d], actual [%d]", expected, actual)
	}
}


func AssertNotEqualsInt(t *testing.T, expected int, actual int) {
	if expected == actual {
		t.Fatal("Not expcted equal values")
	}
}

func AssertEqualsString(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Fatalf("expected [%s], actual [%s]", expected, actual)
	}
}
func AssertNull(t *testing.T, value interface{}) {
	if isNull(value) {
		t.Fatal("Should Be Null")
	}
}

func AssertNotNull(t *testing.T, value interface{}) {
	if !isNull(value) {
		t.Fatal("Should Not Be Null")
	}
}

func isNull(value interface{}) bool {
	return value != nil
}
