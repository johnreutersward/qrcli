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

	flag.Parse()

	if flag.NArg() == 0 && *inputFile == "" && *wifiSSID == "" && *geoLat == "" && *play == "" {
		// no input
		usage()
	}

	var text string
	args := flag.Args()

	switch {
	case *wifiSSID != "":
		text = handleWifi(*wifiType, *wifiSSID, *wifiPassword, *wifiHidden)
	case *geoLat != "":
		text = handleGeo(*geoLat, *geoLong, *geoElev)
	case *play != "":
		text = handlePlay(*play)
	case *inputFile != "":
		text = handleInput(*inputFile)
	case args[0] != "":
		text = args[0]
	}

	code, err := qr.Encode(text, handleLevelFlag(*level))

	if err != nil {
		log.Fatal(err)
	}

	write(code.PNG(), *outputFile)

}

func handleInput(inputFile string) string {
	dat, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(dat)
}

func handlePlay(play string) string {
	return fmt.Sprintf("market://details?id=%s", play)
}

func handleGeo(geoLat, geoLong, geoElev string) string {
	return fmt.Sprintf("geo:%s,%s,%s", geoLat, geoLong, geoElev)
}

func handleWifi(wifiType string, wifiSSID string, wifiPassword string, wifiHidden bool) string {
	return fmt.Sprintf("WIFI:T:%s;S:%s;P:%s;H:%t;", wifiType, wifiSSID, wifiPassword, wifiHidden)
}

func write(img []byte, outputFile string) {
	if outputFile != "" {
		f, err := os.Create(outputFile)
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
