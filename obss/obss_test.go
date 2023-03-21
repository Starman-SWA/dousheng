package obss

import (
	"fmt"
	"testing"
)

func TestPutFile(*testing.T) {

	err := PutFile("1", "../resources/bear.mp4")
	if err != nil {
		return
	}
}

func TestPutAndGet(*testing.T) {
	err := PutFile("1", "../resources/bear.mp4")
	if err != nil {
		return
	}
	fmt.Println(GenGetURL("1"))
}
