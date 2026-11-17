// Introduction to go programming language 
package main // this tells Go that the current file belongs to the main package. why main because go needs to know where to start executing programs.

import "fmt"
func main(){
    // Variable declarartion and intialization
    var name string = "Abebe"
    var age int = 23
    var country string = "Ethiopia"
    fmt.Println("Name : Abel Kebede")
    fmt.Println("Language:Go")
    fmt.Println("Goal:Become a Go Developer")
    fmt.Println("I am learning GO!")
    fmt.Println("My name is ", name , "I am ", age, "years old and I am from ", country)
   

    // Method two of declaring a variable 
    var nname = "Kiros"
    fmt.Println("My name is ",nname)
    
    // Method three of declaring variable in Go
    isStudent := true
    firstName := "Abeba"
    lastName := "Aba"
    fmt.Println(isStudent)

// Understanding Go data types
// 1,String

// Combining strings
fullName := firstName + ""+ " " + lastName
fmt.Println(fullName)

// 2, Integer types
var ageOne int = 23
fmt.Println(ageOne)




}

// func main  tells go where execution begin.
// Exercise one 

// A package a budled programming code that can be used by other programs. A package is a collection of source files in the same directory that are compiled together. Each source file starts with a package declaration, which defines the package name. The main package is a special package that tells Go where to start executing the program.
// Module a single file any programming language file like python .py.


