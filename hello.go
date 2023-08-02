package main

func main() {
	//
	//
	//
	//
	//
	//

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//

	// skillLevel := 5

	// switch {

	// case skillLevel >= 80:
	// 	fmt.Println("high")

	// case skillLevel >= 50:
	// 	fmt.Println("mid")

	// case skillLevel >= 30:
	// 	fmt.Println("low")

	// case skillLevel >= 20, skillLevel >= 0:
	// 	fmt.Println("low")

	// default:
	// 	fmt.Println("don't know don't care")

	// }
	// if 5 > 3 && 9 < 1 {
	// 	fmt.Println("win")
	// } else if 88 > 69 {
	// 	fmt.Println("what the fridge not even 69")
	// } else {
	// 	fmt.Println("you lost deez nuts")
	// }

	// fmt.Println(5 != 8 && 8 != 9)
	// fmt.Println(7 >= 6 || 3 < 1)

	// Go Access, Change, Append and Copy Slices
	//* Append Elements To a Slice  append()

	// mySlice4 := []int{1, 2, 3, 4, 5}
	// fmt.Printf("mySlice4 %v | %d | %d", mySlice4, len(mySlice4), cap(mySlice4))
	// /* mySlice4 [1 2 3 4 5] | 5 | 5 */
	// mySlice4 = append(mySlice4, 6, 7, 8)
	// fmt.Printf("mySlice4 %v | %d | %d", mySlice4, len(mySlice4), cap(mySlice4))
	// /* mySlice4 [1 2 3 4 5 6 7 8] | 8 | 10 */
	//
	//
	//
	//* Append One Slice To Another Slice
	//* slice3 = append(slice1, slice2...)
	// mySlice5 := []int{1, 2, 3}
	// mySlice6 := []int{4, 5, 6}
	// myMixedSlice := append(mySlice5, mySlice6...)

	// fmt.Printf("myMixedSlice %v | %d | %d", myMixedSlice, len(myMixedSlice), cap(myMixedSlice))
	//? myMixedSlice [1 2 3 4 5 6] | 6 | 6
	//
	//
	//* Change The Length of a Slice
	//
	//
	//
	//
	//
	//
	// Go Slices

	// firstSlice := []int32{1000, 50, 700}

	// fmt.Printf("Slice  %T | %v | %#v", firstSlice, firstSlice, firstSlice)
	// Slice  []int32 | [1000 50 700] | []int32{1000, 50, 700}

	// // len() (the number of elements in the slice)
	// // cap() (the number of elements the slice can grow or shrink to)
	// myslice1 := []int{}
	// fmt.Println(len(myslice1))
	// fmt.Println(cap(myslice1))
	// fmt.Println(myslice1)
	// // 0
	// // 0
	// // []
	// myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	// fmt.Println(len(myslice2))
	// fmt.Println(cap(myslice2))
	// fmt.Println(myslice2)
	// // 4
	// // 4
	// // [Go Slices Are Powerful]

	// Create a Slice From an Array
	// arrToSlice := [7]int{1, 2, 3, 4, 5, 6, 7}
	// // [start at index : stop at index ]
	// sliceFromArr := arrToSlice[3:6]

	// fmt.Println(sliceFromArr)                   // [4 5 6]
	// fmt.Println("Length:", len(sliceFromArr))   // 3
	// fmt.Println("capacity:", cap(sliceFromArr)) // 4
	// // WHY: arr cap is 7 & len 7
	// // but we slice from index 3 to 6 but slice take cap from the starting index to the arr end (7)

	// //? Create a Slice With The make() Function
	// slice_name := make([]type, length, capacity)
	// mySlice3 := make([]string, 3, 5)
	// fmt.Printf("mySlice3  %v | %d | %d", mySlice3, len(mySlice3), cap(mySlice3))
	//* mySlice3  [  ] | 3 | 5
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	// Declare an Array

	// var arr1, arr2 = [3]byte{5, 99, 240}, [2]int8{-20, 100}
	// fmt.Printf("arr1  %T | %v \n", arr1, arr1)
	// fmt.Printf("arr2  %T | %v \n", arr2, arr2)

	// var arr3, arr4 = [...]byte{5, 99, 240}, [...]int8{-20, 100}
	// fmt.Printf("arr1  %T | %v \n", arr3, arr3)
	// fmt.Printf("arr2  %T | %v \n", arr4, arr4)

	// arr7 := [...]string{"Volvo", "BMW", "Ford", "Mazda"}
	// fmt.Println(arr7[0]) // Volvo
	// fmt.Println(arr7[2]) // Ford

	// arr7[0] = "Tasla"
	// fmt.Println(arr7[0]) // Tasla
	// fmt.Println(arr7)    // [Tasla BMW Ford Mazda]

	// const arr8 = [2]int{50,69} // so you can't have constant array ooookey!

	// arr8 := [5]int{}               //not initialized
	// arr9 := [5]int{1, 2}           //partially initialized
	// arr10 := [5]int{1, 2, 3, 4, 5} //fully initialized

	// fmt.Println(arr8)
	// fmt.Println(arr9)
	// fmt.Println(arr10)
	// [0 0 0 0 0]
	// [1 2 0 0 0]
	// [1 2 3 4 5]

	// Initialize Only Specific Elements

	// (index) : (default value)
	// arr11 := [4]byte{0: 240, 3: 69} // [240 0 0 69]
	// fmt.Println(arr11)

	// // Find the Length of an Array  len() function
	// /* len(Array) */
	// fmt.Println(len(arr11)) // 4
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	// User Info
	// var (
	// 	// name            = "hello"
	// 	// password int    = 6546840002648
	// 	email string = "something@gmail.com"
	// 	bin   bool   = false
	// )

	// User Info 2
	// const (
	// 	NAME     string = "hello"
	// 	PASSWORD int    = 5
	// 	EMAIL    string = "something@gmail.com"
	// )

	// fmt.Println(name, password, email)
	// fmt.Println(NAME, PASSWORD, EMAIL)

	// fmt.Print(NAME, "\n")
	// fmt.Println(NAME)
	// fmt.Print(email, "\n")
	// fmt.Printf("hi my value is: %v and my type is %T + %% \n %#v \n %b \n %#v", PASSWORD, PASSWORD, NAME, PASSWORD, bin)

	// const MY_NAME string = "AlexEG"
	// fmt.Printf("\n This Const type is: %T and it's value is: %#v", MY_NAME, MY_NAME)

	// x := 500
	// y := 0101011
	// c := 15

	// fmt.Printf("%b", x)

	// fmt.Print("\n")
	// fmt.Printf("%d", y)
	// fmt.Printf("%+d", y)

	// fmt.Print("\n")
	// fmt.Printf("%X", c)

	//////////////////////////
	// String Formatting Verbs
	// str1 := "Hello"
	// fmt.Printf("%s\n", str1)
	// fmt.Printf("%q\n", str1)
	// // %q == $#v && %v == %s

	// fmt.Printf("%x\n", str1)
	// fmt.Printf("% x\n", str1)

	// Float Formatting Verbs
	// const PI float64 = 3.14159265359

	// fmt.Printf("%e\n", PI)
	// fmt.Printf("%f\n", PI)

	// fmt.Printf("%.2f\n", PI)  //precision 2
	// fmt.Printf("%.5f\n", PI)  //precision 5
	// fmt.Printf("%.11f\n", PI) //precision 11

	// fmt.Printf("%6.2f\n", PI)
	// fmt.Printf("%g\n", PI)

	// INT 32/64

	// var num1 int8 = 12 // -128 to 127
	// fmt.Println(num1)

	// var num2 int16 = 12_002 // -32_768 to 32_767
	// fmt.Println(num2)

	// var num3 int32 = 2023_08_02 // -2147483648 to 2_147_483_647
	// fmt.Println(num3)

	// var num4 int64 = 1_000_000_000_000_000_000 // 	-9223372036854775808 to 9223372036854775807
	// fmt.Println(num4)

	// Unsigned Integers
	// var (
	// 	// uint
	// 	// 0 to 4294967295 in 32 bit systems
	// 	// 0 to 18446744073709551615 in 64 bit systems

	// 	num5 uint8  = 240                         //	0 to 255
	// 	num6 uint16 = 50_000                      //	0 to 65_535
	// 	num7 uint32 = 3_000_000_000               //	0 to 4_294_967_295
	// 	num8 uint64 = 1_000_000_000_000_000_000_0 //	0 to 18446744073709551615

	// )

	// fmt.Println(num5)
	// fmt.Println(num6)
	// fmt.Println(num7)
	// fmt.Println(num8)

}
