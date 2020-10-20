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

func main() {
	/*sa := []int64{9, 7, 8, 10, 5, 4, 3, 2, 1, 0, -1, -2, 5}
	fmt.Println(len(sa))
	Burbuja(sa)
	fmt.Println(sa)*/

}
