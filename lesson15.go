package mypackage

import "testing"

func FuzzReverse(f *testing.F) {
	// Seed corpus (initial valid inputs)
	testcases := []string{"Hello", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Add to seed corpus
	}

	// Fuzz target
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig) // Assume Reverse() is the function to test
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q,  is not same as After: %q", orig, doubleRev)
		}
	})
}

// %q , %d ...%v ..also
// for _,tc :=range array

// go test -fuzz=FuzzReverse  # Run a specific fuzz test
// go test -fuzz=.           # Run all fuzz tests
// go test -fuzz=FuzzReverse -fuzztime=10s  # Limit fuzzing duration

func FuzzParseInt(f *testing.F) {
	f.Add("123") // Seed input
	f.Fuzz(func(t *testing.T, s string) {
		_, err := strconv.Atoi(s)
		if err != nil && !strings.Contains(err.Error(), "invalid syntax") {
			t.Fatalf("Unexpected error: %v", err)
		}
	})
}

// strconv.Atoi(s) is to convert a string to int
// so basically err.Error is used when err is not nil and err.Error() returns a string which  is not equal to "invalid syntax"
// so then we print an unexpected error and test fails.

// t.Fatalf is used to fail a test and log an erorr immediately

func TestParseInt_Corpus(t *testing.T) {
	cases := []string{"0", "123", "-1"}
	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			_, _ = strconv.Atoi(c) // Just ensure no panic
		})
	}
}


go test -run=FuzzX/<hash>





type User struct {
    Name string
    Age  int
}

func FuzzUser(f *testing.F) {
    f.Add([]byte(`{"Name":"Alice","Age":30}`))  // JSON input, seed corpus .....
    f.Fuzz(func(t *testing.T, data []byte) {
        var u User
        if err := json.Unmarshal(data, &u); err != nil {
            return  // Skip invalid inputs
        }
        if u.Age < 0 {
            t.Errorf("Negative age: %d", u.Age)
        }
    })
}