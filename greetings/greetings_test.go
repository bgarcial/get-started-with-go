package greetings

import (
	"testing"
	"regexp"
)

/*
Implement test functions in the same package as the code you're testing.
Create two test functions to test the greetings.Hello function.

- Test function names have the form TestName, where Name is specific 
to the test.

- Also, test functions take a pointer to the testing package's testing
https://golang.org/pkg/testing/
T as a parameter. You use this parameter's methods for reporting 
and logging from your test.


*/

// TestHelloName calls greetings.Hello with a name, checking 
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Bernardo"
	want := regexp.MustCompile(`\b`+name+`\b`)

	/*
	TestHelloName calls the Hello function, passing a name value with 
	which the function should be able to return a valid response message
	*/
	msg, err := Hello("Bernardo")

	/*
	If the call returns an error or an unexpected response message 
	(one that doesn't include the name you passed in), you use the 
	t parameter's Fatalf method to print a message to the console 
	and end execution.
	*/
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Bernardo") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string, 
// checking for an error.
/*
This test is designed to confirm that your error handling works.
*/
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	/*
	If the call returns a non-empty string or no error, 
	you use the t parameter's Fatalf method to print a message 
	to the console and end execution.
	*/
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

/*
At the command line in the greetings directory, run the 
go test command to execute the test.
The go test command executes test functions 
(whose names begin with Test) in test files 
(whose names end with _test.go). 

You can add the -v flag to get verbose output that lists 
all of the tests and their results.

The tests should pass.

*/