<h1 align="center">:cookie: Cookiecode :cookie:</h1>
<br/>

Cookiecode is a package that allows you to encrypt and decrypt your cookies for production grade use. Keep your cookies secure.

## Installation

`go get github.com/4thabang/go-cookiecode`

## Demo Usage

### .env variables

`WIP`

### Encode Code Snippet

```go
import (
    "time"
    "log"
    "net/http"

    cookiecode "github.com/4thabang/go-cookiecode"
)

func CookieEncoder(w http.ResponseWriter, r *http.Request) {
  // TODO: Add godotenv .env examples

  value := map[string]string{
    "key":   cookieKey,
    "value": cookievalue,
  }

  encoded, err := cookiecode.Encode(value)
  if err != nil {
    log.Printf("error: %v\n", err)
  }

  cookie := &http.Cookie{
    Name: value.Key,
    Value: encoded, // <- This is our encoded cookie value
    Expires: time.Time, // <- This is when our cookie is set to expire
  }

  // Set our encoded cookie value here
  http.SetCookie(w, cookie)
}
```

### Decode Code Snippet

```go
import (
    "log"
    "net/http"

    cookiecode "github.com/4thabang/go-cookiecode"
    )

func CookieDecoer(w http.ResponseWriter, r *http.Request) {
  cookie, err := r.Cookie("key")
    if err != nil {
      log.Print(err)
    }

  value, err := cookiecode.Decode(cookie.Name, cookie.Value)
    if err != nil {
      log.Print(err)
    }

  fmt.Println(value["value"])
}
```

# License

`WIP`
