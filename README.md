### GoLang

<br/>

> Installation done from - https://go.dev/ , GoLang is an open-source programming language supported by Google

- Go Lang mainly requires packages , which can be declared as main.

- It also requires a main method where it can start the program execution.

- Where it's functions are mainly retrieved from the packages , there are many packages (fmt) is the package for getting the print method.

- If we declare a varibale but not used the variable it will throw a error

- `const conferenceTickets = 30` , In const keyword we can't change the value

- Print format(Printf) function is used to print using format specifier

- %v is used to display the value in a default format

- %#v a Go-syntax representation of the value

- %T a Go-syntax representation of the type of the value

- uint can't be assigned with a negative number

- Like assigning `var name = "Sathish" ` we can assign as `name := "Sathish"`

- Pointers are said to be as special variables
  <br />

#### Arrays

- `var bookings = [10]string{"Sathish","sandi","karthi","chittappa","ratish"}` The following code is the static initialization of array

- Other way of array initialization `var bookingNames [10]string`
  <br />

##### Slice

- Slice is an abstraction of an array , it is also index based . Slices are more flexible and powerful :variable length or to get the sub array of it own

- To declare a slice `var names []string` , it is same as the declaration of the array but the no of elements are not specified

- To append or add value is done by `names = append(names , "Sathish")`

#### Noteable Statement

- There is only one loop that is 'for' loop

- \_underscore is used to ignore the variable

- Conditional statement are same as in the other programming language

- Switch statement is used by checking it with the corresponidng case statement

#### Functions

> `func add(x int, y int) int { return x + y }`

- The syntax of func keyword describing that given is a function along with the function name and () paranthesis which will have empty paramater or a value parameter then the return type

- Notes :- Function is only executed , when "called!",
  you can call a function as many times you want , so a function is also used to reduce code duplication

- We can return multiple elements

#### Package Level Variable

- variable declared outside the function within package , where it can't be decalared with `names := "sathish"` , where it can be declared with const and var keyword

#### Packages

- A package is collection of Go files
