package main
import ("fmt")

func main(){

// Use the switch statement to select one of many code blocks to be executed.

/* switch expression {
	case x:
   // code block
	case y:
   // code block
	case z:
	...
	default:
   // code block
} */

day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  }
}






