package gtb

import (
	"fmt"
	"strconv"
)

// SelectFromArray : select an element in the provided array
func SelectFromArray(a []string) string {
	fmt.Println("Choices:")
	for i := range a {
		fmt.Println("[", i, "]: "+a[i])
	}
	fmt.Println("Enter Number of Selection: ")
	sel, err := strconv.Atoi(GetInput())
	EoE(fmt.Errorf("Error Getting Integer Input from User: %w", err))
	if sel <= len(a)-1 {
		return a[sel]
	}
	return SelectFromArray(a)
}

// SelectFromMap : select an element in the provided map
func SelectFromMap(m map[string]string) string {
	fmt.Println("")
	fmt.Println(MapToString(m))
	fmt.Printf(":")
	sel := GetInput()
	if _, found := m[sel]; found {
		return sel
	}
	fmt.Printf("\"%v\" is an Invalid Selection!!\n", sel)
	fmt.Printf("Please Select From the Following:\n")
	return SelectFromMap(m)
}
