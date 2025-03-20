package mailer

import (
	"math/rand"
	"time"
)

func GenerateOtp() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	str := []int32("asdflkjqeqpixzcvm1234567890")
	hash := ""
	for i := 0; i < 7; i++ {
		index := rand.Intn(len(str) - 1)
		hash += string(str[index])
	}
	return hash
}
