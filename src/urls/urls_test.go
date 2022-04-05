package urls

import (
	"fmt"
	"testing"
)

func Test_shorten(t *testing.T) {
	result, err := Shorten("www.baidu.com")
	if err != nil {
		fmt.Printf("Shorten error: %v\n", err)
	} else {
		fmt.Printf("Shorten succeed: %s", result)
	}
}
