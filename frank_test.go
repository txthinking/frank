package main

import "testing"

func TestFrank(t *testing.T) {
	c, err := NewCase("./case.frank", 0)
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Run(); err != nil {
		t.Fatal(err)
	}
}
