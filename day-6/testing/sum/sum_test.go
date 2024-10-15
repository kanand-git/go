package sum

import "testing"

// test functions should be prefixed with word Test
func TestSumInt(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumInt(x)

	if got != want {
		// Errorf would signal test failure but your test would continue to run
		t.Errorf("sum of 1 to 5 should be %v; got %v", want, got)

		// Fatalf would quit the test as well
		//t.Fatalf("sum of 1 to 5 should be %v; got %v", want, got)
	}
	want = 0
	got = SumInt(nil)
	if got != want {
		t.Errorf("sum of nil should be %v; got %v", want, got)
		//t.Fatalf()
	}
}

func TestSumIntTableTest(t *testing.T) {
	tt := []struct {
		name    string // required
		numbers []int
		want    int
	}{
		// individual test cases
		{
			name:    "one to five",
			numbers: []int{1, 2, 3, 4, 5},
			want:    15,
		},
		{
			name:    "nil",
			numbers: nil,
			want:    0,
		},
		{
			name:    "empty",
			numbers: []int{},
			want:    0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) { // Use 't.Run()' to execute a sub-test for each test case
			got := SumInt(tc.numbers)
			if got != tc.want {
				t.Fatalf("sum of %v should be %v; got %v", tc.numbers, tc.want, got)
			}
		})

	}
}
