#### golang-diary
Golang Data structure  
### The most general form to declare a variable in Golang uses the var keyword.
var i int = 10
var j string = "shashank"

var i int
var s string
i = 10
j = "shashank"
### Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
### Package reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types.extract its dynamic type information by calling TypeOf, which returns a Type.
reflect.TypeOf(i)


### The <strong>:=</strong> short variable assignment operator indicates that short variable declaration is being used. There is no need to use the var keyword or declare the variable type.


## Variable Scope
### Inner block can access its outer block defined variables, but outer block cannot access inner block defined variables.
### Variables declaration can be grouped together into blocks for greater readability and code quality.
var (name  = "shashank"
    qualifiction = Btech
    )

## Constant
## you assign a value at the time of the constant declaration, you can't assign a value later as with variables
### By convention, constant names are usually written in uppercase letters.

