package main

import "fmt"

// ***ALGORITHM***
// - pass first subArray of spiral into output array
// - for every other subArray that is not the last, add the last element of the subArray to the output array
// - for the last subArray, add all elements in reverse order to the output array
// - work back from length - 2 (next to last subArray) through 2nd subArray
//   - for each add the first element to the output
// - increase starting index for spiral AND subArrays by 1 and repeat process

func reverseArray(array []int) []int {
	if len(array) < 2 {
		return array
	}

	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func trimSpiral(spiral [][]int) [][]int {
	var trimmedSpiral [][]int

	for index, subArray := range spiral {
		if index == 0 || index == len(spiral)-1 {
			continue
		} else if len(subArray) > 1 {
			trimmedSpiral = append(trimmedSpiral, subArray[1:len(subArray)-1])
		}
	}
	return trimmedSpiral
}

func spiralMatrix(spiral [][]int) []int {
	var output []int

	if len(spiral) <= 0 {
		return output
	}

	for spiralIndex, subArray := range spiral {
		if len(subArray) < 1 {
			continue
		}

		if spiralIndex == 0 {
			output = append(output, subArray...)
		} else if spiralIndex < len(spiral)-1 {
			fmt.Println("CurrentSpiral:", spiral)
			output = append(output, subArray[len(subArray)-1])
		} else if spiralIndex == len(spiral)-1 {
			output = append(output, reverseArray(subArray)...)
		}
	}

	if len(spiral) > 2 && len(spiral[0]) > 1 {
		for index := len(spiral) - 2; index > 0; index-- {
			output = append(output, spiral[index][0])
		}
	}

	if len(spiral) > 1 {
		fmt.Println("BeforeTrim:", output)
		output = append(output, spiralMatrix(trimSpiral(spiral))...)
	}

	return output
}

func main() {
	fmt.Println(spiralMatrix([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})) // [1,2,3,6,9,8,7,4,5]
	fmt.Println("Reverse:", reverseArray([]int{1, 2, 3}))
	fmt.Println("Trimmed:", trimSpiral([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}))
	fmt.Println(spiralMatrix([][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}})) // [1,2,3,4,8,12,11,10,9,5,6,7]
	fmt.Println(spiralMatrix([][]int{[]int{7}, []int{9}, []int{6}}))
	fmt.Println(trimSpiral([][]int{[]int{7}, []int{9}, []int{6}}))

	fmt.Println(spiralMatrix([][]int{[]int{1, 11}, []int{2, 12}, []int{3, 13}, []int{4, 14}, []int{5, 15}, []int{6, 16}, []int{7, 17}, []int{8, 18}, []int{9, 19}, []int{10, 20}}))

}
