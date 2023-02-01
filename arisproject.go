package main

import (
	"fmt"

	s "aris.com/scrapp"
	"github.com/soluchok/freeproxy"
)

// s "aris.com/scrapp"

func main() {

	gen := freeproxy.New()
	fmt.Println(gen.Get())
	// s.WalletData = []string{"ars", "test"}
	s.Start()
	// y.Getsum()
}
