package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestDesignMain(t *testing.T) {
	// DesignMain()

	f, err := os.Open("./design/test.obj")
	if err != nil {
		panic(err)
	}
	defer exec(f.Close)

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}

func TestGob(t *testing.T) {
	v := struct {
		N []int
		L int
	}{[]int{1, 2, 3}, 3}

	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(v)
	if err != nil {
		panic(err)
	}
	v.N[2] = 10
	decoder := gob.NewDecoder(buf)
	err = decoder.Decode(&v)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
