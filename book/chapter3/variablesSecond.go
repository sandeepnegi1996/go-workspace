package main
import "fmt"
func main() {

	//values assigned to the variables can be reassigned
	var x string
	x="first"
	fmt.Println(x)

	x="second"    // this is variable reassignment
	fmt.Println(x)

	// == 

	var t string="hello"
	var k string ="hello"

	fmt.Println(t==k)

	//shorter form to create and assign variable

	name:="sandeep"  
	fmt.Println(name)

	counter:=1
	fmt.Println(counter)
}