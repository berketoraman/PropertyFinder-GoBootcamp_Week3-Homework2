package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

//Checks if the rectangle is valid.
func (instance *Rectangle) Info() {
	if instance.length <= 0 {
		fmt.Println("Error. The length of rectangle can not be 0 or negative.")
	}else if instance.breadth <= 0 {
		fmt.Println("Error. The breadth of rectangle can not be 0 or negative.")
	}
}


func main() {

	//Creating instances of Rectangle and experimenting on them.
	var rect1 = Rectangle{10, 20}
	fmt.Println("The sizes of the rectangle is: ",rect1)
	fmt.Println("The area of given rectangle is: " ,area(rect1))
	fmt.Println("The Circumference of given rectangle is: ",circumference(rect1))
    rect1.Info()// checks if the rectangle is valid

	var rect2 = Rectangle{length: 20, breadth: 40} 
	fmt.Println("\nThe sizes of the rectangle is: ",rect2)
	fmt.Println("The area of given rectangle is: " ,area(rect2))
	fmt.Println("The Circumference of given rectangle is: ",circumference(rect2))
	rect2.Info()// checks if the rectangle is valid

	rect3 := Rectangle{20, 30}
	fmt.Println("\nThe sizes of the rectangle is: ",rect3)
	fmt.Println("The area of given rectangle is: " ,area(rect3))
	fmt.Println("The Circumference of given rectangle is: ",circumference(rect3))
	rect3.Info()// checks if the rectangle is valid

	rect4 := Rectangle{length: 50, breadth: 30}
	fmt.Println("\nThe sizes of the rectangle is: ",rect4)
	fmt.Println("The area of given rectangle is: " ,area(rect4))
	fmt.Println("The Circumference of given rectangle is: ",circumference(rect4))
	rect4.Info()// checks if the rectangle is valid

	rect5 := Rectangle{length: 20, breadth: 30}
	fmt.Println("\nThe sizes of the rectangle is: ",rect5)
	fmt.Println("The area of given rectangle is: " ,area(rect5))
	fmt.Println("The Circumference of given rectangle is: ",circumference(rect5))
	rect5.Info() // checks if the rectangle is valid


	rect6 := Rectangle{breadth: 10}
	fmt.Println("\nThe sizes of the rectangle is" ,rect6)
	fmt.Println("The area of given rectangle is: " ,area(rect6))
	fmt.Println("The Circumference of given rectangle is: ",circumference(rect6))
	rect6.Info()// checks if the rectangle is valid

	rect7 := Rectangle{length: 50}
	fmt.Println("\nThe sizes of the rectangle is" ,rect7)
	fmt.Println("The area of given rectangle is: " ,area(rect7))
	fmt.Println("The Circumference of given rectangle is: ",circumference(rect7))
	rect7.Info() // checks if the rectangle is valid
}

//Area function
func area(r Rectangle) int{
   if r.length <= 0 {
	   return 0
   }else if r.breadth <= 0 {
	   return 0
   }
	
   return  r.length * r.breadth
}

//Circumference function
func circumference(rect Rectangle) int {
	if rect.length <= 0 {
		return 0
	}else if rect.breadth <= 0 {
		return 0
	}

	return  2 * (rect.length + rect.breadth)
}





