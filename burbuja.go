package main

//import "fmt"

func Burbuja(s []int64) {

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++{
			if s[i] > s[j] {
				aux := s[i]
				s[i] = s[j]
				s[j] = aux
		}
		
		}
	}
}

