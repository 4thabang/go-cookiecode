# Cookiecode

Cookiecode is a package that allows you to encrypt and decrypt your cookies for production grade use. Keep your cookies secure.

## Encode Example

```go
import (
  "time"
  "log"
  "net/http"

  cookiecode "github.com/4thabang/go-cookiecode"
)

func CookieEncoder(w http.ResponseWriter, r *http.Request) {
  value := map[string]string{
    "key": cookieKey, // <- Enter cookie key here
    "value": cookieValue, // <- Enter cookie value here
  }

  encoded, err := cookiecode.Encode(value)
  if err != nil {
    log.Print(err)
  }

  cookie := &http.Cookie{
    Name: value.Key,
    Value: encoded, // <- This is our encoded cookie value
    Expires: time.Time, // <- This is when our cookie is set to expire
  }

  // Set our encoded cookie value here.
  http.SetCokkie(w, cookie)
}
```

## Decode Example

```go
import (
  "log"
  "net/http"

  cookiecode "github.com/4thabang/go-cookiecode"
)

func CookieDecoder(w http.ResponseWriter, r *http.Request) {
  cookie, err := r.Cookie("key")
  if err != nil {
    log.Print(err)
  }

  value, err := cookiecode.Decode(cookie.Value)
  if err != nil {
    log.Print(err)
  }

  fmt.Println(value["value"])
}
```
