package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func isIPv4SubString(subString string) bool {
	if len(subString) > 1 && strings.HasPrefix(subString, "0") || len(subString) == 0 {
		return false
	}

	regexString := fmt.Sprintf("[0-9]{%v}", len(subString))

	match, _ := regexp.MatchString(regexString, subString)

	if !match {
		return false
	}

	subStringInt, _ := strconv.Atoi(subString)

	if subStringInt < 0 || subStringInt > 255 {
		return false
	}

	return true
}

func isIPv4(ip string) bool {
	subStrings := strings.Split(ip, ".")

	if len(subStrings) != 4 {
		return false
	}

	for _, subString := range subStrings {
		if !isIPv4SubString(subString) {
			return false
		}
	}

	return true
}

func isIPv6SubString(subString string) bool {
	if len(subString) > 4 || len(subString) < 1 {
		return false
	}

	regexString := fmt.Sprintf("[a-f,A-F,0-9]{%v}", len(subString))

	match, _ := regexp.MatchString(regexString, subString)

	if !match {
		return false
	}

	return true
}

func isIPv6(ip string) bool {
	subStrings := strings.Split(ip, ":")

	if len(subStrings) != 8 {
		return false
	}

	for _, subString := range subStrings {
		if !isIPv6SubString(subString) {
			return false
		}
	}

	return true
}

func validateIp(ip string) string {

	if isIPv4(ip) {
		return "IPv4"
	} else if isIPv6(ip) {
		return "IPv6"
	}

	return "Neither"
}

func main() {
	fmt.Println(validateIp("172.16.254.1"))                           // "IPv4"
	fmt.Println(validateIp("172.016.254.1"))                          // "Neither"
	fmt.Println(validateIp("2001:0db8:85a3:0:0:8A2E:0370:7334"))      // "IPv6"
	fmt.Println(validateIp("2001:0db8:85a3:0:0:8A2E:0370:7334:7334")) // "Neither"
	fmt.Println(validateIp("2001:0db8:85a3:0:0:8A2E:0370:7334Z"))     // "Neither"
	fmt.Println(validateIp("256.256.256.256"))                        // "Neither"
	fmt.Println(validateIp("1e1.4.5.6"))                              // "Neither"
	fmt.Println(validateIp("1.0.1."))                                 // "Neither"
	fmt.Println(validateIp("2001:0db8:85a3:00000:0:8A2E:0370:7334"))  // "Neither"

}
