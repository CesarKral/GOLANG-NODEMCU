package main

import (
	"fmt"
	"net/http"
	"net"
	"bufio"
)

func main()  {
	
	http.HandleFunc("/nodemcu", func(res http.ResponseWriter, req *http.Request) {
		conn, err := net.Dial("tcp", "192.168.1.39:80")
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		conn.Write([]byte("11"))
		resp, _, _ := bufio.NewReader(conn).ReadLine()
		fmt.Println(string(resp))
	})
	
	http.ListenAndServe(":8080", nil)

}
