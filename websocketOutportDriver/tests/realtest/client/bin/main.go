package main

import (
	"flag"
	"log"

	"github.com/Reshusk23/sr-me-core/marshal"
	"github.com/Reshusk23/sr-me-core/websocketOutportDriver/tests/realtest/client"
)

var (
	addr = flag.String("name", "client 0", "-")
	port = flag.Int("port", 21112, "-")
)

func main() {
	tc, err := client.NewTempClient(*addr, &marshal.JsonMarshalizer{})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer tc.Stop()

	tc.Run(*port)
}
