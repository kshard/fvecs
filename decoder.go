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

// Vector base types
type Base interface {
	float32 | uint32 | byte
}

// Vector decoder
type Decoder[T Base] struct {
	r     io.Reader
	n     int
	codec func(int, []byte) []T
	d     [4]byte
	b     []byte
}

// New instance of decoder
func NewDecoder[T Base](r io.Reader) *Decoder[T] {
	d := &Decoder[T]{r: r}

	switch any(*new(T)).(type) {
	case float32:
		d.n = 4
		d.codec = toFVec
	case uint32:
		d.n = 4
		d.codec = toIVec
	case byte:
		d.n = 1
		d.codec = toBVec
	}

	return d
}

// Read vector from stream
func (d *Decoder[T]) Read() ([]T, error) {
	s, err := d.readSize()
	if err != nil {
		return nil, err
	}

	l := s * d.n
	if len(d.b) != int(l) {
		d.b = make([]byte, l)
	}

	_, err = d.r.Read(d.b)
	if err != nil {
		return nil, err
	}

	return d.codec(s, d.b), nil
}

// read vector size
func (d *Decoder[T]) readSize() (int, error) {
	ps := d.d[:]
	_, err := d.r.Read(ps)
	return int(d.d[0]) | int(d.d[1])<<8 | int(d.d[2])<<16 | int(d.d[3])<<24, err
}

// bytes to .fvecs
func toFVec[T Base](s int, b []byte) []T {
	v := make([]float32, s)

	p := 0
	for i := 0; i < len(b); i += 4 {
		v[p] = math.Float32frombits(binary.LittleEndian.Uint32(b[i : i+4]))
		p++
	}

	return any(v).([]T)
}

// bytes to .ivecs
func toIVec[T Base](s int, b []byte) []T {
	v := make([]uint32, s)

	p := 0
	for i := 0; i < len(b); i += 4 {
		v[p] = binary.LittleEndian.Uint32(b[i : i+4])
		p++
	}

	return any(v).([]T)
}

// bytes to .bvecs
func toBVec[T Base](s int, b []byte) []T {
	v := make([]byte, s)
	copy(v, b)
	return any(v).([]T)
}
