package main
import "fmt"

func main() {
	fmt.Println("1+1 = ",1+1)

	fmt.Println("1.0+1.0 = ",1.0+1.0)

	//length of string
	fmt.Println(len("test "))


	//string concetenation

	fmt.Println("hello 1 "+ "world 1")

	fmt.Println("hello"[2])     //108

	// l --> asci value is 108
	//each character is a byte 
	// 1 byte = 8 bit


	//Booleans  --> named after George Boole 
	// it is special 1 bit integer type 
	//logical operators are used with bool
	//&& ,|| , !
	fmt.Println((true && false) || (false && true) || !(false && false))

	
}