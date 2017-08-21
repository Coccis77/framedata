package util

import "testing"

func TestContainsIgnoreCaseAndWithoutSpaceEmptySlice(t *testing.T) {
	characters := []string{}
	if ContainsIgnoreCaseAndWithoutSpace(characters[:], "Sora") {
		t.Error("Empty slice must return false")
	}
}

func TestContainsIgnoreCaseAndWithoutSpaceNilSlice(t *testing.T) {
	var characters []string
	if ContainsIgnoreCaseAndWithoutSpace(characters[:], "Sora") {
		t.Error("Nil slice must return false")
	}
}

func TestContainsIgnoreCaseAndWithoutSpaceOneExistingElement(t *testing.T) {
	characters := []string{"Sora"}
	if !ContainsIgnoreCaseAndWithoutSpace(characters[:], "Sora") {
		t.Error("Existing value must return true")
	}
}

func TestContainsIgnoreCaseAndWithoutSpaceOneNonExistingElement(t *testing.T) {
	characters := []string{"Sora"}
	if ContainsIgnoreCaseAndWithoutSpace(characters[:], "Kairi") {
		t.Error("Non-existing value must return false")
	}
}

func TestContainsIgnoreCaseAndWithoutSpaceOneExistingElementDifferentCase(t *testing.T) {
	characters := []string{"Sora"}
	if !ContainsIgnoreCaseAndWithoutSpace(characters[:], "SORA") {
		t.Error("Existing value with different case must return true")
	}
}

func TestEqualFoldWithoutSpaceNoSpace(t *testing.T) {
	if !EqualFoldWithoutSpace("Sora", "sora") {
		t.Error("Existing value with different case must return true")
	}
}

func TestEqualFoldWithoutSpaceWithSpace(t *testing.T) {
	if !EqualFoldWithoutSpace("So ra", "s  ora") {
		t.Error("Existing value with different case must return true")
	}
}