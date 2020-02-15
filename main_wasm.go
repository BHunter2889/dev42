// +build wasm

package main

import (
	"github.com/vugu/vugu"
	"log"
	"os"
)

func main() {

	println("Entering main()")
	defer println("Exiting main()")

	// Init any global state here if needed...

	rootInst, err := vugu.New(&Root{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	env := vugu.NewJSEnv("#root_mount_parent", rootInst, vugu.RegisteredComponentTypes())
	env.DebugWriter = os.Stdout

	for ok := true; ok; ok = env.EventWait() {
		err = env.Render()
		if err != nil {
			panic(err)
		}
	}

}
