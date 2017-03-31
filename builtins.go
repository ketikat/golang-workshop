package workshop

import "fmt"

func builtins() {
	var slice1 = []string{
		"I am a string",
		"I am the other string",
	}
	slice := append(slice1, "Just another string")
	const capLength = 9
	sliceFinal := make([]string, len(slice), capLength)
	copy(sliceFinal, slice)
	comparator := func(slice, slicefinal []string) bool {
		fmt.Println(slice, sliceFinal)
		for i := range slice {
			if slice[i] != slicefinal[i] {
				return false
			}
		}
		return true
	}
	// Cannot compare slices with == like you can with arrays
	assert(comparator(slice, sliceFinal) == true)

	capacity := cap(sliceFinal)
	assert(capacity == capLength)

	length := len(sliceFinal)
	assert(length == 3)

	list := make(map[string]int)
	list["num1"] = 5
	list["num2"] = 3

	delete(list, "num1")
	assert(len(list) == 1)

	// new returns a pointer to the newly allocated type
	names := new([]string) // return &[]
	fmt.Println(names)

	// since new returns a pointer you must use the deference operator
	num := new(int)
	*num++
	*num--
	assert(*num == 0)
}