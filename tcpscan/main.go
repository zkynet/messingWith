package main

import (
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var timeout = 1 * time.Second
var IPMap = make(map[string]net.Conn)
var IPMapLock = sync.Mutex{}
var WaitGroup = sync.WaitGroup{}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	for i := 1; i <= 255; i++ {
		WaitGroup.Add(1)
		go raw_connect("192.168.1."+strconv.Itoa(i), "1234")
	}
	WaitGroup.Wait()

	for i, v := range IPMap {
		log.Println(i, v)
	}
	log.Println(IPMap)
}

func raw_connect(host string, port string) {
	defer func() {
		if r := recover(); r != nil {
			WaitGroup.Done()
			log.Println(r)
		} else {
			WaitGroup.Done()
		}
	}()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err == nil {
		IPMapLock.Lock()
		IPMap[host] = conn
		IPMapLock.Unlock()
	}
}
