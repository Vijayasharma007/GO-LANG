package main
import ("fmt")

func main(){

	// Loop are handy if you want to run the same code over and over again, each time with a different value.
	// Each execution of a loop can take up to three statement
	/* syntax
	
	for statement1; statement2, statement3 {
	// block of codes....
	...
	} */

	for i:=0; i<5;i++{
		fmt.Println(i)
	}
	

	/* for i:=0; i < 5; i++ {
		if i == 3 {
		  continue
		}
	   fmt.Println(i)
	  }
    } */

	/* days := [] string {"sunday", "monday", "tuesday", "wednesday", "thursday","friday","saturday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	} */



	/* The range keyword is used to more easily iterate over an array, 
	slice or map. It returns both the index and the value. */

	/* syntax
	    for index, value := array|slice|map {
   	    code to be executed for each iteration
        } */
	
	/* func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
	 fmt.Printf("%v\t%v\n", idx, val)
	}
  } */