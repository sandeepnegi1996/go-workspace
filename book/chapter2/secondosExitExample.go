package main
import (
	"fmt" ;
	"os"
)

func main() {
	
	for  i := 0; i <10 ;i++ {
	
		fmt.Printf("[%d] hello \n",i)
		if(i==5) {
			fmt.Printf("Bye")
			os.Exit(0)
		}
}
}

