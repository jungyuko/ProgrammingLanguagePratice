package main

import (
"fmt"
"io"
"os"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
	dat, err := os.Open("input.txt")
	check(err)

	var perline int
	var num []int
	for{
		_,err := fmt.Fscanf(dat, "%d\n",&perline);

		if err != nil{
			if err == io.EOF{break}
			fmt.Println(err)
			os.Exit(1)
		}
		num = append(num, perline)
	}
	fmt.Println(num)
	for i:= 0; i< len(num); i++{
		for j:= 0; j< i; j++{
			if(num[i] < num[j]){num[i], num[j] = num[j], num[i]}
		}
	}
	fmt.Println(num)

	f, err := os.Create("output.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	for i := range num{
		fmt.Fprintf(f, "%d\n", num[i])
	}
	f.Close()
}
