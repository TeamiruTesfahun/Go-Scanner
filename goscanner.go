package main 

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func Scanport(host string, port int, wg *sync.WaitGroup){
	defer wg.Done()

	address := host + ":" + strconv.Itoa(port)

	conn, err := net.DialTimeout("tcp", address, 1*time.Second)

	if err != nil {
		//fmt.Printf("Port %d is closed\n", port)
		return
	}

	defer conn.Close()

	fmt.Printf("Port %d is open\n", port)

}

func main(){
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run try.go <host>")
		os.Exit(1)
	}

	host := os.Args[1]
	startPort := 1
	endPort := 1000

	var wg sync.WaitGroup

	fmt.Printf("Scanning ports %d to %d on %s...\n", startPort, endPort, host)

	for port:= startPort; port <= endPort; port++ {
		wg.Add(1)

		go Scanport(host, port, &wg)
	}

	wg.Wait()

	fmt.Println("Scan completed")
	
}