package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//make a byte slice init w 99999 empty elements
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	lw := logWriter{}
	//now to do it in one go....lol
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println((string(bs)))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
