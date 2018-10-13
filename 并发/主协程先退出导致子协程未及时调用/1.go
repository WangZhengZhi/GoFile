package main

import "fmt"

/*主协程先退出，导致子协程没来得及调用 */
func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程i=", i)
		}
	}()
}
