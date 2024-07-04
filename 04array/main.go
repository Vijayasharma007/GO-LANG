package main
import ("fmt")

//var array_name = [length] datatype {values} // here length is defined
//array_name := [length] datatype {values} // here length is defined


/* func main(){
	var arr1 = [5] int {1,2,3,4,5}
	arr2 := [5] int {4,5,6,7,8}
	fmt.Println(arr1)
	fmt.Println(arr2)
} */

/* func main(){
	var cars = [4] string {"volvo","BMW","ford","mazda"}
	fmt.Println(cars)
	fmt.Println(cars[0])
	fmt.Println(cars[3])
} */

/* If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type. */
// The default value for int is 0, and the default value for string is "".

 /* func main(){
	arr1 := [5] int {} //not initialized
	arr2 := [5] int { 1,2 } //partially initialized
	arr3 := [5] int { 1,2,3,4,5} //fully initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
} */

/* func main(){
	fmt.Println("welcome to array in golang")

	var fruitlist [2] string

	fruitlist[0]= "apple"
	fruitlist[1]= "mango"
	fmt.Println("fruit list is:", fruitlist)
	fmt.Println("fruit list is:", len(fruitlist))

} */

func main(){
	//Array declaration basic form
	var a [5] int
	fmt.Println(a)

	//Array declaration with element
	var b = [5] int {1,2,3,4,5}
	fmt.Println(b)

	c := [5] int {6,7,8,9,0}
	fmt.Println(c)

	//sparse array declaration 
	d := [5] int {1, 2: 10, 2, 3}
	fmt.Println(d)

	//implicit length declaration
	e := [...] int {1,2,3,4,5}
	fmt.Println(e)
	fmt.Println("length of the array is:", len(e)) //find the length of the array

	f := e[3]
	fmt.Println(f)

	h := [2][2] int {{1,2},{3,4}}
	fmt.Println(h)

}

