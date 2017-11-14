package main

import "testing"

func TestNewCase(t *testing.T) {
	c, err := NewCase("./case.frank")
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Run(); err != nil {
		t.Fatal(err)
	}
}
