package obss

import (
	"fmt"
	"testing"
)

func TestPutFile(*testing.T) {
	Init()

	putFile("1", "../resources/bear.mp4")
}

func TestPutAndGet(*testing.T) {
	Init()
	putFile("1", "../resources/bear.mp4")
	fmt.Println(genGetURL("1"))
}
