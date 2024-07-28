package utils

import (
	"bytes"
	"encoding/binary"
)

func GenericMarshalBinary[T any](obj T) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, obj)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
