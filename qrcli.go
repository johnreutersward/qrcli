package main

import (
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

const (
	version = "0.1.0"
)

func printUsage() {
	fmt.Fprintf(os.Stderr,
		"usage: qrcli [flags] [text]\n"+
			"       qrcli -out qr.png \"http://golang.org/\"\n")
	flag.PrintDefaults()
	os.Exit(0)
}

func printVersion() {
	fmt.Fprintf(os.Stdout, "qrcli v%s\n", version)
	os.Exit(0)
}

func handleLevelFlag(lvl string) qr.ErrorCorrectionLevel {
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
		showHelp    = flag.Bool("help", false, "Show usage help")
		showVersion = flag.Bool("version", false, "Show version")
		inputFile   = flag.String("file", "", "Input file")
		outputFile  = flag.String("out", "", "Output file")
		level       = flag.String("level", "L", "Error correction level (L|M|Q|H)")
		size        = flag.Int("size", 250, "Output image size")

		// wifi
		wifiAuth     = flag.String("wifi-auth", "WPA", "Wifi authentication (WPA|WEP|nopass)")
		wifiSSID     = flag.String("wifi-ssid", "", "Wifi SSID")
		wifiPassword = flag.String("wifi-pw", "", "Wifi password")
		wifiHidden   = flag.Bool("wifi-hidden", false, "Wifi hidden (true|false)")

		// geo
		geoLat  = flag.String("geo-lat", "", "Geo deg N latitude")
		geoLong = flag.String("geo-long", "", "Geo deg W longitude")
		geoElev = flag.String("geo-elev", "", "Geo elevation")

		// Google play
		play = flag.String("googleplay", "", "Google Play uri, e.g. \"org.example.app\"")
	)

	flag.Parse()

	if *showVersion {
		printVersion()
	}

	if (flag.NArg() == 0 && *inputFile == "" && *wifiSSID == "" && *geoLat == "" && *play == "") || *showHelp {
		printUsage()
	}

	var text string
	args := flag.Args()

	switch {
	case *wifiSSID != "":
		text = handleWifi(*wifiAuth, *wifiSSID, *wifiPassword, *wifiHidden)
	case *geoLat != "":
		text = handleGeo(*geoLat, *geoLong, *geoElev)
	case *play != "":
		text = handlePlay(*play)
	case *inputFile != "":
		text = handleInput(*inputFile)
	case args[0] != "":
		text = args[0]
	}

	qrcode, err := qr.Encode(text, handleLevelFlag(*level), qr.Auto)
	if err != nil {
		log.Fatal(err)
	}

	qrcode, err = barcode.Scale(qrcode, *size, *size)
	if err != nil {
		log.Fatal(err)
	}

	var f *os.File
	if *outputFile != "" {
		f, err = os.Create(*outputFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f = os.Stdout
	}
	defer f.Close()

	if err := png.Encode(f, qrcode); err != nil {
		log.Fatal(err)
	}
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
