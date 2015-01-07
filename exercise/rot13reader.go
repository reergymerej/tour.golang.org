package main

import (
        "io"
        "os"
        "strings"
        "fmt"
)

type rot13Reader struct {
        r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
        // read from inner reader
        
        fmt.Println(r.r)
        return len(b), err
}

func main() {
        s := strings.NewReader("Lbh penpxrq gur pbqr!")
        
        b := make([]byte, 8)
        bytesRead, err := s.Read(b)
        if err == nil {
                fmt.Println(bytesRead)  
        }
        r := rot13Reader{s}
        io.Copy(os.Stdout, &r)
}
