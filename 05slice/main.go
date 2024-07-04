package main
import ("fmt")
	
// Slices are similar to arrays, but are more powerful and flexible.
// Syntax ----  slice_name := [] datatype {values}

/* len() function - returns the length of the slice (the number of elements in the slice)
    cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to) */

/* func main(){
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := [] string {"go", "slices", "python"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
} */

// Create a Slice From an Array
// syntax -- var myarray = [length] datatype {values} // An array
//           myslice := myarray [start:end]           // A slice made from the array

/* func main(){
	var myarray = [6] int {1,2,3,4,5,6}
	myslice := myarray [1:2]
	fmt.Printf("myslice = %v \n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
} */

 /* func main(){
	//The make() function can also be used to create a slice.
	// syntax --- slice_name := make([]type, length, capacity)
  	myslice1 := make( [] int, 5, 10)
  	fmt.Printf("myslice1 = %v\n", myslice1)
  	fmt.Printf("length = %d\n", len(myslice1))
  	fmt.Printf("capacity = %d\n", cap(myslice1))

 	 // with omitted capacity
 	 myslice2 := make( [] int, 5)
  	fmt.Printf("myslice2 = %v\n", myslice2)
  	fmt.Printf("length = %d\n", len(myslice2))
  	fmt.Printf("capacity = %d\n", cap(myslice2))
} */

func main(){

	var fruitlist = [] string {"apple", "tomoto", "peach"}
	fmt.Printf("Which type is this is: %T \n", fruitlist)

	// You can append elements to the end of a slice using the append()function:
	// syntax --- slice_name = append(slice_name, element1, element2, ...)

	fruitlist = append(fruitlist, "mango","orange","banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist [1:3])
	fmt.Println(fruitlist)

	// how to remove a value from slices based on index
	var courses = [] string {"go","javascript","swift"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)


}



