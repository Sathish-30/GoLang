package concepts

import "fmt"

type MyInt = int32
type Celcius float64
type Fahrenhiet float64


// Where in type alias it doesn't change the datatype to the Name given by the users , instead it take the primitive data
func TypeAlias(){
	var num MyInt
	num = 4121
	fmt.Printf("The dataType of the num is %T\n",num)
}

func CheckType() {
	c := Celcius(11.2)
	fahrenhiet := c.calculate()
	fmt.Printf("The fahrenhiet for %v celcius is %v\n",c,fahrenhiet)
}

func (c Celcius) calculate() Fahrenhiet {
	return Fahrenhiet(c*(9/5) + 32)
}