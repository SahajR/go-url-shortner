package utils

import (
	"strings"
)

var characters = "23456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ-_"
var base = len(characters)

// To make sure the encoded string is at least 5 characters long
var adjustment = 10000000

func Encode(num int) string {
	encodedString := ""
	num += adjustment
	for num > 0 {
		encodedString = string(characters[num % base]) + encodedString
		num = num/base
	}
	return encodedString
}

func Decode(str string) int {
	num,i := 0,0
	for i < len(str) {
		character := string(str[i])
		num = num * base + strings.Index(characters, character)
		i++
	}
	return num - adjustment
}