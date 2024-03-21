package main

import "fmt"

// Group constants in a block instead of declaring them each on their own line
const (
    french = "French"
    spanish = "Spanish"

    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
    frenchHelloPrefix = "Bonjour, "
)

// Public functions start with a capital letter
func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

// Private functions start with lowercase
// Named return value: (prefix string)
// this will create a variable called prefix in the function
// you can return wheaterver it's set to by just calling return
// rather than return prefix
func greetingPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("world", ""))
}

