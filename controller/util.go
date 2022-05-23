package controller

import "testing"

func assertStringValues(t *testing.T, testName string, expected, actual string) {
	if expected != actual {
		t.Errorf(" %s Expected value %s. Got %s\n", testName, expected, actual)
	}
}

func assertStatus(t *testing.T, testName string, expected, actual int) {
	if expected != actual {
		t.Errorf(" %s Expected response code %d. Got %d\n", testName, expected, actual)
	}
}
