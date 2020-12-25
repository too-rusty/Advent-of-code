package main

import "fmt"

const (
	mod = 20201227
)

func main() {

	// subject :=
	// value := 1
	// loop_size :=

	// value = (value * subject)%mod

	// value := 1
	// for i := 0; i < 8; i++ {
	// 	value *= subject
	// 	value %= mod
	// }
	// fmt.Println(value)

	cardPublicKey, doorPublicKey := 14082811, 5249543
	cardLoopSize, doorLoopSize := getLoopSize(cardPublicKey), getLoopSize(doorPublicKey)
	fmt.Println(transform(doorPublicKey, cardLoopSize))
	fmt.Println(transform(cardPublicKey, doorLoopSize))

}

func max(arr ...int) int {
	mx := 0
	for _, v := range arr {
		if v > mx {
			mx = v
		}
	}
	return mx
}

func transform(subject, loop_size int) (val int) {
	val = 1
	for i := 0; i < loop_size; i++ {
		val *= subject
		val %= mod
	}
	return
}

func getLoopSize(cardPublicKey int) (cardLoopSize int) {
	card := 1
	const subject = 7
	for {
		if card == cardPublicKey {
			break
		}
		card *= subject
		card %= mod
		cardLoopSize++
	}
	return
}
