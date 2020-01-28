package ch1

import "fmt"

func NumericTypesConversion() {
	myInt := 2
	myFloat := 1.5
	myOtherFloat := float64(myInt) * myFloat
	fmt.Println(myOtherFloat) // Outputs 3
}

func ArraysAndSlices() {
	mySlice := []int{1, 2, 3, 4, 5}
	mySlice2 := mySlice[0:3]
	mySlice3 := mySlice[1:4]
	fmt.Println(mySlice2, mySlice3, mySlice[2])
	// Outputs: [1 2 3] [2 3 4] 3

}
func ArraysAndSlices2() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(mySlice)) // Outputs: 5
	mySlice2 := []string{"Hi", "there"}
	fmt.Println(len(mySlice2)) // Outputs: 2
}

func Looping() {
	animals := []string{"Cat", "Dog", "Emu", "Warthog"}
	for i, animal := range animals {
		fmt.Println(animal, "is at index", i)
	}
}
func Looping2() {
	// 用下划线丢弃某个变量
	animals := []string{"Cat", "Dog", "Emu", "Warthog"}
	for _, animal := range animals {
		fmt.Println(animal)
	}
}

// Maps 代表其他语言中的哈希表 字典
func Maps() {
	// maps are not output in any particular order.
	starWarsYears := map[string]int{
		"A New Hope":              1977,
		"The Empire Strikes Back": 1980,
		"Return of the Jedi":      1983,
		"Attack of the Clones":    2002,
		"Revenge of the Sith":     2005,
	}
	// add
	starWarsYears["The Force Awakens"] = 2015

	fmt.Println(len(starWarsYears)) // Correctly outputs: 6

	// Looping over Maps
	for title, year := range starWarsYears {
		fmt.Println(title, "was released in", year)
	}
}

func Map2() {
	colours := map[string]string{
		"red":     "#ff0000",
		"green":   "#00ff00",
		"blue":    "#0000ff",
		"fuchsia": "#ff00ff",
	}
	redHexCode := colours["red"]
	fmt.Println("redHexCode: ", redHexCode)

	// built-in delete function

	delete(colours, "fuchsia")

	// key exists
	code, exists := colours["burgundy"]
	if exists {
		fmt.Println("Burgundy's hex code is", code)
	} else {
		fmt.Println("I don't know burgundy")
	}
}
