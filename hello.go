package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
   
)

func main() {
	if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }
	
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	count, _ := lines(file)
	fmt.Printf("Lines count: %d \n", count)
    
}

func lines(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
}