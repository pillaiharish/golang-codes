package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"sort"
	"strconv"
)

func main() {
	// 1) := vs =
	// := declares variables.
	x := 10
	// = assigns to existing variables.
	x = 20
	fmt.Println("Value of x is:", x)

	// The important trap:
	var1 := 10
	if true {
		var1 := 20
		fmt.Println("value of var1 is:", var1)
	}
	fmt.Println("value of var1 is outside if is:", var1)

	// The inner var1 is a new variable because of lexical scope.
	// This is called shadowing.

	//But this is valid:
	a := 10
	a, b := 20, 30

	// in above atleast one variable is new.
	fmt.Println("Value of a is:", a, "value of b is:", b)

	// rule:
	// := requires at least one new variable in the same scope.

	// 2) Unused variables are compile errors
	// Go requires local variables to be used.

	// Same for imports.
	// If unused we can get compile errors

	// We can intentionally discard values using _.
	var2, _ := strconv.Atoi("123")
	fmt.Println("Value of var2:", var2)

	// 3) Multiple return values extremely common in Go:
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error is:", err, "Result:", result)
	}

	// 4) `if` can have initialization
	if result, err = divide(50, 0); err != nil {
		fmt.Println("Error is:", err, "Result:", result)
	}

	if result, _ = divide(150, 50); result > 2 {
		fmt.Printf("Result %d is greater than 2\n", result)
	}

	// 5) for is Go's only loop

	// classic for loop
	for i := 0; i < 2; i++ {
		fmt.Println("Inside classic for loop", i)
	}

	// equivalent to while
	i := 1
	for i > 0 {
		fmt.Println("Inside `while` style for loop", i)
		break
	}

	// infinite for loop
	for {
		fmt.Println("Inside infinite for loop", i)
		break
	}

	// inside range for loop
	arr := []int{3, 5}
	for j, k := range arr {
		fmt.Println("Inside range for loop with index", j, "and value", k)
	}

	// ignore index and capture value only
	for _, j := range arr {
		fmt.Println("Inside range for looping arr for values", j)
	}

	// only index capture
	for j := range arr {
		fmt.Println("Inside range for looping only index", j)
	}

	// 6) switch does NOT fall through by default
	d := 1
	switch d {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}
	// the above prints `one` as output

	switch d {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
	}
	// Since we used `fallthrough` this will print both `one` and `two` in separate lines

	// 7. Arrays are values in Go
	arr1 := [3]int{3, 5, 7}
	fmt.Println("Array1 value is", arr1)
	arr2 := arr1
	arr2[0] = 100
	fmt.Println("Array2 was assigned Array1 and first element changed to 100")
	fmt.Println("Array2 value is", arr2)
	fmt.Println("arr2:=arr1 copies the entire array.")

	// Also: [2]int and [3]int are different types.
	fmt.Println("We cannot use [3]int := [2]int array of different sizes will not work")
	// arr3 := [2]int{}
	// arr3 = arr1
	// fmt.Println("Array3 will not work",arr3)

	// 7.1) Now the important part: Slices
	// A slice is not an array.

	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array3 is an array ", arr3)

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice1 is %d a slice\n", slice1)

	// Conceptually:
	// type slice struct {
	// 	ptr *T
	// 	len int
	// 	cap int
	// }

	// Think of a slice approximately as:
	// +--------------------+
	// | pointer to array   |
	// | length             |
	// | capacity           |
	// +--------------------+

	// 8.) len vs cap

	s := make([]int, 3, 5) // len = 3, cap = 5
	fmt.Println("Length of slice s:", len(s))
	fmt.Println("Capacity of slice s", cap(s))
	fmt.Println("Default slice s contents:", s)
	// for i := range cap(s) {
	// 	s[i] = i
	// }
	// fmt.Println("iterating through cap=5")
	// fmt.Println("slice s contains", s)

	// -----------------------
	// Output:
	// Default slice s contents: [0 0 0]
	// panic: runtime error: index out of range [3] with length 3

	// goroutine 1 [running]:
	// main.main()
	// 		~/golang-codes/go-syntax-traps/main.go:163 +0xabc
	// exit status 2
	// ------------------------

	// Underlying storage conceptually:
	// [0 0 0 _ _]
	//  ↑
	//  len = 3
	//  cap = 5

	// We can access:
	// s[0], s[1] and s[2]
	// But: s[3] panics even though capacity is 5.
	// Why?
	// Because indexing is limited by length, not capacity.

	// 9.) make([]int, 3) does NOT create an empty slice
	s2 := make([]int, 4) // len = 4, cap = 4
	fmt.Println("Print slice s2", s2, "of size", len(s2))
	// Output: Print slice s2 [0 0 0 0] of size 4

	// append a value 10
	s2 = append(s2, 10)
	fmt.Println("After append 10 to slice s2", s2)
	// Output: After append 10 to slice s2 [0 0 0 0 10]

	// People sometimes intend:
	s3 := make([]int, 0, 3) // len = 0, cap = 3
	s3 = append(s3, 11)
	fmt.Println("After append 11 to slice s3", s3)
	// Output: After append 11 to slice s3 [11]

	// 10.) Slice shares underlying arrays
	// This is one of the most important Go concepts.
	// so i can clone that data to next slice2 and if data in slice2
	// changes then data at origin slice1 also changes.
	slice2 := []int{20, 30, 40, 70}
	slice3 := slice2[2:3]
	fmt.Println("slice3 value is", slice3)
	slice3 = slice2[2:] // did not copy the values.
	// It references the same backing array.
	fmt.Println("slice3 value is for [2:] is", slice3)
	slice3[0] = 114
	fmt.Println("slice2 after changing first element of slice3[0]=114", slice2)
	// Output:
	// slice3 value is [40]
	// slice3 value is for [2:] is [40 70]
	// slice2 after changing first element of slice3[0]=114 [20 30 114 70]

	s4 := []int{4, 6, 9, 1}
	s5 := s4
	s5[0] = 123
	fmt.Println(s4)
	fmt.Println(s5)
	// Output:
	// [123 6 9 1]
	// [123 6 9 1]

	// 11.) Slicing syntax
	// s[a:b]
	// includes a, excludes b.

	// 12.) Slice capacity after slicing
	slice4 := []int{20, 44, 51, 17, 23, 68}
	slice5 := slice4[1:3]
	fmt.Println("slice5 := slice4[1:3]", slice5)
	fmt.Println("Len of slice5", len(slice5))
	fmt.Println("cap of slice5", cap(slice5))
	slice6 := slice4[3:6]
	fmt.Println("slice6 := slice4[3:6]", slice6)
	fmt.Println("Len of slice6", len(slice6))
	fmt.Println("cap of slice6", cap(slice6))
	// Output:
	// slice5 := slice4[1:3] [44 51]
	// Len of slice5 2
	// cap of slice5 5
	// slice6 := slice4[3:6] [17 23 68]
	// Len of slice6 3
	// cap of slice6 3

	// 13.) append may or may not allocate
	// Consider:
	slice7 := make([]int, 2, 4)

	slice7[0] = 10
	slice7[1] = 20

	slice8 := append(slice7, 30)

	// There is available capacity.
	// So b probably uses the same backing array.
	// Conceptually:
	// slice7 -> [10 20 _ _]
	// slice8 -> [10 20 30 _]

	// Now:
	slice8[0] = 999
	fmt.Println("slice8 first value is", slice8[0])
	// Output:
	// slice8 first value is 999

	// because backing storage is shared.
	// But once capacity is exhausted:
	slice9 := []int{1, 2}

	slice10 := append(slice9, 3)

	// Go may allocate a new backing array.
	// Then modifications to b may no longer affect a.
	// Important rule:
	// append() can change where a slice points.

	// Therefore always capture its return value:
	// s = append(s, value)

	// Don't write:
	// append(s, value)

	// In fact, Go won't let we discard it like that as a standalone call.
	fmt.Println("slice9 is", slice9)
	fmt.Println("After append slice10 with slice9 and value=3 is:")
	fmt.Println("slice10 is", slice10)
	// Output:
	// slice9 is [1 2]
	// After append slice10 with slice9 and value=3 is:
	// slice10 is [1 2 3]

	// 14.) Slice function trap
	fmt.Println("----------------------")
	fmt.Println("14.) Slice function trap")
	slice11 := []int{1, 2}
	addSlice1(slice11)
	fmt.Println("slice11 after appending 100 inplace", slice11)
	slice11 = addSlice2(slice11)
	fmt.Println("slice11 after appending 101", slice11)
	// however modifying slice inside func without returning it still works.
	// The copied slice header still points to the same backing array.
	modifySlice1(slice11)
	fmt.Println("slice11 modify inplace index 0", slice11)
	// Output:
	// 14.) Slice function trap
	// slice11 after appending 100 inplace [1 2]
	// slice11 after appending 101 [1 2 101]
	// slice11 modify inplace index 0 [145 2 101]

	// 15. nil slice vs empty slice
	fmt.Println("----------------------")
	fmt.Println("15. nil slice vs empty slice")
	var slice12 []int
	slice13 := []int{}

	fmt.Println("slice12: ", slice12, "slice13: ", slice13)
	fmt.Println("length of slice12:", len(slice12), "and slice13:", len(slice13))
	fmt.Println("But slice12==nil gives:", slice12 == nil)
	fmt.Println("slice13==nil gives:", slice13 == nil)
	// Output:
	// 15. nil slice vs empty slice
	// slice12:  [] slice13:  []
	// length of slice12: 0 and slice13: 0
	// But slice12==nil gives: true
	// slice13==nil gives: false

	// Usually both work similarly, but serialization can differ.
	// For example JSON commonly distinguishes them:
	slice12Struct := sliceExample{SliceKey: slice12}
	slice12JSON, err := json.Marshal(slice12Struct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Serialized slice12 JSON gives:", string(slice12JSON))

	slice13Struct := sliceExample{SliceKey: slice13}
	slice13JSON, err := json.Marshal(slice13Struct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Serialized slice13 JSON gives:", string(slice13JSON))
	// Output:
	// Serialized slice12 JSON gives: {"Slice":null}
	// Serialized slice13 JSON gives: {"Slice":[]}

	// nil slice   -> null
	// empty slice -> []

	slice12 = append(slice12, 101)
	slice13 = append(slice13, 172)
	fmt.Println("After append slice12", slice12)
	fmt.Println("After append slice13", slice13)
	// Output:
	// After append slice12 [101]
	// After append slice13 [172]

	// 16. We cannot compare slices
	fmt.Println("----------------------")
	fmt.Println("16. We cannot compare slices")
	// This is illegal:
	// a := []int{1, 2}
	// b := []int{1, 2}
	// fmt.Println(a == b)
	// This will give Compile error.

	// Slices can only be compared to:
	if slice13 != nil {
		fmt.Println("slice13 is not nil")
	}
	// For modern Go, we can use:
	fmt.Println("slice.Equal() output:", slices.Equal(slice13, slice12))
	compareSlices1, err := sliceEqual(slice13, slice12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("slice comparison function sliceEqual() output:", compareSlices1)

	// Output:
	// 16. We cannot compare slices
	// slice13 is not nil
	// slice.Equal() output: false
	// Slice contents not equal at index 0
	// slice comparison function sliceEqual() output: false

	fmt.Println("----------Maps------------")
	fmt.Println("17. Reading nonexistent map keys doesn't panic")
	map1 := map[string]int{"alice": 10, "bob": 20}
	// or
	map2 := make(map[string]string)
	map3 := make(map[string]bool)

	fmt.Println("map1 value is int output:", map1["Hello"])
	fmt.Println("map2 value is string output:", map2["Hello"], "<-empty string")
	fmt.Println("map3 value is bool output:", map3["Hello"])

	// These above results give ambiguity when key is not present, it will return
	// default result. for example: below "bob" key is not at all present
	// but still output will return false
	map4 := map[string]bool{"alice": true}
	fmt.Println("map4 without key value is bool output:", map4["bob"])

	// Output:
	// 17. Reading nonexistent map keys doesn't panic
	// map1 value is int output: 0
	// map2 value is string output:  <-empty string
	// map3 value is bool output: false
	// map4 without key value is bool output: false

	fmt.Println("----------------------")
	fmt.Println("18. Comma-ok idiom")
	// map returns false if the key does not exist using that extra return data we can
	// distinguish if a key is absent.
	val1, ok := map4["tom"]
	fmt.Println("In map4 key named tom present:", ok, ", its value:", val1)

	// Output:
	// 18. Comma-ok idiom
	// In map4 key named tom present: false , its value: false

	fmt.Println("----------------------")
	fmt.Println("19. Nil map: read okay, write panic")
	var map5 map[string]int
	fmt.Println("Read from nil map5:", map5["hello"])
	fmt.Println("Writing to nil map will throw panic")
	// map5["hi"] = 20
	// Writing to nil map will throw panic for above map5
	// panic: assignment to entry in nil map

	// goroutine 1 [running]:
	// main.main()
	//         ~/golang-codes/go-syntax-traps/main.go:437 +0x21a8
	// exit status 2
	fmt.Println("To solve writing to empty map we need to initialize with zero value")
	map6 := make(map[string]int)
	map6["Hola"] = 200
	fmt.Println("Value of map6:", map6)

	// Output:
	// Read from nil map5: 0
	// Writing to nil map will throw panic
	// To solve writing to empty map we need to initialize with zero value
	// Value of map6: map[Hola:200]

	fmt.Println("----------------------")
	// 20. Map iteration order is unspecified
	fmt.Println("20. Map iteration order is unspecified")
	fmt.Println("Never depend on for loop iteration of map for ordering")
	fmt.Println("Map iteration order is intentionally unspecified.")
	// If ordering matters:
	// Use an array to store the keys and sort the keys then use the sorted array
	// to loop the keys in the map and return ordered values.
	map7 := map[string]int{"saturday": 6, "tuesday": 2, "sunday": 7, "wednusday": 3, "monday": 1}
	arr4 := make([]string, 0, len(map7))
	for i := range map7 {
		arr4 = append(arr4, i)
	}
	sort.Strings(arr4)
	for _, k := range arr4 {
		fmt.Println("Sorted key:", k, ", its value:", map7[k])
	}

	// Output:
	// Never depend on for loop iteration of map for ordering
	// Map iteration order is intentionally unspecified.
	// Sorted key: monday , its value: 1
	// Sorted key: saturday , its value: 6
	// Sorted key: sunday , its value: 7
	// Sorted key: tuesday , its value: 2
	// Sorted key: wednusday , its value: 3

	fmt.Println("----------------------")
	// 21. Delete doesn't complain if key doesn't exist
	fmt.Println("21. Delete doesn't complain if key doesn't exist")
	delete(map7, "thursday")
	fmt.Println("delete(map7, \"thursday\") is Safe even if \"thursday\" doesn't exist.")
	// Output:
	// 21. Delete doesn't complain if key doesn't exist
	// delete(map7, "thursday") is Safe even if "thursday" doesn't exist.

	fmt.Println("----------------------")
	// 22. Maps aren't safe for unsynchronized concurrent writes
	fmt.Println("22. Maps aren't safe for unsynchronized concurrent writes")
	fmt.Println("Have covered more about this in 1-map-sync-unsafe, 2-map-sync-fix1 and 3-map-sync-fix2")

	fmt.Println("---------INTERFACES-------------")
	// 23. Interfaces are implemented implicitly
	fmt.Println("23. Interfaces are implemented implicitly")
	// Implicit Interfaces
	// In Go, we do not need to explicitly declare that Dog implements Speaker
	// (there is no implements keyword). As long as Dog implements all the methods
	// defined in the Speaker interface (Speak() string), it automatically satisfies the interface.

	// 1. Initialize the Dog struct
	dog := Dog{breed: "Golden Retriever"}
	// 2. Assign it to a variable of type Speaker
	var dogSpeaker Speaker = dog

	// 3. Call the method through the interface
	fmt.Println("Sound:", dogSpeaker.Speak())
	fmt.Println("Breed:", dog.breed)
	// Alternatively, we can pass it directly to a function that accepts the interface:
	makeSound(dog)

	fmt.Println("----------------------")
	// 24. Pointer receiver vs value receiver
	fmt.Println("24. Pointer receiver vs value receiver")
	// For more idea on this we can refer this code "4-interface-pointer-reciver-vs-value-receiver-1"
	// and "5-interface-pointer-reciver-vs-value-receiver-2"
	// When working with interfaces in Go, choosing between a value receiver (func (d Dog)) and
	// a pointer receiver (func (d *Dog)) changes how a struct satisfies an interface.

	// The Golden Rule of Receivers and Interfaces
	// 1) Value Receiver (func (d Dog)): Both values (Dog) and pointers (*Dog) satisfy the interface.

	// 2) Pointer Receiver (func (d *Dog)): Only pointers (*Dog) satisfy the interface. (Strictly
	// speaking, non-pointer values will fail to implement the interface unless they are addressable
	// variables, and even then, it's safer to use pointers).

	fmt.Println("----------------------")
	// 25. any
	fmt.Println("25. any")
	// Previously Go used "interface{}"
	// For example:
	// func printAnything(v interface{}) {
	// }

	// Modern Go has the alias: "any"

	// So:
	// func printAnything(v any) {
	// }

	// means essentially:
	// func printAnything(v interface{}) {
	// }

	// Because an empty interface has zero required methods, virtually
	// every type satisfies it.

	fmt.Println("----------------------")
	// 26. Type assertion
	fmt.Println("26. Type assertion")
	// for more info read "6-type-assertions-interface-1" in same folder.

	fmt.Println("----------------------")
	// 27. Type switch
	fmt.Println("27. Type switch")
	describe(10)
	describe("hello")
	// Output:
	// 27. Type switch
	// integer: 10
	// string: hello
	fmt.Println("----------------------")
	// 28. The infamous nil-interface trap
	fmt.Println("28. The infamous nil-interface trap")
	// we can read more about this at "7-nil-interface-trap" folder
	fmt.Println("----------------------")
	// 29. Interface design: accept interfaces, return concrete types
	fmt.Println("29. Interface design: accept interfaces, return concrete types")
	// 	A useful Go design guideline is generally:
	// func NewService() *Service

	// rather than:
	// func NewService() ServiceInterface

	// and consumers accept the smallest interface they need.
	// For example:
	// type Reader interface {
	//     Read([]byte) (int, error)
	// }

	// This is one reason Go's standard library composes so well.
	// We'll see tiny interfaces like:
	// io.Reader
	// io.Writer
	// error
	// fmt.Stringer

	// 	For example error itself is simply:
	// type error interface {
	//     Error() string
	// }

	// That's why our custom type becomes an error by implementing:
	// Error() string

	// similar to 18. comma ok idiom we can find in interface
	var example_interface interface{} = "hi"
	val_interface, ok := example_interface.(string)
	fmt.Println("value of variable type interface:", val_interface, ", value exists:", ok)
	// Output:
	// value of variable type interface: hi , value exists: true
}

func divide(a, b int) (int, error) {
	fmt.Println("Divide", a, "by", b)
	if b == 0 {
		return 0, errors.New("{\"Divide by zero error\"}")
	}
	return a / b, nil
}

// inplace slice changes will not be reflected when slice passed as value.
func addSlice1(s []int) {
	s = append(s, 100)
}

// returning the inplace slice changes then store them in original slice will work.
func addSlice2(s []int) []int {
	return append(s, 101)
}

// but modify inplace does work even in call by value
func modifySlice1(s []int) {
	s[0] = 145
}

func sliceEqual(a, b []int) (bool, error) {
	if len(a) != len(b) {
		return false, errors.New("Length not equal")
	}
	for i := range len(a) {
		if a[i] != b[i] {
			return false, fmt.Errorf("Slice contents not equal at index %d", i)
		}
	}
	return true, nil
}

type sliceExample struct {
	SliceKey []int `json:"slices"`
}

type Speaker interface {
	Speak() string
}

type Dog struct {
	breed string
}

func (Dog) Speak() string {
	return "Woof Woof!"
}

// function that accept interface
func makeSound(s Speaker) {
	fmt.Println("Inside function that accept interface")
	fmt.Println("Animal sound:", s.Speak())
}

func describe(v any) {
	switch x := v.(type) {
	case int:
		fmt.Println("integer:", x)

	case string:
		fmt.Println("string:", x)

	case bool:
		fmt.Println("boolean:", x)

	default:
		fmt.Println("unknown")
	}
}
