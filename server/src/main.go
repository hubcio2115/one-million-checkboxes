package main

import (
	"net/http"
)

func pow(base uint8, pow uint8) uint8 {
	result := uint8(1)

	for i := uint8(0); i < pow; i++ {
		result *= base
	}

	return result
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/checkboxes", func(w http.ResponseWriter, r *http.Request) {
		bits := make([]bool, 1_000_000)

		for i := 0; i < 1_000_00; i++ {
			bits[i] = i%2 == 0
		}

		buf := make([]byte, 125_000)

		// This has to be like this other way first byte is filled with zeros
		currentByte := uint8(0)
		for i := -8; i < 999_992; i++ {
			if i%8 == 0 && i >= 0 {
				buf[i/8] = currentByte
				currentByte = 0
			}

			if bits[i+8] {
				if i < 0 {
					currentByte += pow(2, uint8(-i%8))
				} else {
					currentByte += pow(2, uint8(i%8))
				}
			}
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(buf)
	})

	http.ListenAndServe(":8080", mux)
}
