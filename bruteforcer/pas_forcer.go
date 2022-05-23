package bruteforcer

import (
	"crypto/md5"
	"encoding/hex"
)

func Force(startSymb rune, endSymb rune, hexPassword string, channel chan string) {
	var pasSymbols [6]rune
	for firstSymb := startSymb; firstSymb <= endSymb; firstSymb++ {
		pasSymbols[0] = firstSymb

		for secSymb := 33; secSymb < 127; secSymb++ {
			pasSymbols[1] = rune(secSymb)

			for thirdSymb := 33; thirdSymb < 127; thirdSymb++ {
				pasSymbols[2] = rune(thirdSymb)

				for fourthSymb := 33; fourthSymb < 127; fourthSymb++ {
					pasSymbols[3] = rune(fourthSymb)

					for fivethSymb := 33; fivethSymb < 127; fivethSymb++ {
						pasSymbols[4] = rune(fivethSymb)

						for sixthSymb := 33; sixthSymb < 127; sixthSymb++ {
							pasSymbols[5] = rune(sixthSymb)
							password := asciiListToString(pasSymbols[:])
							passwordMd5 := getMD5(password)
							if passwordMd5 == hexPassword {
								channel <- "The target passowrd is: " + password
							}
						}

					}
				}
			}
		}
	}
}

func asciiListToString(asciiList []rune) string {
	var convertedText string

	for _, char := range asciiList {
		convertedText += string(char)
	}

	return convertedText
}

func getMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
