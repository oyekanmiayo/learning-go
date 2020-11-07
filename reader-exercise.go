package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr *MyReader) Read(b []byte) (int, error) {
	count := 0
	for i,_ := range b {
		if b[i] != byte('A'){
			b[i] = byte('A')
			count++
		}
	}
	
	return count, nil
}

func main() {
	reader.Validate(&MyReader{})
}