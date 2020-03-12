package helper

import "fmt"

// CheckErr checks if there is an error, if true
// print the error.
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
