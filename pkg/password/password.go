package password

import (
	"crypto/rand"
	"math/big"
	"os"
)

var PasswordCharacters = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"@", "#", "$", "%", "&", "*", "(", ")", "{", "}", "[", "]", "<", ">", "?",
}

func GeneratePassword(pc *PasswordConfig) string {

	alphabetsMaxIndex := 52
	maxIndex := len(PasswordCharacters)

	pass := PasswordCharacters[generateRandomNumber(alphabetsMaxIndex)]
	for c := 0; c < (pc.Length - 2); c++ {
		pass = pass + PasswordCharacters[generateRandomNumber(maxIndex)]
	}
	pass = pass + PasswordCharacters[generateRandomNumber(alphabetsMaxIndex)]

	return pass
}

func generateRandomNumber(max int) int {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		os.Exit(1)
	}
	return int(r.Int64())
}
