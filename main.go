package main
import (
	"fmt"
	"os"
)
func main() {
	val, ok := os.LookupEnv("XXXYYY")
	if !ok {
		panic("NOPE")
	}
	fmt.Println("XXXYYY:", val)
}