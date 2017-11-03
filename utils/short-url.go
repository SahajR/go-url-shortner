package utils

import (
	"strings"
	"regexp"
)

var characters = "23456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ-_"
var base = len(characters)

// To make sure the encoded string is at least 5 characters long
var adjustment = 10000000

// Encodes the ID of the URL in the database to a string that serves as its short code
func Encode(num int) string {
	encodedString := ""
	num += adjustment
	for num > 0 {
		encodedString = string(characters[num % base]) + encodedString
		num = num/base
	}
	return encodedString
}

// Decodes a short code to get the ID of the URL in the database
func Decode(str string) int {
	num,i := 0,0
	for i < len(str) {
		character := string(str[i])
		num = num * base + strings.Index(characters, character)
		i++
	}
	return num - adjustment
}

// Returns the complete URL given host and short code
func GetURL(host, code string) string {
	urlRegexp := "^(f|ht)tps?://"
	hasProtocol, _ := regexp.MatchString(urlRegexp, host)
	if hasProtocol {
		return host + "/" + code
	}
	return "https://" + host + "/" + code
}
