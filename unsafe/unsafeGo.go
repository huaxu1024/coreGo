package unsafe

import (
	"fmt"
	"unsafe"
)

func unsafeAs() {
	fmt.Println(unsafe.Sizeof(float64(0)))
}
