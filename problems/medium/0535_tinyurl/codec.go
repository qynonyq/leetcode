package main

import (
	"log"
	"math/rand"
	"path"
	"regexp"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	digits    = "0123456789"
	letters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphanums = digits + letters
	tinyUrl   = "https://tinyurl.com/"
)

type Codec struct {
	repo map[string]string
	rx   *regexp.Regexp
}

func Constructor() Codec {
	return Codec{
		repo: make(map[string]string, 1),
		rx: regexp.MustCompile("^(ftp|https?)://.?[-a-zA-Z0-9@:%._+~#=]{0,62}" +
			"\\b([-a-zA-Z0-9@:%_+.~#?&/=]*)$"),
	}
}

func (c *Codec) encode(longUrl string) string {
	if len(longUrl) < 1 || len(longUrl) > 10000 {
		log.Println("wrong longUrl size")
		return ""
	}
	if ok := c.rx.MatchString(longUrl); !ok {
		log.Println("invalid url format")
		return ""
	}
	encoded := randUrl(6)
	c.repo[encoded] = longUrl
	return tinyUrl + encoded
}

func (c *Codec) decode(shortUrl string) string {
	base := path.Base(shortUrl)
	return c.repo[base]
}

func randUrl(n int) string {
	rr := make([]byte, n)
	for i := range rr {
		rr[i] = alphanums[rand.Intn(len(alphanums))]
	}
	return string(rr)
}
