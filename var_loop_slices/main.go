// Online Go compiler to run Golang program online
// Print "Hello, World!" message

package main
import "fmt"
func main() {
    var firstname string = "anamika"
    var lastname string ="saxena"
    var age uint = 23
    var email string = "anu@gmail.com"
    fmt.Printf("nice to meet you !  this is %v %v and i am %v years old my emal address is %v , please contact by using my mail . thanku  ", firstname, lastname, age,email)
    fmt.Printf("\nthe datatypes of all variavles are %T ,%T ,%T and %T", firstname,lastname,age, email)
    a := 10
    b := 3
    fmt.Printf("\naddition is  %v", a+b)
    fmt.Printf("\nmultiplication is %v", a*b)
    fmt.Printf("\n subtraction is  %v ", a-b)
    fmt.Printf("\n division is %v ",a/b )
    var marks int 
    fmt.Scan(&marks)

switch {
case marks >= 90:
    fmt.Printf("\nthe grade is A+")

case marks >= 80 && marks < 90:
    fmt.Printf("A")

case marks >= 70 && marks < 80:
    fmt.Printf("B")

case marks >= 60 && marks < 70:
    fmt.Printf("C")

default:
    fmt.Printf("fail ho gaya tera bachha chhi !!!")
}
var k int 
    fmt.Scan(&k)
    fmt.Printf("factorial of %v is %v", k, fact(k))


    numbers := []int{10, 20, 30, 40, 50}

    index := 2

    numbers = append(numbers[:index], numbers[index+1:]...)

    fmt.Println(numbers)
    mark := map[string]int{
    "A": 80,
    "B": 95,
    "C": 72,
    "D": 88,
}
    mark["E"]=90
    mark["C"]=78
    delete(mark, "B")
    fmt.Println(mark)
}


func fact(a int) int {
    if a == 1 || a == 0 {
        return 1
    }

    return a * fact(a-1)
}
    
    
    

