package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type CustomHandler struct{}

func (c *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL received: %s", r.URL.String())
	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	resp, err := client.Get("https://catfact.ninja/fact/")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	io.WriteString(w, string(body))
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL received: %s", r.URL.String())
	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	resp, err := client.Get("https://catfact.ninja/fact/")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(body))
}

func randomJsonHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, string("{\"Received url\":\""+r.URL.String()+"\"}"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

//func main() {
//	cfg := config.LoadFromFile("./config.yaml")
//
//	goagentClose := goagent.Init(cfg)
//	fmt.Println(os.Getenv("TA_LOG_LEVEL"))
//	defer goagentClose()
//	r := mux.NewRouter()
//	r.Handle("/cat-fact/random-endpoint/{id}", traceablehttp.NewHandler(http.HandlerFunc(serveHTTP), "/cat-fact/random-endpoint/{id}"))
//	r.HandleFunc("/random-endpoint/{id}", randomJsonHandler)
//
//	log.Fatal(http.ListenAndServe(":8080", r))
//
//}

func main() {
	// Choose the port you want to listen on
	port := "8443"
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listen.Close()
	fmt.Println("Server listening on port", port)

	for {
		// Wait for a connection
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// Handle the connection in a goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		// Read data from the connection
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("error matched")
			}
			fmt.Println("Error reading:", err)
			return
		}

		// Print the received data
		fmt.Print(string(buffer[:n]))

		// You can add your processing logic here based on the received data
		// For example, you could parse the data or perform some action

		// You can also send a response back to the client if needed
		conn.Write([]byte("Hello, client!"))
	}
}
