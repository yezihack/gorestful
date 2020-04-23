[![](https://img.shields.io/badge/GoDoc-gorestful-orange)](https://godoc.org/github.com/yezihack/gorestful)
[![Coverage Status](https://coveralls.io/repos/github/yezihack/gorestful/badge.svg)](https://coveralls.io/github/yezihack/gorestful)
# gorestful
Gorestful is golang simple web 

# Install 
`go get github.com/yezihack/gorestful`

# How
```go
import "github.com/yezihack/gorestful"

restful := gorestful.New()
restful.GET("/", func(writer http.ResponseWriter, request *http.Request) {
    writer.WriteHeader(200)
    writer.Write([]byte("ok"))
})
restful.Run(":8080")
```

# Feature
Support GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, TRACE
