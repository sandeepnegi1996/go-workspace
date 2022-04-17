// Converting feets intp meters


package main
import "fmt"

func main() {

	fmt.Println("Enter number of feets")
	var input float64
	fmt.Scanf("%f",&input)

	output:=input * 0.3048
	fmt.Printf("Meter :  %f ",output)
}