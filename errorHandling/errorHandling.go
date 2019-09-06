package errorHandling

import (
	"fmt"
	"os"
)

func ErrorHandingFunc() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
