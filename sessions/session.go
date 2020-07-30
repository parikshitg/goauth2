package sessions

import (
	"fmt"
	"math/rand"
)

var oauth2states = make(map[string]string)

func RandomString(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func NewOauth2State() string {
	k := RandomString(16)
	v := RandomString(16)
	oauth2states[k] = v
	return fmt.Sprintf("%s%s", k, v)
}

func VerifyOauth2State(state string) bool {
	chars := []rune(state)
	k := string(chars[:15])
	v := string(chars[16:])
	if val, ok := oauth2states[k]; ok && v == val {
		return true
	}
	return false
}
