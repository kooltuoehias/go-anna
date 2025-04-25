package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/speps/go-hashids/v2"
)

func main() {
	hd := hashids.NewData()
        hd.Salt = ""
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)

	log.Println("Anna Go here")
	for {
		var word string
		fmt.Scan(&word)
		integerToEncode, err := strconv.Atoi(word)
		if err != nil {
			Decode(h, word)
		} else {
			Encode(h, integerToEncode)
		}
	}
}

func Decode(h *hashids.HashID, input string) {
        d, _ := h.DecodeWithError(input)
        if len(d) == 1 {
                log.Println(PaddedResult(input, strconv.Itoa(d[0])))
        }
}

func Encode(h *hashids.HashID, input int) {
        e, _ := h.Encode([]int{input})
        log.Println(PaddedResult(strconv.Itoa(input), e))
}

func PaddedResult(input string, output string) string {
        return fmt.Sprintf("Anna figures out %-10s <---->    %s", input, output)
}
