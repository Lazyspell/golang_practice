package main

import "fmt"

// import (
// 	"fmt"
// 	"math"
// 	"sort"
// 	"strings"
// )

func main() {
	v := 15 % 4
	switch v {
	case 3:
		fmt.Println(100)
		fallthrough

	case 2:
		fmt.Println(42)

	case 1:
		fmt.Println(1)
		fallthrough

	default:
		fmt.Println("default")
	}
}

// func main() {
// 	// anagram("hello world", "hello world")

// 	// if !CheckCategory("chair") {
// 	// 	fmt.Println("chair not in list")
// 	// }

// 	// var myName models.Name

// 	// myName.First = "Jeremy"
// 	// fmt.Println(myName.Age)

// 	// fmt.Println(CheckBool(myName.First))

// 	// rect := rectangle{length: 5, width: 2}

// 	// fmt.Println(rect.area())
// 	// fmt.Println(rect.perimeter())

// 	// testList := []int{1, 2, 3, 4, 5}

// 	// removeItem(testList, 1)

// 	// testString := []string{"nail", "screw", "hammer"}
// 	// testString2 := []string{"wheel", "saw"}

// 	// testParts := parts{random: testString}

// 	// // testParts.printParts()
// 	// testParts.addParts(testString2)
// 	// // testParts.printParts()
// 	// testParts.removeParts(0)
// 	// testParts.removeParts(0)
// 	// testParts.printParts()
// 	// ["a", "b", "c", 'd', 'e', 'f','p', 'o', 'o', 'p']

// 	// list := []string{"a", "b", "c", "d", "e", "f", "p", "o", "o", "p"}

// 	// test := ransomNote(list, "pooop")
// 	// fmt.Println(test)

// 	// fibValue := fib(8)
// 	// fmt.Println(fibValue)
// 	// twoValue := twoSum([]int{3, 2, 4}, 6)
// 	// fmt.Println(twoValue)

// 	// indicesValue := indices([]int{1, 1, 3, 3, 5, 7, 8, 8, 9, 9, 9, 15}, 5)
// 	// fmt.Println(indicesValue)

// 	// badVersionValue := firstBadVersion(5)

// 	// fmt.Println("___________________________________")
// 	// fmt.Println(badVersionValue)

// 	// searchInsert := searchInsert2([]int{1, 3, 5, 6}, 0)
// 	// fmt.Println(searchInsert)

// 	// pivotList := []int{1, 7, 3, 6, 5, 6}
// 	// left := pivotList

// 	// rotateList := []int{1, 2, 0, 3, 4, 0, 5, 6, 7}
// 	// zeroList := []int{1, 0, 0}

// 	// rotateList = rotate(rotateList, 3)
// 	// fmt.Println(rotateList)

// 	// fmt.Println(reverseString("tes"))

// 	// fmt.Println(10 % 10)

// 	// deleteFromValue := deleteFromSlice(zeroList)

// 	// fmt.Println(deleteFromValue)

// 	// moveZeroes(rotateList)

// 	// reverseValue := reverseString2("Hello world")
// 	// fmt.Println(reverseValue)

// 	// testMap := make(map[string]int)

// 	// testMap["test1"] = 3
// 	// testMap["test2"] = 4

// 	// stringTest := "tes"

// 	// stringList := []string{"h", "e", "l", "l", "o"}
// 	// reverseString3(stringList)
// 	// fmt.Println(len(rotateList))

// 	// valueWord := reverseWords("Let's take LeetCode contest")
// 	// fmt.Println(valueWord)

// 	// FirstLastIndices.first_last([1,1,3, 3, 5, 7, 8, 8, 9, 9, 9, 15], 9)
// 	// test := []int{1, 1, 3, 3, 5, 7, 8, 8, 9, 9, 9, 15}
// 	// var final []int
// 	// for index, value := range test {
// 	// 	if value == 9 {
// 	// 		final = append(final, index)
// 	// 	}
// 	// }

// 	// fmt.Println(final[0], final[len(final)-1])
// 	// // lengthOfLongestSubstring("testt")

// 	answer := anagram("dog", "good")
// 	fmt.Println(answer)

// 	test := []int{-2, 45, 0, 11, -9, 88, -97, -202, 747}

// 	fmt.Println(mergeSort(test))

// 	// arr1 :=

// 	another := mergeTwo([]int{1, 2, 3, 4, 5, 5}, []int{1, 1, 6, 7, 7, 7, 8, 9, 1})
// 	fmt.Println(another)

// }

// func anagram(word1 string, word2 string) bool {

// 	count := make(map[string]int)

// 	for _, value := range word1 {
// 		count[string(value)]++
// 	}

// 	for _, value := range word2 {
// 		count[string(value)]--
// 		if count[string(value)] < 0 {
// 			return false
// 		}
// 	}
// 	return len(word1) == len(word2)
// }

// func mergeSort(items []int) []int {
// 	if len(items) <= 1 {
// 		return items
// 	}
// 	middle := len(items) / 2
// 	left := mergeSort(items[:middle])
// 	right := mergeSort(items[middle:])

// 	return merge(left, right)
// }
// func merge(left []int, right []int) []int {
// 	var final []int
// 	i := 0
// 	j := 0

// 	for i < len(left) && j < len(right) {
// 		if left[i] <= right[j] {
// 			final = append(final, left[i])
// 			i++
// 		} else {
// 			final = append(final, right[j])
// 			j++
// 		}
// 	}

// 	for i < len(left) {
// 		final = append(final, left[i])
// 		i++
// 	}
// 	for j < len(right) {
// 		final = append(final, right[j])
// 		j++
// 	}
// 	return final

// }

// func mergeTwo(arr1 []int, arr2 []int) []int {

// 	count := make(map[int]bool)
// 	var final []int

// 	for _, value := range arr1 {
// 		count[value] = true
// 	}

// 	for _, value := range arr2 {
// 		count[value] = true
// 	}

// 	for key, _ := range count {
// 		final = append(final, key)
// 	}

// 	sort.Ints(final)
// 	return final

// }

// func deleteSlice(list []int, index int) []int {

// 	list = append(list[:index], list[index+1:]...)

// 	return list

// }

// func CheckCategory(item string) bool {

// 	categories := []string{"seating", "sleeping or laying", "entertainment", "tables", "storage"}

// 	for _, value := range categories {
// 		if value == item {
// 			return true
// 		}
// 	}

// 	return false
// }

// func CheckBool(item string) bool {
// 	return item != ""
// }

// type circle struct {
// 	radius float64
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// type square struct {
// 	length float64
// }

// func (s square) area() float64 {
// 	return s.length * s.length
// }

// type triangle struct {
// 	height float64
// 	base   float64
// }

// func (t triangle) area() float64 {
// 	return t.base * t.height / 2
// }

// type shape interface {
// 	area() float64
// }

// type calculator struct {
// }

// func (a calculator) areaSum(shapes ...shape) float64 {
// 	var sum float64
// 	for _, shape := range shapes {
// 		sum += shape.area()
// 	}
// 	return sum
// }

// type rectangle struct {
// 	length float64
// 	width  float64
// }

// func (m *rectangle) area() float64 {

// 	return m.length * m.width
// }

// func (m *rectangle) perimeter() float64 {
// 	return (2 * m.length) + (2 * m.width)
// }

// func removeItem(list []int, index int) {
// 	left := list[index+1:]
// 	right := list[:index]

// 	list = append(right, left...)
// 	fmt.Println(list)
// }

// type parts struct {
// 	random []string
// }

// func (m *parts) printParts() {
// 	for _, part := range m.random {
// 		fmt.Println(part)
// 	}
// }

// func (m *parts) addParts(list []string) []string {
// 	m.random = append(m.random, list...)
// 	return m.random
// }

// func (m *parts) removeParts(index int) []string {
// 	m.random = append(m.random[index+1:], m.random[:index]...)

// 	return m.random
// }

// func ransomNote(list []string, word string) bool {
// 	count := make(map[string]int)

// 	for _, item := range list {
// 		count[item]++
// 	}

// 	for _, letter := range word {
// 		count[string(letter)]--
// 		if count[string(letter)] < 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func fib(num int) int {

// 	if num == 0 || num == 1 {
// 		return num
// 	} else {
// 		return fib(num-1) + fib(num-2)
// 	}
// }

// // func twoSums(list []int, target int) [][]int {
// // 	var finalList [][]int
// // 	seen := make(map[int]bool)

// // 	for _, value := range list {
// // 		difference := target - value
// // 		if !seen[difference] {
// // 			seen[value] = true
// // 		} else {
// // 			combination := []int{difference, value}
// // 			finalList = append(finalList, combination)
// // 			delete(seen, difference)
// // 		}

// // 	}
// // 	return finalList
// // }

// // [3,2,4]
// // 6
// func twoSum(nums []int, target int) []int {
// 	seen := make(map[int]bool)
// 	indexes := make(map[int]int)
// 	var final []int
// 	for index, value := range nums {
// 		diff := target - value
// 		fmt.Println(seen[diff])
// 		if !seen[diff] {

// 			seen[value] = true
// 			indexes[value] = index
// 		} else {
// 			final = append(final, indexes[diff], index)
// 		}
// 	}
// 	return final
// }

// func indices(list []int, num int) []int {
// 	var finalList []int
// 	var emptyList []int
// 	var seen []int
// 	for index, value := range list {
// 		if value == num {
// 			seen = append(seen, index)
// 		}
// 	}

// 	if len(seen) <= 1 {
// 		return emptyList
// 	} else {

// 		finalList = append(finalList, seen[0], seen[len(seen)-1])
// 	}

// 	return finalList

// }

// func isBadVersion(num int) bool {
// 	if num != 1 {
// 		return false
// 	}
// 	return true
// }

// func firstBadVersion(n int) int {

// 	return firstBadRec(1, n)
// }

// func firstBadRec(left, right int) int {

// 	if left == right {
// 		return left
// 	}

// 	if middle := (left + right) / 2; isBadVersion(middle) {
// 		return firstBadRec(left, middle)
// 	} else {
// 		return firstBadRec(middle+1, right)
// 	}

// }

// func searchInsert(nums []int, target int) int {
// 	return sort.SearchInts(nums, target)
// }

// func searchInsert2(nums []int, target int) int {
// 	for index, value := range nums {
// 		if value > target {
// 			return index
// 		}
// 		if value == target {
// 			return index
// 		}
// 	}
// 	return len(nums)
// }

// func pivotIndex(list []int) int {
// 	var left int
// 	for _, value := range list {
// 		left += value
// 	}

// 	var right int
// 	for index, value := range list {
// 		if left-value-right == right {
// 			return index
// 		}
// 		right += value
// 	}
// 	return -1
// }

// // func merge2(left, right []int) []int {

// // }

// // func mergeSort2(list []int) []int {
// // 	if len(list) <= 1 {
// // 		return list
// // 	}
// // 	middle := len(list) / 2
// // 	left := mergeSort2(list[:middle])
// // 	right := mergeSort2(list[middle:])

// // 	final := merge2(left, right)

// // 	return final

// // }

// func rotate(nums []int, k int) []int {
// 	idx := len(nums) - (k % len(nums))

// 	nums = append(nums[idx:], nums[:idx]...)
// 	return nums
// }

// func reverseString(str string) (result string) {
// 	// iterate over str and prepend to result
// 	for _, v := range str {
// 		result = string(v) + result
// 		fmt.Println(result)
// 	}
// 	return
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	cur := new(ListNode)
// 	ret := cur
// 	sum := 0

// 	for l1 != nil || l2 != nil {

// 		if l1 != nil {

// 			sum = sum + l1.Val
// 			l1 = l1.Next
// 		}

// 		if l2 != nil {
// 			sum = sum + l2.Val
// 			l2 = l2.Next
// 		}
// 		fmt.Print(sum)

// 		cur.Next = &ListNode{
// 			sum % 10,
// 			nil,
// 		}

// 		cur = cur.Next
// 		sum = sum / 10

// 	}

// 	if sum > 0 {
// 		cur.Next = &ListNode{sum, nil}
// 	}

// 	return ret.Next

// }

// func deleteFromSlice(list []int) []int {

// 	var zeroCount []int
// 	var countingZero int

// 	for _, value := range list {
// 		if value == 0 {
// 			countingZero++
// 		}
// 	}

// 	if countingZero == len(list) {
// 		return list
// 	}

// 	for index, value := range list {
// 		if value == 0 {
// 			zeroCount = append(zeroCount, value)
// 			list = append(list[:index], list[index+1:]...)
// 		}
// 	}

// 	fmt.Println(zeroCount)

// 	list = append(list, zeroCount...)
// 	return list
// }

// func moveZeroes(nums []int) {
// 	writepointer := 0
// 	for index, value := range nums {
// 		fmt.Println(nums, index, writepointer)
// 		if value != 0 {
// 			nums[index], nums[writepointer] = nums[writepointer], nums[index]
// 			writepointer++
// 		}
// 	}
// }

// func reverseString2(word string) string {

// 	var final string
// 	for _, value := range word {
// 		final = string(value) + final
// 	}

// 	return final
// }

// func isSubsequence(s, t string) bool {
// 	for _, c := range s {
// 		if i := strings.IndexRune(t, c); i == -1 {
// 			return false
// 		} else {
// 			t = t[i+1:]
// 		}
// 	}
// 	return true
// }

// func reverseWords(s string) string {
// 	stringSlice := strings.Split(s, " ")
// 	for index, value := range stringSlice {
// 		stringSlice[index] = reverseString2(value)
// 	}

// 	final := strings.Join(stringSlice[:], " ")

// 	return final
// }

// func fib2(num int) int {

// 	if num == 0 || num == 1 {
// 		return num
// 	}

// 	return fib2(num-1) + fib2(num-2)
// }

// func lengthOfLongestSubstring(s string) int {
// 	sLen := len(s)
// 	if sLen < 1 {
// 		return sLen
// 	}

// 	max := 0
// 	m1 := [256]int{}
// 	left := 0
// 	for right, val := range s {
// 		if m1[val] > 0 && left < m1[val] {
// 			left = m1[val]
// 		}
// 		m1[val] = right + 1
// 		fmt.Println(right, left, m1[val])
// 		tempMax := right - left + 1
// 		if max < tempMax {
// 			max = tempMax
// 		}
// 	}
// 	return max
// }

// func Solution(S string, K int) string {

// 	dayOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
// 	var weekIndex int

// 	for index, value := range dayOfWeek {
// 		if value == S {
// 			weekIndex = index
// 		}
// 	}

// 	return dayOfWeek[(weekIndex+K)%7]

// }
