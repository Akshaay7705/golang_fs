package main
import "fmt"


func main(){
	// ifelsemod()
	// forCondional()
	forArray()
}
func ifelse(){
	var age int
	fmt.Print("Enter your age : ")
    fmt.Scan(&age)
	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}

func ifelsemod(){
	var a, b int
	fmt.Print("Enter 1st num : ")
    fmt.Scan(&a)
	fmt.Print("Enter 2nd num : ")
    fmt.Scan(&b)
	if a > b {
		fmt.Println("a is greater")
	} else if a == b {
		fmt.Println("a is equal to b")
	} else {
		fmt.Println("b is greater")
	}
}


func forCondional(){
	n := 1
	for n < 5{
		n *= 2
	}
	fmt.Println(n)
}

func forArray() {
	strings := []string{"hello", "world", "Golang", "NIE"}
	for _,s := range strings {
		fmt.Println(s) 
	}
}