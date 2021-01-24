# Cookiecode
Cookiecode is a package that allows you to encrypt and decrypt your cookies for production grade use. Keep your cookies secure.

## Example Code
```go
package main

import (
  "time"
  "log"
  "net/http"
  
  cookiecode "github.com/4thabang/go-cookiecode"
)

func CookieHandler(w http.Response, r *http.Request) {
  value := map[string]string{
    "key": cookieKey,
    "value": cookieValue
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
  
  http.SetCokkie(w, cookie)
}
```
