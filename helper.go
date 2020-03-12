package helper

import "fmt"

// PrintErr checks if the err is nil, if not,
// print the error.
func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
