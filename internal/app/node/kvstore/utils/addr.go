package utils

import "strings"

func EncodeIPAddr(ip string) string {
	return strings.ReplaceAll(ip, ":", "-")
}

func DecodeIPAddr(ip string) string {
	return strings.ReplaceAll(ip, "-", ":")
}
