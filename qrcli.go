package main

import (
	//"flag"
	"fmt"
	//"io/ioutil"
	"bufio"
	"log"
	"os"

	"code.google.com/p/rsc/qr"
	"github.com/spf13/cobra"
)

var level = flag.String("level", "L", "QR error correction level: L, M, Q, H")
var input = flag.String("file", "file.txt", "Input filename")
var output = flag.String("out", "qr.png", "Output filename")

type Wifi struct {
	Type     string
	SSID     string
	Password string
	Hidden   bool
}

func (w *Wifi) String() string {
	return fmt.Sprintf("WIFI:T:%s;S:%s;P:%s;H:%s;", w.Type, w.SSID, w.Password, w.Hidden)
}

type Email string

func (e *Email) String() string {
	return fmt.Sprintf("mailto:%s", *e)
}

func main() {

	//flag.Parse()
	//args := flag.Args()
	//fmt.Println(len(args))
	//for _, a := range args {
	//fmt.Println(a)
	//}
	//code, err := qr.Encode("WIFI:T:WPA;S:suggis;P:xei3eere;;", 0)

	code, err := qr.Encode("bajs", 0)
	//fmt.Println(args[0])

	if err != nil {
		log.Fatal(err)
	}

	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	f.Write(code.PNG())

	//err = ioutil.WriteFile("qr.png", code.PNG(), 0644)

}
