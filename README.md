# ![qrcli](http://i.imgur.com/dhnYicm.png) qrcli

qrcli is a cli tool for creating QR codes.

## Usage

Install from source: 

```bash
$ go get github.com/rojters/qrcli
```

Writes to stdout by default:

```bash
$ qrcli "Hello, world" > qr.png
```

Options:

```bash
$ qrcli
usage: qrcli [flags] [text]
       qrcli -out qr.png "http://golang.org/"
  -elev="": GEO: elevation
  -file="": input file
  -hidden=false: Wifi: hidden (true|false)
  -lat="": GEO: deg N latitude
  -level="L": error correction level (L|M|Q|H)
  -long="": GEO: deg W longitude
  -out="": output file; stdout if empty
  -playstore="": Google play store uri, eg. org.example.app
  -pw="": Wifi: password
  -ssid="": Wifi: ssid
  -type="WPA": Wifi: network type (WPA|WEP)
```

More examples:

```bash
# wifi authentication (Android, may require plugins)
$ qrcli -ssid hotspot -pw secret -out wifi.png

# geographic information
$ qrcli -out geo.png -lat 69.7241573 -long 30.0583198 -elev 1000

# email
$ qrcli "mailto:hello@example.com" > email.png

# specify input file, use highest error correction level
$ qrcli -file mecard.txt -out contact.png -level H 

# Google play uri (opens app in play store on Android)
$ qrcli -playstore com.github.mobile -out app.png
```

For more information on diffrent types of contents see:
https://github.com/zxing/zxing/wiki/Barcode-Contents

## Attribution
This package imports and uses [code.google.com/p/rsc/qr][rsc-qr], a QR generation package by Russ Cox.

## License

MIT
