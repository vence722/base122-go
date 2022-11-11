# base122-go
Encoder and decoder implementation of [base122](https://blog.kevinalbs.com/base122) encoding scheme in Go programing language.

Base122 is a space efficient binary-to-text encoding scheme, which encodes the binary into UTF-8 strings. It use less space to encode your binary data compared to base64 (12.5% higher efficiency according to the [wiki page](https://en.wikipedia.org/wiki/Binary-to-text_encoding)).

This repo is inspired by [patrickfav/base122-java](https://github.com/patrickfav/base122-java) and [kevinAlbs/Base122](https://github.com/kevinAlbs/Base122), and aims to provide easy-to-use interface for Gophers to use base122 to encode your data efficiently.

## Intro

## Requirement
```
Go >= 1.18
```

## Installation
```
go get github.com/vence722/base122-go
```

## Usage
Simple encode & decode example:
```
import "github.com/vence722/base122-go"

...

// Encode
s := "hello base122!!!"
encText, err := base122.EncodeToString([]byte(s))
if err != nil {
    // error handling
}
fmt.Println(encText) // 4-Fc<@b0׬S	Hd!H 

...

// Decode
txt, err := base122.DecodeFromString(encText)
if err != nil {
    // error handling
}
fmt.Println(txt == s)  // true

```

## Contributors
Vence Lin (vence722@gmail.com)
