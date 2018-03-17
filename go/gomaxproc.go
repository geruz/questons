//Что выведется в консоль?

package main
import(
	"runtime"
	"fmt"
)
func showNumber(num int) {
  //time.Sleep(1)
	fmt.Println(num)
}
func main() {
	runtime.GOMAXPROCS(0)
	iterations := 10
	for i := 0; i<=iterations; i++ {
		go showNumber(i)
	}
	fmt.Println("Goodbye!")
}
