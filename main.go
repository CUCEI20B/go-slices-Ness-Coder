package main

import "fmt"

func main()  {
	var n int
	var suma int
	//var numero int
	//var s  []int
	 
	fmt.Scan(&n)

	s := make([]int,n)//tomamos el tamaño de n

	for i:=0; i < len(s); i++ {
		fmt.Scan(&s[i])
		suma = suma + s[i]
	}

	fmt.Println(suma)
}