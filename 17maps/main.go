package main
import ("fmt")

// var a = map[KeyType] ValueType {key1:value1, key2:value2,...}
// b := map[KeyType] ValueType {key1:value1, key2:value2,...}

/* func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
  
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
  } */

func main(){
	languages := make(map[int]string)

	languages[1] = "Golang"
	languages[2] = "python"
	languages[3] = "javascript"
	languages[4] = "java"

	fmt.Println("Most popular programming languages:", languages)
}