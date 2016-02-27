# <img src="http://i.imgur.com/dhnYicm.png" alt="qrcli" align="left" /> qrcli

[![travis-ci status](https://api.travis-ci.org/rojters/qrcli.png)](https://travis-ci.org/rojters/qrcli)

qrcli is a cli tool for creating QR codes.

## Install

Install from source: 

```
$ go get github.com/rojters/qrcli
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

Specify input and output file.

```
$ qrcli -file mecard.txt -out contact.png
```

Wifi authentication. Android devices can use this to automatically connect to a Wifi network (may require plugins). 

```
$ qrcli -wifi-ssid hotspot -wifi-pw secret -out wifi.png
```

Geographic information.

```
$ qrcli -geo-lat 12.357222 -geo-long -1.535278 -geo-elev 11 -out geo.png
```

Email.

```
$ qrcli "mailto:hello@example.com" > email.png
```

Google play URI. Opens app in Google Play on Android.

```
$ qrcli -googleplay com.google.android.youtube -out app.png
```

## License

MIT
