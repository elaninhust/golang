package main

import "fmt"

func main(){
	s1 := make([]int, 5)
	fmt.Printf("the length of s1 is : %d\n", len(s1))
	fmt.Printf("the cap of s1 is : %d\n", cap(s1))
	fmt.Printf("the value of s1 is : %d\n", s1)

	s2 := make([]int, 5, 8)
	fmt.Printf("the length of s2 is : %d\n", len(s2))
	fmt.Printf("the cap of s2 is : %d\n", cap(s2))
	fmt.Printf("the value of s2 is : %d\n", s2)

	s3 := []int{1,2,3,4,5,6,7}
	s4 := s3[2:4]

	fmt.Printf("length of s4: %d\n", len(s4))
	fmt.Printf("cap of s4: %d\n", cap(s4))
	fmt.Printf("value of s4: %d\n", s4)
}