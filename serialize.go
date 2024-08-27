package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
)

func Serialize(writer io.Writer, obj ...any) error {
	for _, i := range obj {
		err := binary.Write(writer, binary.LittleEndian, i)
		if err != nil {
			return err
		}
	}
	return nil
}

func SerializeMain() {
	buf := new(bytes.Buffer)
	err := Serialize(buf, int32(5), false, []byte("Hello"), int32(123), float64(5.6))
	if err != nil {
		panic(err)
	}
	for _, b := range buf.Bytes() {
		fmt.Printf("%X ", b)
	}

	v1 := *get[int32](buf)
	v2 := *get[bool](buf)
	v3 := *get[[]byte](buf)

	fmt.Println(v1, v2, v3)
}

func get[T any](buf io.Reader) *T {
	v := new(T)
	err := binary.Read(buf, binary.LittleEndian, v)
	if err != nil {
		panic(err)
	}
	return v
}

func gobEncode(writer io.Writer, val ...any) error {
	encoder := gob.NewEncoder(writer)
	for _, v := range val {
		err := encoder.Encode(v)
		if err != nil {
			return err
		}
	}
	return nil
}
