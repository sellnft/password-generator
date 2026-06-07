package main

import (
	"crypto/rand"
	"math/big"
)

func generatePassword(digits *bool, symbols *bool, hardnessValue *string, n int) string {
	var res string

	hardnessValues := hardnesses[*hardnessValue]
	for range hardnessValues[0] {
		ind, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		index := int(ind.Int64())
		res += string(letters[index])
	}
	var usedPositions = make(map[int]bool)
	runes := []rune(res)
	if *digits == true {
		i := 0
		for i < n {
			ind1, _ := rand.Int(rand.Reader, big.NewInt(int64(len(res))))
			index1 := int(ind1.Int64())
			if !usedPositions[index1] {
				usedPositions[index1] = true
				ind2, _ := rand.Int(rand.Reader, big.NewInt(int64(len(DIGITS))))
				index2 := int(ind2.Int64())
				runes[index1] = DIGITS[index2]
				i++
			}
		}
	}

	if *symbols == true {
		i := 0
		for i < hardnessValues[1] {
			ind1, _ := rand.Int(rand.Reader, big.NewInt(int64(len(res))))
			index1 := int(ind1.Uint64())

			if !usedPositions[index1] {
				usedPositions[index1] = true

				ind2, _ := rand.Int(rand.Reader, big.NewInt(int64(len(specialSymbols))))
				index2 := int(ind2.Uint64())
				runes[index1] = specialSymbols[index2]
				i++
			}
		}
	}
	res = string(runes)
	return res
}
