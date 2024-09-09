// Start from the 'func main'.
package main //main package is included.

import ( //here you will include the necessary libraries for specific operation. You can learn about each separately.
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func add(a int, b int) int { //this is the function to add two numbers(first start understanding from main...)
	return a + b
}

func swap(x, y, z string) (string, string, string) { //swipe two strings.
	return z, y, x
}

func split(sum int) (x, y int) { //spliting the two digit number. It take one number and returns two digits.
	x = sum * 4 / 9 //a simple logic to separate numbers.
	y = sum - x
	return
}

func breaking(number int) (x, y int) { //in Golang function can return more than one value.
	x = number / 10
	y = number % 10
	return
}

var C, python, Go, Java bool //declaring the variables of bool type.
var a, b int = 1, 2          //declaring variable .

var ( //another way to initialize variables.
	what    bool       = true
	Maxint  uint32     = 1<<32 - 1
	complex complex128 = cmplx.Sqrt(4 + 24i) //Sqrt is square root function in math.
)

const ( //declaring the constant variable.
	Big   = 1 << 100
	Small = Big >> 99
)

func getint(x int) int { //just of funtion.
	return x * 10
}
func getfloat(y float32) float32 {
	return y * 0.1
}

func sqrt(x float64) string { //This is the square root function it is not build-in.
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 { //this is power function not build-in.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g>=%g", v, lim)
	}
	return lim
}
func main() {
	fmt.Println("The random number is:", rand.Intn(99))         //random number generating anf printing.
	fmt.Printf("Now I have %g problems.\n", math.Sqrt(9999999)) //square root function.
	fmt.Println(math.Pi)                                        //print the value of pie.
	fmt.Println(add(9, 9))                                      //add function i wrote see above.
	fmt.Println(swap("Do", "or", "Die"))                        //see above.
	fmt.Println(split(64))                                      //split function written above.
	fmt.Println(breaking(24))                                   //break function is also written above.
	var i int                                                   //declaring integer type variable.
	Go = true                                                   //declaring outside of main(see above) here just initialized.
	fmt.Println(i, C, python, Go, Java)
	fmt.Println(a, b)
	var q, r = "Tom", "Jerry"
	fmt.Println(q, r)

	k := 18 //this syntax of initialization is always inside a function
	fmt.Println(k)
	//the variables written are already declared outside of main.
	fmt.Printf("The type is %T and its value is %v\n", what, what) //here %T will print the type of variable and %v gives its value.
	fmt.Printf("The type is %T and its value is %v\n", Maxint, Maxint)
	fmt.Printf("The type is %T and its value is %v\n", complex, complex)
	var y int
	var w float64
	var e bool
	var u string
	fmt.Printf("%v %v %v %q\n", y, w, e, u) //for u we have %q to display the value.
	fmt.Println(y, w, e, u)
	var m, n int = 3, 4
	var l float64 = math.Sqrt(float64(m*m + n*n))
	var g uint = uint(l) //unassigned int
	fmt.Println(m, n, g)
	j := 44.4
	fmt.Printf("j is of type %T\n", j)
	const hello = "Tom" //Constant
	const c = 33
	fmt.Println(hello, c)
	fmt.Println(getint(Small)) //these functions are written above.
	fmt.Println(getfloat(Small))
	fmt.Println(getfloat(Big))
	//for loop
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)
	s := 1
	for s < 1000 {
		s += s
	}
	fmt.Println(s)
	v := 1
	for v < 1000 {
		v += v
	}
	fmt.Println(v)

	//if statment up in function
	fmt.Println(sqrt(2), sqrt(-16)) //this function are mentioned above.

	//if with a short statement,else
	fmt.Println(pow(3, 2, 10), pow(3, 4, 25))
}

//this is another program where we make a logic to find square root called newton and Raphson method
/*
package main
import (
	"fmt"
	"math"
)
func sqrt(x float64) float64{
	z:=1.0
	for i:=0;i<10;i++{
		z-=(z*z-x)/(2*z)
		fmt.Printf("Guess %d: %.10f.\n",i+1,z)
	}
	return z
}

func main(){
	fmt.Println(sqrt(6))
	fmt.Println("The square root of 2 using build-in function:",math.Sqrt(6))
}
*/
