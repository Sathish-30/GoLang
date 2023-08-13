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
	var names = [4]string{}
	for i := 0; i < 4; i++ {
		fmt.Print("Enter the names : ")
		fmt.Scan(&names[i])
	}
	for _, name := range names {
		fmt.Print(name + " ")
	}
}
