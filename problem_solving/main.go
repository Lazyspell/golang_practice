package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// anagram("hello world", "hello world")

	// if !CheckCategory("chair") {
	// 	fmt.Println("chair not in list")
	// }

	// var myName models.Name

	// myName.First = "Jeremy"
	// fmt.Println(myName.Age)

	// fmt.Println(CheckBool(myName.First))

	// rect := rectangle{length: 5, width: 2}

	// fmt.Println(rect.area())
	// fmt.Println(rect.perimeter())

	// testList := []int{1, 2, 3, 4, 5}

	// removeItem(testList, 1)

	// testString := []string{"nail", "screw", "hammer"}
	// testString2 := []string{"wheel", "saw"}

	// testParts := parts{random: testString}

	// // testParts.printParts()
	// testParts.addParts(testString2)
	// // testParts.printParts()
	// testParts.removeParts(0)
	// testParts.removeParts(0)
	// testParts.printParts()
	// ["a", "b", "c", 'd', 'e', 'f','p', 'o', 'o', 'p']

	list := []string{"a", "b", "c", "d", "e", "f", "p", "o", "o", "p"}

	test := ransomNote(list, "pooop")
	fmt.Println(test)

	fibValue := fib(8)
	fmt.Println(fibValue)
	twoValue := twoSums([]int{1, 3, 2, 2, 2}, 4)
	fmt.Println(twoValue)

}

func anagram(word1 string, word2 string) bool {
	word1 = strings.ReplaceAll(word1, " ", "")
	// word2 = strings.ReplaceAll(word2, " ", "")

	count := make(map[string]int)

	arrString := strings.Split(word1, "")
	// arrString2 := strings.Split(word2, "")

	for i := 0; i < len(arrString); i++ {
		count[arrString[i]]++
	}

	for i, dog := range count {
		fmt.Println(i, dog)
	}

	return true
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	left := mergeSort(items[:len(items)/2])
	right := mergeSort(items[len(items)/2:])
	return merge(left, right)
}
func merge(left []int, right []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		final = append(final, left[i])
	}
	for ; j < len(right); j++ {
		final = append(final, right[j])
	}
	return final
}

func deleteSlice(list []int, index int) []int {

	list = append(list[:index], list[index+1:]...)

	return list

}

func CheckCategory(item string) bool {

	categories := []string{"seating", "sleeping or laying", "entertainment", "tables", "storage"}

	for _, value := range categories {
		if value == item {
			return true
		}
	}

	return false
}

func CheckBool(item string) bool {
	return item != ""
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

type shape interface {
	area() float64
}

type calculator struct {
}

func (a calculator) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

type rectangle struct {
	length float64
	width  float64
}

func (m *rectangle) area() float64 {

	return m.length * m.width
}

func (m *rectangle) perimeter() float64 {
	return (2 * m.length) + (2 * m.width)
}

func removeItem(list []int, index int) {
	left := list[index+1:]
	right := list[:index]

	list = append(right, left...)
	fmt.Println(list)
}

type parts struct {
	random []string
}

func (m *parts) printParts() {
	for _, part := range m.random {
		fmt.Println(part)
	}
}

func (m *parts) addParts(list []string) []string {
	m.random = append(m.random, list...)
	return m.random
}

func (m *parts) removeParts(index int) []string {
	m.random = append(m.random[index+1:], m.random[:index]...)

	return m.random
}

func ransomNote(list []string, word string) bool {
	count := make(map[string]int)

	for _, item := range list {
		count[item]++
	}

	for _, letter := range word {
		count[string(letter)]--
		if count[string(letter)] < 0 {
			return false
		}
	}
	return true
}

func fib(num int) int {

	if num == 0 || num == 1 {
		return num
	} else {
		return fib(num-1) + fib(num-2)
	}
}

func twoSums(list []int, target int) [][]int {
	var finalList [][]int
	seen := make(map[int]bool)

	for _, value := range list {
		difference := target - value
		if !seen[difference] {
			seen[value] = true
		} else {
			combination := []int{difference, value}
			finalList = append(finalList, combination)
			delete(seen, difference)
		}

	}
	return finalList
}
