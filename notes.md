# Notes

## General

- Every go repo needs a `go.mod` file which defines module name, go version and dependencies
- All code must belong to a package
- First statement in a go file must be `package ...`
- Default package should be `package main`
- `main` function is the entry point of a go program

## Variables & Constants

### Variables

- `var myVariables type = "Hello World"`
- Shorthand for defining variables `myVariable := "Hello World"`
  - No explicit type declaration
  - Not usable for constants

### Constants

- `const myConstant type = "The only constant thing in life is change"`

### Package-level variables

- Package-level variables hold shared state
- Created at the top outside all functions
- Accessible inside any function
- Accessible **in all files inside the same package**
- Can't be defined using the `:=` shorthand

## Formatting Strings

- `fmt.Println("My variable has a value of", myVariable)` takes multile arguments and separates them with spaces
- `fmt.Printf("Use placeholders %v to replace variables", myVariable)` uses placeholders and applies placeholder-specific formatting. See https://pkg.go.dev/fmt

## Data Types

- Statically typed with type inference

## User Input

`fmt.Scan(&name)` requires a pointer and stores the input at that address

## Arrays & Slices

### Arrays

- Arrays have a fixed size
- `var myArray = [50]string{}`
- Shorthand for empty arrays: `var myArray [10]int`

### Slices

- Abstraction of an array
- Variable length
- Also index-based and have a size but automatically resized when needed
- `var mySlice = []string{}`
- Shorthand:
  - `var mySlice []string`
  - `mySlice := []string{}`
- Built-in `append(mySlice, element)` function to add elements to the end of the slice

## Maps

- `map[string]int` defines the types of the map keys and values
- Use `make(map[string]int)` to create an empty map
- Add data using `myMap[key] = value`

## Loops

- ONLY "for loop" for all use cases
- No while, do-while or forEach keywords

### `while`

```go
for condition {
  // Your code
}
```

### Iterating through a list

```go
for index, element := range mySlice {
  // Your code
}

// Use _ as blank identifier to ignore the index
```

## Conditional logic

```go
if condition {

} else if condition {

} else {

}
```

```go
switch variable {
  case value1:
    // Code
  case value2:
    // Code
  case value3, value4:
    // Code
  default:
    // Code
}
```

## Functions

```go
func myFunction (input1 string, input2 int) bool {
  // Code
  return myBoolean
}

// You can also return multiple values
func myFunction (input string) (string, int, bool) {
  // Code
  return myString, myInt, myBool
}
```

## Packages

- Importing files from other packages in the same module:

```go
import (
  "moduleName/packageName"
)
```

- Exporting functionality is done by simply capitalizing the function/variable name

```go
func ValidateUserInput() {
  // Code
}
```

## Concurrency

- Go actually doesn't create real OS threads but rather a virtual thread (aka green thread) called **goroutine**
- More lightweight and faster
- Easy communication between goroutines with **channels**
- Use the `go` keyword to start a goroutine

```go
bookTicket(numberOfTickets, firstName, lastName, email)
go  sendTicket(numberOfTickets, firstName, lastName, email)
```

- Synchronizing multiple threads works with `sync.WaitGroup{}`
- Call `wg.Add(2)` before starting a new threads to add the number of threads you started
- Call `wg.Done()` inside the multi-threaded function when it is done
- Call `wg.Wait()` at the end of the `main` function to wait for all running threads

## Random

- `strings.Fields()` splits a string with white-space as separator and returns a slice
- `strconv` has utilities for converting data types into strings
