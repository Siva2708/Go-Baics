package main

import "fmt"

func main() {

	var arr = []int{}
loop:
	for {
		// k := 0
		// k++
		fmt.Println("\n1. Show Array")
		fmt.Println("2. Show size of array")
		fmt.Println("3. Append element")
		fmt.Println("4. Append multiple elements")
		fmt.Println("5. update Element by index")
		fmt.Println("6. Remove Element by index")
		fmt.Println("7. Remove element")
		fmt.Println("8. check if element exist")
		fmt.Println("9. count the given element")
		fmt.Println("10. check capacity of slice")
		fmt.Println("100. Quit")
		var a int
		fmt.Println("Enter the operation to be executed:")
		fmt.Scan(&a)

		switch a {

		case 1:
			showArray(arr)
		case 2:
			arrayLength(arr)
		case 3:
			fmt.Println("Enter the value to be added:")
			var val int
			fmt.Scan(&val)
			arr = append(arr, val)
			for i := range arr {
				fmt.Println(arr[i])
			}

		case 4:
			var end int
			fmt.Println("Enter how many numbers to be appended:")
			fmt.Scan(&end)
			fmt.Println("Enter the numbers:")
			var v int
			for i := 0; i < end; i++ {
				fmt.Scan(&v)
				arr = append(arr, v)
			}

		case 5:
			var index int
			fmt.Println("Enter the index:")
			fmt.Scan(&index)
			var num int
			fmt.Println("Enter the number:")
			fmt.Scan(&num)
			for r := 0; r < len(arr); r++ {
				if index == r {
					arr[index] = num
				}
			}
		case 6:
			var index int
			fmt.Println("Enter the index:")
			fmt.Scan(&index)
			for r := index; r < len(arr)-1; r++ {
				arr[r] = arr[r+1]
			}
		case 7:
			var num int
			fmt.Println("Enter the element to be deleted:")
			fmt.Scan(&num)
			for r := 0; r < len(arr); r++ {
				if arr[r] == num {
					for k := r; k < len(arr)-1; k++ {
						arr[k] = arr[k+1]
					}
				}
			}
			arr[len(arr)-1] = 0
		case 8:
			var num int
			fmt.Println("Enter the element to be checked:")
			fmt.Scan(&num)
			for r := 0; r < len(arr); r++ {
				if arr[r] == num {
					fmt.Println("Element is present")
				}

			}

		case 100:
			break loop
		}
	}
}

func showArray(arr []int) {
	fmt.Println("Array:", arr)
}

func arrayLength(arr []int) {
	var c int
	for i, _ := range arr {
		c = i

	}
	c = c + 1
	fmt.Println("Size of array is:", c)
}
