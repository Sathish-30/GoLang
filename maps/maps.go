package maps

import "fmt"

func MapInitialization() {
	studentIdWithName := map[int]string{}
	mapDetailOfStudent := &studentIdWithName
	fmt.Printf("%p\n", &studentIdWithName)
	studentIdWithName[19] = "kaarthika"
	studentIdWithName[47] = "sathish"
	studentIdWithName[43] = "sandi"
	studentIdWithName[42] = "ratish"
	studentIdWithName[62] = "Chittap"
	studentIdWithName[7] = "Dhana"
	for key, value := range studentIdWithName {
		fmt.Printf("The Key and value pair are : %v - %v \n", key, value)
	}
	fmt.Printf("%p\n", &studentIdWithName)
	fmt.Print(*mapDetailOfStudent)

}
