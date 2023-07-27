package main

import "testing"

func Test_run(t *testing.T) {
	// TODO: Add test cases.
	want := "Command Alias Console"
	got := "Command Alias Console"
	if got != want {
		t.Errorf("run() = %q, want %q", got, want)
	}
}
