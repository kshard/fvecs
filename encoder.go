//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/fvecs
//

package fvecs

import (
	"encoding/binary"
	"io"
	"math"
)

// Vector decoder
type Encoder[T Base] struct {
	w     io.Writer
	n     int
	codec func([]T, []byte)
	b     []byte
}

// New instance of decoder
func NewEncoder[T Base](w io.Writer) *Encoder[T] {
	e := &Encoder[T]{w: w}

	switch any(*new(T)).(type) {
	case float32:
		e.n = 4
		e.codec = fromFVec
	case uint32:
		e.n = 4
		e.codec = fromIVec
	case byte:
		e.n = 1
		e.codec = fromBVec
	}

	return e
}

// Read vector from stream
func (e *Encoder[T]) Write(v []T) error {
	l := len(v)*e.n + 4
	if len(e.b) != int(l) {
		e.b = make([]byte, l)
	}

	binary.LittleEndian.PutUint32(e.b[0:4], uint32(len(v)))
	e.codec(v, e.b[4:])

	_, err := e.w.Write(e.b)
	if err != nil {
		return err
	}

	return nil
}

// .fvecs to bytes to
func fromFVec[T Base](v []T, b []byte) {
	vv := any(v).([]float32)

	p := 0
	for i := 0; i < len(v); i++ {
		u := math.Float32bits(vv[i])
		binary.LittleEndian.PutUint32(b[p:p+4], u)

		p += 4
	}

}

// .ivecs to bytes
func fromIVec[T Base](v []T, b []byte) {
	vv := any(v).([]uint32)

	p := 0
	for i := 0; i < len(v); i++ {
		binary.LittleEndian.PutUint32(b[p:p+4], vv[i])

		p += 4
	}

}

// .bvecs to bytes
func fromBVec[T Base](v []T, b []byte) {
	vv := any(v).([]byte)
	copy(b, vv)
}
