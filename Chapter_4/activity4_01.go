package main

import "fmt"

func main() {
	l1, l2, l3 := linked()
	fmt.Println("Linked: ", l1, l2, l3)

	nl1, nl2 := noLink()
	fmt.Println("No Link: ", nl1, nl2)

	cl1, cl2 := capLinked()
	fmt.Println("Cap Link: ", cl1, cl2)

	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link: ", cnl1, cnl2)

	cp1, cp2, cp := copyNoLink()
	fmt.Print("Copy No Link: ", cp1, cp2)
	fmt.Printf(" (Number of elements copied : %v)\n", cp)

	a1, a2 := appendNoLink()
	fmt.Println("Append No Link: ", a1, a2)
}

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	s1[2] = 99
	return s1[2], s2[2], s3[2]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	fmt.Println("New S1: ", s1)
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3], s2[3]
}
