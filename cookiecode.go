package cookiecode

import (
	"fmt"

	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
)

/*
	TODO: how will the users add their keys
	TODO: better godocs comments
*/

// Encrypt holds the values for our token encoding/decoding.
type Encrypt struct {
	HashKey  []byte
	BlockKey []byte
}

// Keys allows us to store our encryption keys securely for re-use.
func Keys(v map[string]string) (*securecookie.SecureCookie, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	encrypt := &Encrypt{
		HashKey:  []byte(v["HASH_KEY"]),  // AES 256-bit
		BlockKey: []byte(v["BLOCK_KEY"]), // AES 128-bit
	}
	secure := securecookie.New(encrypt.HashKey, encrypt.BlockKey)

	return secure, nil
}

/* [sample code below]

value := map[string]interface{}{
	"token": cookieValue,
}

encoded, err := cookiecode.Encode(value)
if err != nil {
	fmt.Println(err)
}

cookie := &http.Cookie{
	Name: "token",
	Value: encoded, <-- our encoded cookie value
	Expires: time.Time, <-- time to expiry
}

*/

// EncodeType is a struct that houses the value of our key which will be
// determined by the user at runtime.
type EncodeType struct {
	Key string
}

// Encode allows us to encode our cookie in order to keep it secure, safe and unexposed.
func Encode(value map[string]string) (string, error) {
	// TODO: fix invalid map argument
	secure, err := Keys(value)
	if err != nil {
		fmt.Println(err)
	}

	e := &EncodeType{
		Key: value["key"],
	}

	// cookiecode.Encode("access_token", cookie.Value)
	encode, err := secure.Encode(e.Key, value)
	if err != nil {
		fmt.Println(err)
	}

	return encode, nil
}

/* [sample code below]

cookie, err := r.Cookie("key")
if err != nil {
	log.Print(err)
}

value, err := cookiecode.Decode(cookie.Value)
if err != nil {
	log.Print(err)
}

return value["key"]
*/

// Decode allows us to decode our cookie in order to consume it safely, awaay from prying eyes.
func Decode(key, cookie string) (map[string]string, error) {
	// TODO: fix invalid map argument
	vm := make(map[string]string)

	secure, err := Keys(vm)
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
