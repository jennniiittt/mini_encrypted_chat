package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println("Choose mode: server or client")
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Mode: ")
    mode, _ := reader.ReadString('\n')
    mode = strings.TrimSpace(mode)

    switch mode {
		case "server":
			StartServer()
		case "client":
			StartClient()
		default:
			fmt.Println("Invalid mode. Choose 'server' or 'client'.")
	}

}
