package mylib

import "testing"

func TestGetOutlineList(t *testing.T) {
	const filepath = "Feeder_test.opml"
	siteList := GetOutlineList(filepath)
	for _, site := range siteList {
		t.Log(site)
	}
}
