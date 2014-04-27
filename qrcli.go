package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"code.google.com/p/rsc/qr"
)

var (
	// global
	inputFile  = flag.String("file", "", "input file")
	outputFile = flag.String("out", "", "output file; stdout if empty")
	level      = flag.String("level", "L", "error correction level (L|M|Q|H)")

	// wifi
	wifiType     = flag.String("type", "WPA", "Wifi: network type (WPA|WEP)")
	wifiSSID     = flag.String("ssid", "", "Wifi: ssid")
	wifiPassword = flag.String("pw", "", "Wifi: password")
	wifiHidden   = flag.Bool("hidden", false, "Wifi: hidden (true|false)")

	// geo
	geoLat  = flag.String("lat", "", "GEO: deg N latitude")
	geoLong = flag.String("long", "", "GEO: deg W longitude")
	geoElev = flag.String("elev", "", "GEO: elevation")

	// Google play
	play = flag.String("playstore", "", "Google play store uri, eg. org.example.app")
)

func usage() {
	fmt.Fprintf(os.Stderr,
		"usage: qrcli [flags] [text]\n"+
			"       qrcli -out qr.png \"http://golang.org/\"\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func handleLevelFlag(lvl string) qr.Level {
	switch lvl {
	case "M":
		return qr.M
	case "Q":
		return qr.Q
	case "H":
		return qr.H
	}
	return qr.L
}

func main() {

	flag.Parse()

	if flag.NArg() == 0 && *inputFile == "" && *wifiSSID == "" && *geoLat == "" && *play == "" {
		// no input
		usage()
	}

	var text string
	args := flag.Args()

	switch {
	case *wifiSSID != "":
		text = handleWifi()
	case *geoLat != "":
		text = handleGeo()
	case *play != "":
		text = handlePlay()
	case *inputFile != "":
		text = handleInput()
	case args[0] != "":
		text = args[0]
	}

	encode(text)
}

func handleInput() string {
	dat, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(dat)
}

func handlePlay() string {
	return fmt.Sprintf("market://details?id=%s", *play)
}

func handleGeo() string {
	return fmt.Sprintf("geo:%s,%s,%s", *geoLat, *geoLong, *geoElev)
}

func handleWifi() string {
	return fmt.Sprintf("WIFI:T:%s;S:%s;P:%s;H:%s;", *wifiType, *wifiSSID, *wifiPassword, *wifiHidden)
}

func encode(text string) {
	code, err := qr.Encode(text, handleLevelFlag(*level))

	if err != nil {
		log.Fatal(err)
	}

	write(code.PNG())
}

func write(img []byte) {
	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		f.Write(img)
	} else {
		f := bufio.NewWriter(os.Stdout)
		defer f.Flush()
		f.Write(img)
	}
}
