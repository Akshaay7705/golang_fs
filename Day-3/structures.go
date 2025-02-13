package main
import "fmt"

type student struct{
	Name string
	RgNo float64
	Dept string
}

func main(){
	st1 := student{Name : "Stud1", RgNo : 13.5, Dept : "CS"}
	fmt.Println("Name : ", st1.Name, "Regno : ", st1.RgNo, "Dept :",st1.Dept)
}