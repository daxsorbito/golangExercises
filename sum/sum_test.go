package sum_test

import (
	"testing"
)

func TestIntsTTWithMessageSubTest(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"no numbers", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			sum := sum.Ints(tc.numbers...)
			if sum != tc.sum {
				t.Errorf("sum of %v should be %v; got %v", tc.name, tc.sum, sum)
			}
		})
	}
}

func TestIntsTTWithMessage(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"no numbers", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}

	for _, tc := range tt {
		sum := sum.Ints(tc.numbers...)
		if sum != tc.sum {
			t.Errorf("sum of %v should be %v; got %v", tc.name, tc.sum, sum)
		}
	}
}
func TestIntsTT(t *testing.T) {
	tt := []struct {
		numbers []int
		sum     int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{nil, 0},
		{[]int{1, -1}, 0},
	}

	for _, tc := range tt {
		sum := sum.Ints(tc.numbers...)
		if sum != tc.sum {
			t.Errorf("sum of %v should be %v; got %v", tc.numbers, tc.sum, sum)
		}
	}
}
func TestInts(t *testing.T) {
	// t.Fail()
	// t.Errorf("this test failed because I said so")
	// t.Fatalf("this test failed and stopped running")
	// t.Errorf("this should not execute if Fatalf is executed")

	s := sum.Ints(1, 2, 3, 4, 5)
	if s != 15 {
		t.Errorf("sum of one to five should be 15; got %v", s)
	}

	s = sum.IntsArray([]int{1, 2, 3, 4, 5})
	if s != 15 {
		t.Errorf("sum of one to five should be 15; got %v", s)
	}

	s = sum.Ints()
	if s != 0 {
		t.Errorf("sum of no numbers should be 0; got %v", s)
	}

	s = sum.Ints(1, -1)
	if s != 0 {
		t.Errorf("sum of one and -one should be 0; got %v", s)
	}
}
