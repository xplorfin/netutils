[![Coverage Status](https://coveralls.io/repos/github/xplorfin/netutils/badge.svg?branch=master)](https://coveralls.io/github/xplorfin/netutils?branch=master)
[![Renovate enabled](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://app.renovatebot.com/dashboard#github/xplorfin/netutils)
[![Build status](https://github.com/xplorfin/netutils/workflows/test/badge.svg)](https://github.com/xplorfin/netutils/actions?query=workflow%3Atest)
[![Build status](https://github.com/xplorfin/netutils/workflows/goreleaser/badge.svg)](https://github.com/xplorfin/netutils/actions?query=workflow%3Agoreleaser)
[![](https://godoc.org/github.com/xplorfin/netutils?status.svg)](https://pkg.go.dev/github.com/xplorfin/netutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/xplorfin/netutils)](https://goreportcard.com/report/github.com/xplorfin/netutils)

# Netutils

This is a series of networking utilities and test wrappers by [entropy](http://entropy.rocks/) for building robust networked services in golang. See the godoc for details

 
 # What can I do with this?
 
The godoc should cover most of it. I've highlighted a few things below and will add more examples as time goes on. Examples are also present in the godoc
 
Mocking:
 
One thing peculiarity of the [`httmock`](https://github.com/jarcoal/httpmock) library is you can't actually pass it a handler. `WrapHandler` let's you do so:
 
 Handler Mocking:
 
 (see [`mock_test`](testutils/mock_test) for a more detailed example)
 ```go
    package main
    import (
    	"github.com/jarcoal/httpmock"
    )

        func main(){
        httpmock.Activate()
        defer httpmock.Deactivate()
        ctx := context.Background()
        requestCount := 0
        testServer := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
            rw.WriteHeader(200)
        })
    
        httpmock.RegisterResponder("POST", "https://api.github.com/graphql", testutils.WrapHandler(testServer))
}   
```

There's also a fasthttp module for mocking fasthttp servers/clients with http mock (see tests)

