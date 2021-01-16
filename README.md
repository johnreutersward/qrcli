# qrcli

[![Build Status](https://travis-ci.org/johnreutersward/qrcli.svg?branch=master)](https://travis-ci.org/johnreutersward/qrcli)
[![Go Reference](https://pkg.go.dev/badge/github.com/johnreutersward/qrcli.svg)](https://pkg.go.dev/github.com/johnreutersward/qrcli)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnreutersward/qrcli)](https://goreportcard.com/report/github.com/johnreutersward/qrcli)

![qrcli](qr.png?raw=true "qrcli")

> qrcli is a cli tool for creating QR codes.

## Install

Binary releases: 

https://github.com/johnreutersward/qrcli/releases

Or build from source:

```
$ go get github.com/johnreutersward/qrcli
```

## Usage

Writes to stdout by default:

```
$ qrcli "Hello, world" > qr.png
```

## Options

```
$ usage: qrcli [flags] [text]
       qrcli -out qr.png "http://golang.org/"
  -file string
      Input file
  -geo-elev string
      Geo elevation
  -geo-lat string
      Geo deg N latitude
  -geo-long string
      Geo deg W longitude
  -googleplay string
      Google Play uri, e.g. "org.example.app"
  -help
      Show usage help
  -level string
      Error correction level (L|M|Q|H) (default "L")
  -out string
      Output file
  -size int
      Output image size (default 250)
  -version
      Show version
  -wifi-auth string
      Wifi authentication (WPA|WEP|nopass) (default "WPA")
  -wifi-hidden
      Wifi hidden (true|false)
  -wifi-pw string
      Wifi password
  -wifi-ssid string
      Wifi SSID
```

## Examples

Specify input and output file:

```
$ qrcli -file mecard.txt -out contact.png
```

Wifi authentication - Android devices can use this to automatically connect to a Wifi network (may require plugins):

```
$ qrcli -wifi-ssid hotspot -wifi-pw secret -out wifi.png
```

Geographic information:

```
$ qrcli -geo-lat 12.357222 -geo-long -1.535278 -geo-elev 11 -out geo.png
```

Email:

```
$ qrcli "mailto:hello@example.com" > email.png
```

Google play URI - Opens app in Google Play on Android:

```
$ qrcli -googleplay com.google.android.youtube -out app.png
```

## License

MIT
