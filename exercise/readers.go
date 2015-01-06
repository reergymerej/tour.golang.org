package main

import "code.google.com/p/go-tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
        for i, _ := range b {
                // same effect as b[i] = 65
                b[i] = 'A'
        }
        return len(b), err
}

func main() {
        reader.Validate(MyReader{})
}
