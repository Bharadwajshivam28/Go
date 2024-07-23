package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, _ := net.Dial("tcp", "localhost:3000")
    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Text to send: ")
        text, _ := reader.ReadString('\n')
        fmt.Fprintf(conn, text + "\n")
    }
}