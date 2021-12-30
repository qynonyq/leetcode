package main

import "fmt"

func main() {
	codec := Constructor()
	tiny := codec.encode("ftp://174.123.452.34/directory/file")
	fmt.Println(tiny)
	long := codec.decode(tiny)
	fmt.Println(long)
}
