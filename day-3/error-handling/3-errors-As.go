package main

import (
	"errors"
	"fmt"
	"log"
)

//Error should be suffixed in the name of the struct use for error handling

// QueryError is a custom error used to create dynamic error messages
type QueryError struct {
	Func  string
	Input string
	Err   error // this field must be present when working with Error structs
}

// implementing the error interface// the error method must be implemented over a pointer
func (q *QueryError) Error() string {
	return "main." + q.Func + ": " + "input " + q.Input + " " + q.Err.Error()
}

func main() {
	//fmt.Println(strconv.Atoi("abc"))
	//fmt.Println(strconv.Atoi("xyz"))
	//fmt.Println(strconv.ParseInt("ajay", 10, 64))
	////fmt.Println(strconv.ParseFloat("dev", 64))
	//errors.New()
	//os.PathError{}
	//os.OpenFile()

	err := SearchSomething("abc")
	if err != nil {
		// we need to update the QueryError in errors.As func so we need to make it a pointer because
		// when we pass structs there values are copied to param not the address
		var qe *QueryError // nil // must be a pointer

		// errors.As check if a struct is present in the chain or not for error,
		// if yes than fill the values in qe var with the values of QueryError type
		if errors.As(err, &qe) { // must be passed as reference to update nil pointer
			fmt.Println(err)
			fmt.Println("Query Error", qe.Func, qe.Err)
			return
		}
		log.Println("Error", err)
	}

}

func SearchSomething(s string) error {
	// assume that search code is written and we need to return an error
	// we can return QueryError as error because it implements error interface
	return &QueryError{
		Func:  "SearchSomething",
		Input: s,
		Err:   errors.New("not found"),
	}

}
