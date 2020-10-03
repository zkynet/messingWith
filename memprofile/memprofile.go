package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	runtime.SetCPUProfileRate(5000)
	f2, err := os.Create("mem.pprof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f2); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
	defer f2.Close()

	runtime.SetCPUProfileRate(5000)
	f, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer f.Close()

}

func WriteHeap() {
	f, err := os.Create("heapx")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()
	pprof.Lookup("heap").WriteTo(f, 0)
}
