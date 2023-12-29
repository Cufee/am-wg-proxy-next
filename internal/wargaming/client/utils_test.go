package client

import "testing"

func TestParseProxySettings(t *testing.T) {
	bucket, err := parseProxySettings("user:password@host:80?wgAppId=your_app_id&maxRps=20&realm=na", "not-valid-id", 99)
	if err != nil {
		t.Error(err)
	}
	if bucket.port != "80" {
		t.Error("port not parsed")
	}
	if bucket.host != "host" {
		t.Error("host not parsed")
	}
	if bucket.username != "user" {
		t.Error("username not parsed")
	}
	if bucket.password != "password" {
		t.Error("password not parsed")
	}
	if bucket.wgAppId != "your_app_id" {
		t.Error("wgAppId not parsed")
	}
	if bucket.rps != 20 {
		t.Error("maxRps not parsed")
	}
	if bucket.realm != "NA" {
		t.Error("realm not parsed")
	}

	bucket2, err := parseProxySettings("user1:password1@host1:8080?wgAppId=your_app_id1&maxRps=5", "not-valid-id", 99)
	if err != nil {
		t.Error(err)
	}
	if bucket2.port != "8080" {
		t.Error("port not parsed")
	}
	if bucket2.host != "host1" {
		t.Error("host not parsed")
	}
	if bucket2.username != "user1" {
		t.Error("username not parsed")
	}
	if bucket2.password != "password1" {
		t.Error("password not parsed")
	}
	if bucket2.wgAppId != "your_app_id1" {
		t.Error("wgAppId not parsed")
	}
	if bucket2.rps != 5 {
		t.Error("maxRps not parsed")
	}
	if bucket2.realm != "" {
		t.Error("realm not parsed")
	}
}
