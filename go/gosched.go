package main
import(
	"runtime"
	"fmt"
)
func showNumber(num int) {
	fmt.Println(num)
}
//что выведется в коснсоль?
func main() {
	runtime.GOMAXPROCS(1)
	iterations := 10
	for i := 0; i<=iterations; i++ {
		go showNumber(i)
	}
	runtime.Gosched()
	fmt.Println("Goodbye!")
}
