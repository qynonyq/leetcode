package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"path"
	"regexp"
)

const (
	digits    = "0123456789"
	letters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphanums = digits + letters
	tinyURL   = "https://tinyurl.com/"
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

func (c *Codec) encode(longURL string) string {
	if len(longURL) < 1 || len(longURL) > 10000 {
		log.Println("wrong longURL size")
		return ""
	}
	if ok := c.rx.MatchString(longURL); !ok {
		log.Println("invalid url format")
		return ""
	}
	encoded := randURL(6)
	c.repo[encoded] = longURL
	return tinyURL + encoded
}

func (c *Codec) decode(shortURL string) string {
	base := path.Base(shortURL)
	return c.repo[base]
}

func randURL(n int) string {
	rr := make([]byte, n)
	for i := range rr {
		max := big.NewInt(int64(len(alphanums)))
		random, _ := rand.Int(rand.Reader, max)
		idx := random.Int64()
		rr[i] = alphanums[idx]
	}
	return string(rr)
}
