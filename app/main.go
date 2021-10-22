// +build ignore

package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	httpPort string
	appID    string
	tcpPort  string
)

func init() {
	flag.StringVar(&httpPort, "http", "8080", "the http port to listen on")
	flag.StringVar(&tcpPort, "tcp", "0", "the tcp port to listen on")
	flag.StringVar(&appID, "id", "foo", "the content to return in http response")
}

func handleByCode(w http.ResponseWriter, r *http.Request) {
	o := strings.Split(r.URL.Path, "/")
	code := 503
	if len(o) >= 3 {
		c, err := strconv.Atoi(o[2])
		if err == nil {
			code = c
		}
	}
	fmt.Printf("handleByCode: %v\n", code)
	w.WriteHeader(code)
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%v] request path: %v\n", time.Now(), r.URL.Path)
	w.Header().Add("app-server", appID)
	if strings.HasPrefix(r.URL.Path, "/code/") {
		handleByCode(w, r)
		return
	}
	fmt.Fprintf(w, "hello from %v", appID)
}

func handleTCP(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	payload := string(buf[:len])
	fmt.Printf("received tcp conn\nremote address %v\nlocal address: %v\ndata\n%v",
		conn.RemoteAddr(), conn.LocalAddr(), string(payload))
	lines := strings.Split(payload, "\n")
	cmd := "reset"
	for _, i := range lines {
		if strings.Contains(i, "cmd") {
			cmd = strings.TrimSpace(strings.Split(i, " ")[1])
			break
		}
	}
	fmt.Printf("execute command:%v\n", cmd)
	switch cmd {
	// not working yet.
	// case "echo":
	// 	conn.Write([]byte("HTTP/1.1 200 OK\n"))
	// 	conn.Write(buf[:len])
	// 	// go idle not reset.
	// 	time.Sleep(time.Second * 600)
	// break
	case "hang":
		time.Sleep(3600 * time.Second)
		break
	default:
		conn.Close()
	}
	// Send a response back to person contacting us.
	// conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
}

// TODO(incfly): not necessarily, since http server can reset the connection as well.
// https://gist.github.com/incfly/68098fa245c54f6e5abf9d9680b5546d
func startTcpServer() {
	if tcpPort == "0" {
		fmt.Println("--tcp is not specified, skip.")
		return
	}
	// Listen for incoming connections.
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", tcpPort))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		panic("failed to listen tcp server")
	}
	defer l.Close()
	fmt.Println("tcp server initialized.")
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleTCP(conn)
	}
}

func main() {
	flag.Parse()
	fmt.Printf("hello world, starting the server at port %v\n", httpPort)
	http.HandleFunc("/", handleHTTP)
	go startTcpServer()
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%v", httpPort), nil); err != nil {
		panic(fmt.Sprintf("failed to listen http server: %v", err))
	}
}
