// Tests need to be in a file with a name like "XXX_test.go"
package main

// In order to use the "*testing.T" type, you need to import "testing"
import "testing"

// Test function must start with the word "Test"
// The test function takes one argument only "t *testing.T"
func TestHello(t *testing.T) {
    // Subtests
    // Sometimes it is useful to group tests around a "thing" and then
    // have subtests describing different scenarios
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Lucas")
        want := "Hello, Lucas"
        assertCorrectMessage(t, got, want)
    })

    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"
        assertCorrectMessage(t, got , want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    // Is needed to tell the test suite that this method is a helper.
    // By doing this when it fails the line number reported will be in our
    // function call rather than inside our test helper.
    // This will help other developers track down problems easier
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

