package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("===== Name + Hashing Program =====")

	var input1, input2 string
	fmt.Print("Please input Value 1: ")
	fmt.Scanln(&input1)
	fmt.Print("Please input Value 2: ")
	fmt.Scanln(&input2)
	proofMe(input1, input2)

}

func proofMe(txt1, txt2 string) {
	// MD5
	md5A := md5.Sum([]byte(txt1))
	md5B := md5.Sum([]byte(txt2))
	printResult("MD5", hex.EncodeToString(md5A[:]), hex.EncodeToString(md5B[:]))

	// SHA1
	sha1A := sha1.Sum([]byte(txt1))
	sha1B := sha1.Sum([]byte(txt2))
	printResult("SHA1", hex.EncodeToString(sha1A[:]), hex.EncodeToString(sha1B[:]))

	// SHA256
	sha256A := sha256.Sum256([]byte(txt1))
	sha256B := sha256.Sum256([]byte(txt2))
	printResult("SHA256", hex.EncodeToString(sha256A[:]), hex.EncodeToString(sha256B[:]))

	// SHA256
	sha512A := sha512.Sum512([]byte(txt1))
	sha512B := sha512.Sum512([]byte(txt2))
	printResult("SHA512", hex.EncodeToString(sha512A[:]), hex.EncodeToString(sha512B[:]))

	// SHA256
	sha3A := sha3.Sum256([]byte(txt1))
	sha3B := sha3.Sum256([]byte(txt2))
	printResult("SHA3-256", hex.EncodeToString(sha3A[:]), hex.EncodeToString(sha3B[:]))

}

func printResult(name, hashA, hashB string) {
	match := "No Match!"
	if hashA == hashB {
		match = "Match!"
	}

	fmt.Printf("\nHash (%s): \n", name)
	fmt.Println("Output A = ", hashA)
	fmt.Println("Output B = ", hashB)
	fmt.Println("=> ", match)
}
