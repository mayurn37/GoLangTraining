package main

import "testing"

//test case to check age of student
func TestGetAge(t *testing.T) {
	ss := Student{19, "Vivek", 28, ContactInfo{"1234", "Calsoft Pune"}}
	age := 28
	if age != ss.getAge() {
		t.Errorf("Age is %d, expected %d", ss.getAge(), age)
	}
}