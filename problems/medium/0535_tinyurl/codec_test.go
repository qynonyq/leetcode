package main

import (
	"testing"
)

var (
	tinyLen = len(tinyURL + randURL(6))
)

func TestCodecEncodeDecode(t *testing.T) {
	tests := map[string]struct {
		url         string
		expectedLen int
	}{
		"ok":        {url: "https://google.com", expectedLen: tinyLen},
		"ok2":       {url: "ftp://174.123.452.34/directory/file", expectedLen: tinyLen},
		"empty_url": {url: "", expectedLen: 0},
		"wrong_url": {url: "htt://google.com", expectedLen: 0},
	}

	codec := Constructor()

	for name, in := range tests {
		t.Run(name, func(t *testing.T) {
			res := codec.encode(in.url)
			if in.expectedLen != len(res) {
				t.Errorf("expected %d, got %d", in.expectedLen, len(res))
			}
			url := codec.decode(res)
			if in.expectedLen == 0 {
				if in.expectedLen != len(url) {
					t.Errorf("expected %q, got %q", in.expectedLen, len(url))
				}
				return
			}
			if in.url != url {
				t.Errorf("expected %q, got %q", in.url, url)
			}
		})
	}
}
