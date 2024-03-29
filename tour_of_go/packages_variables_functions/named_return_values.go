package main

import "fmt"

// Go's return values may be named
// If so, they are treated as variables defined at the top of the function
// These names should be used to document the meaning of the return values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    // A return statement without arguments returns the named return values
    // This is known as a "naked" return
    // Naked return statements should be used only in short functions
    return
}

func main(){
    fmt.Println(split(17))
}

