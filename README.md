# Gache

[![GoDoc](https://godoc.org/github.com/efritz/gache?status.svg)](https://godoc.org/github.com/efritz/gache)
[![Build Status](https://secure.travis-ci.org/efritz/gache.png)](http://travis-ci.org/efritz/gache)
[![Maintainability](https://api.codeclimate.com/v1/badges/1b4448cb4f5672631beb/maintainability)](https://codeclimate.com/github/efritz/gache/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/1b4448cb4f5672631beb/test_coverage)](https://codeclimate.com/github/efritz/gache/test_coverage)

A simple cache interface for Go applications.

## Overview

A cache instance can be configured to store items in memory or in Redis.
Use the `NewMemoryCache` or `NewRedisCache` to create an instance of the
cache.

```go
// Store a value in the cache with the tag `products`
if err := cache.SetValue("product-123", "...", "products"); err != nil {
    // ...
}

// Retrieve the value by key from the cache
val, err := cache.GetValue("product-123")
if err != nil {
    // ...
}

// Process the value if it exists
if val != "" {
    // ...
}

// Remove all keys with the tag `products` after a change to the
// data which backs that set of cache keys.
if err := cache.BustTags("products"); err != nil {
    // ...
}
```

## License

Copyright (c) 2018 Eric Fritz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
