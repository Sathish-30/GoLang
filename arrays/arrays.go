package arrays

import "fmt"

func SliceDeclaration() {
	//var nameList []string
	var nameList []string
	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Enter the name : ")
		fmt.Scan(&name)
		nameList = append(nameList, name)
	}
	for _, name := range nameList {
		fmt.Print(name + " ")
	}
}

func ArraysDeclaration() {
	//var names [4]string
	var names = make([]string, 4)
	for i := 0; i < 4; i++ {
		fmt.Print("Enter the names : ")
		fmt.Scan(&names[i])
	}
	for _, name := range names {
		fmt.Print(name + " ")
	}
}

func AddElementAtFirst(ele int) {
	var values []int
	for i := 0; i < 5; i++ {
		values = append(values, i+1)
	}
	values = append([]int{ele}, values...)
	//fmt.Print(len(values), "\n")
	for _, e := range values {
		fmt.Print(e, " ")
	}
}

func DeleteArrayAtIndex(index int, itr int) {
	var values []int
	for i := 0; i < itr; i++ {
		values = append(values, i+1)
	}
	values = append(values[:index], values[index+1:]...)
	for _, e := range values {
		fmt.Print(e, " ")
	}
}
