// Tests need to be in a file with a name like "XXX_test.go"
package main

// In order to use the "*testing.T" type, you need to import "testing"
import "testing"

// Test function must start with the word "Test"
// The test function takes one argument only "t *testing.T"
func TestHello(t *testing.T) {
    got := Hello("Lucas")
    want := "Hello, Lucas"

    if got != want {
        // Print out a message and fail the test
        // "f" stands for format which allow us to build a string with values
        // inserted into the placeholder values %q
        t.Errorf("got %q want %q", got, want)
    }
}

