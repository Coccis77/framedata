package util

import "testing"

func TestContainsIgnoreCaseEmptySlice(t *testing.T) {
	characters := []string{}
	if ContainsIgnoreCase(characters[:], "Sora") {
		t.Error("Empty slice must return false")
	}
}

func TestContainsIgnoreCaseNilSlice(t *testing.T) {
	var characters []string
	if ContainsIgnoreCase(characters[:], "Sora") {
		t.Error("Nil slice must return false")
	}
}

func TestContainsIgnoreCaseOneExistingElement(t *testing.T) {
	characters := []string{"Sora"}
	if !ContainsIgnoreCase(characters[:], "Sora") {
		t.Error("Existing value must return true")
	}
}

func TestContainsIgnoreCaseOneNonExistingElement(t *testing.T) {
	characters := []string{"Sora"}
	if ContainsIgnoreCase(characters[:], "Kairi") {
		t.Error("Non-existing value must return false")
	}
}

func TestContainsIgnoreCaseOneExistingElementDifferentCase(t *testing.T) {
	characters := []string{"Sora"}
	if !ContainsIgnoreCase(characters[:], "SORA") {
		t.Error("Existing value with different case must return true")
	}
}