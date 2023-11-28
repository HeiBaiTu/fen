package fen

import (
	"log"
)

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		log.Printf(WarningColor, "PORT is undefined. Using port :8080 by default")
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}
