package main

import (
	"testing"

	"code.google.com/p/rsc/qr"
)

var leveltests = []struct {
	in       string
	expected qr.Level
}{
	{"L", qr.L},
	{"Q", qr.Q},
	{"M", qr.M},
	{"H", qr.H},
	{"", qr.L},
	{"__A__!/(##/(!!/(#", qr.L},
}

func TestLevel(t *testing.T) {
	for _, tt := range leveltests {
		got := handleLevelFlag(tt.in)
		if got != tt.expected {
			t.Errorf("handleLevelFlag(%s): expected %d, actual %d", tt.in, tt.expected, got)

		}
	}
}

var playtests = []struct {
	in       string
	expected string
}{
	{"", "market://details?id="},
	{"org.example.app", "market://details?id=org.example.app"},
}

func TestPlay(t *testing.T) {
	for _, tt := range playtests {
		got := handlePlay(tt.in)
		if got != tt.expected {
			t.Errorf("handlePlay(%s): expected %d, actual %d", tt.in, tt.expected, got)

		}
	}
}

type wifi struct {
	wifiType     string
	wifiSSID     string
	wifiPassword string
	wifiHidden   bool
}

var wifitests = []struct {
	in       wifi
	expected string
}{
	{
		wifi{
			"WPA",
			"hotspot",
			"secret",
			false,
		}, "WIFI:T:WPA;S:hotspot;P:secret;H:false;",
	},
	{
		wifi{
			"WEP",
			"hotspot",
			"",
			true,
		}, "WIFI:T:WEP;S:hotspot;P:;H:true;",
	},
}

func TestWifi(t *testing.T) {
	for _, tt := range wifitests {
		got := handleWifi(tt.in.wifiType, tt.in.wifiSSID, tt.in.wifiPassword, tt.in.wifiHidden)
		if got != tt.expected {
			t.Errorf(
				"handleWifi(%s, %s, %s, %t)\n: expected %s, actual %s",
				tt.in.wifiType,
				tt.in.wifiSSID,
				tt.in.wifiPassword,
				tt.in.wifiHidden,
				tt.expected,
				got,
			)
		}
	}
}

type geo struct {
	geoLat  string
	geoLong string
	geoElev string
}

var geotests = []struct {
	in       geo
	expected string
}{
	{
		geo{
			"69.7241573",
			"30.0583198",
			"1000",
		}, "geo:69.7241573,30.0583198,1000",
	},
	{
		geo{
			"",
			"",
			"",
		}, "geo:,,",
	},
}

func TestGeo(t *testing.T) {
	for _, tt := range geotests {
		got := handleGeo(tt.in.geoLat, tt.in.geoLong, tt.in.geoElev)
		if got != tt.expected {
			t.Errorf(
				"handleGeo(%s, %s, %s)\n: expected %d, actual %d",
				tt.in.geoLat,
				tt.in.geoLong,
				tt.in.geoElev,
				tt.expected,
				got,
			)
		}
	}
}
