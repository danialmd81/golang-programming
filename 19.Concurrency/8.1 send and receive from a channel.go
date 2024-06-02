package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Start Main Function ........")

	ch := make(chan int)

	go send(ch)

	go receive(ch)

	time.Sleep(time.Microsecond * 100)

	fmt.Println("End of Main Function ........")

}

func send(ch chan int) {

	ch <- 50 //اینجا به چنل 50 را فزستادیم
}

func receive(ch chan int) {

	fmt.Println(200 + <-ch) //اینجا چیزی که به چنا فرستاده بودیمو با 200 جمع کردیم
}
