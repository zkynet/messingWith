package memprofile

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func GetMemoryProfile() {
	runtime.SetCPUProfileRate(5000)
	f, err := os.Create("mem.pprof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
func WriteHeap() {
	f, err := os.Create("heapx")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()
	pprof.Lookup("heap").WriteTo(f, 0)
}
