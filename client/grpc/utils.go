package grpc

import (
	"bytes"
	"encoding/binary"
	"time"
)

func encodeUvarint(value uint64) []byte {
	// Get the size required to represent the value.
	size := binary.Size(value)
	if size == 0 {
		return nil
	}

	buf := make([]byte, size)
	n := binary.PutUvarint(buf, value)
	return buf[:n]
}

func encodeTime(t time.Time) []byte {
	var buf bytes.Buffer

	s := t.Unix()
	if s != 0 {
		buf.Write(encodeFieldNumberAndTyp3(1, 0))
		buf.Write(encodeUvarint(uint64(s)))
	}

	ns := int32(t.Nanosecond()) // this int64 -> int32 cast is safe (nanos are in [0, 999999999])
	if ns != 0 {
		buf.Write(encodeFieldNumberAndTyp3(2, 0))
		buf.Write(encodeUvarint(uint64(ns)))
	}

	return buf.Bytes()
}

func encodeFieldNumberAndTyp3(num uint32, typ uint8) []byte {
	return encodeUvarint((uint64(num) << 3) | uint64(typ))
}
