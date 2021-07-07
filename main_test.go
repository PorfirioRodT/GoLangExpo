package main

import "testing"

func TestValuesShouldBeSorted (t *testing.T){

	w := "bac"

	if sortingValues(w) != "abc" {

		t.Error("Values are not getting sorted")

	}

}