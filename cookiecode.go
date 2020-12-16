package cookiecode

import (
	"fmt"

	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
)

// Encrypt holds the values for our token encoding/decoding.
type Encrypt struct {
	HashKey  []byte
	BlockKey []byte
}

type i interface{}

// Keys allows us to store our encryption keys securely for re-use.
func Keys(v map[string]string) (*securecookie.SecureCookie, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	encrypt := &Encrypt{
		HashKey:  []byte(v["HASH_KEY"]),
		BlockKey: []byte(v["BLOCK_KEY"]),
	}
	secure := securecookie.New(encrypt.HashKey, encrypt.BlockKey)

	return secure, nil
}

// Encode allows us to encode our cookie in order to keep it secure, safe and unexposed.
func Encode(key string, value map[string]string) (string, error) {
	// TODO: fix invalid map argument
	secure, err := Keys(map[string]string)
	if err != nil {
		fmt.Println(err)
	}

	// cookiecode.Encode("access_token", cookie.Value)
	encode, err := secure.Encode(key, value)
	if err != nil {
		fmt.Println(err)
	}

	return encode, nil
}

// Decode allows us to decode our cookie in order to consume it safely, awaay from prying eyes.
func Decode(key, cookie string) (map[string]string, error) {
	// TODO: fix invalid map argument
	secure, err := Keys(map[string]string)
	if err != nil {
		fmt.Println(err)
	}

	var value map[string]string
	err = secure.Decode(key, cookie, &value)
	if err != nil {
		fmt.Println(err)
	}

	return value, nil
}
