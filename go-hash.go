package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
)

func main() {
	it := os.Args[1]
	ht := os.Args[2]
	input := os.Args[3]

	var h hash.Hash
	if ht == "sha512" {
		h = sha512.New()
	} else if ht == "sha256" {
		h = sha256.New()
	} else if ht == "md5" {
		h = md5.New()
	}

	if it == "file" {
		f, err := os.Open(input)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if _, err := io.Copy(h, f); err != nil {
			panic(err)
		}
		fmt.Printf("%x\n", h.Sum(nil))
	} else if it == "string" {
		h.Write([]byte(input))
		fmt.Printf("%x\n", h.Sum(nil))
	}
}
