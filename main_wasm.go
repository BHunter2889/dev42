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

	initState()
	state.initGrid(20)

	/**
	Tried the below but apparently it is non-reactive. Instead using state directly in root.vugu
	*/
	//var rd RootData
	//rootDataMeta := reflect.TypeOf(rd)
	//rootDataLen := rootDataMeta.NumField()
	//
	//props := make(vugu.Props, rootDataLen)
	////for i := 0; i < rootDataLen; i++ {
	////	props[rootDataMeta.Field(i)] = state
	////}
	//props["NumSides"] = state.NumSides
	//props["NumRows"] = state.NumRows
	//props["TotalRollCount"] = state.TotalRollCount
	//props["MinNumberOfRolls"] = state.DieBalanceComputationValues.DieConstants.MinNumberOfRolls
	//props["IsMinNumRollsMet"] = state.DieBalanceComputationValues.DieConstants.MinNumberOfRolls <= state.TotalRollCount
	//props["BalanceThreshold"] = state.DieBalanceComputationValues.BalanceThreshold
	//props["ExpectedRollsPerSide"] = state.DieBalanceComputationValues.ExpectedRollsPerSide
	//props["SumSquaredError"] = state.DieBalanceComputationValues.SumSquaredError
	//props["IsBalanced"] = state.DieBalanceComputationValues.IsBalanced

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
