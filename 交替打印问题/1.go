package main

import "fmt"

/*交替打印问题*/
func main() {
	chan_n := make(chan bool)
	chan_c := make(chan bool, 1)
	done := make(chan interface{})
	go func() {
		for i := 0; i < 11; i = i + 2 {
			<-chan_c
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_n <- true
		}
	}()
	go func() {
		char_seq := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
		for i := 0; i < 10; i = i + 2 {
			<-chan_n
			fmt.Print(char_seq[i])
			fmt.Print(char_seq[i+1])
			chan_c <- true

		}
		done <- true

	}()
	chan_c <- true
	<-done
}
