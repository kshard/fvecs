//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/fvecs
//

package fvecs_test

import (
	"testing"

	"github.com/fogfish/it/v2"
	"github.com/kshard/fvecs"
)

func TestDecodeBVecs(t *testing.T) {
	for i := 0; i < 10; i++ {
		v, err := dbvecs.Read()
		it.Then(t).Should(
			it.Nil(err),
			it.Seq(v).Equal(1, 2, 3, 4, 5, 6, 7, 8, 9),
		)
	}
}

func TestDecodeIVecs(t *testing.T) {
	for i := 0; i < 10; i++ {
		v, err := divecs.Read()
		it.Then(t).Should(
			it.Nil(err),
			it.Seq(v).Equal(1, 2, 3, 4, 5, 6, 7, 8, 9),
		)
	}
}

func TestDecodeFVecs(t *testing.T) {
	for i := 0; i < 10; i++ {
		v, err := dfvecs.Read()
		it.Then(t).Should(
			it.Nil(err),
			it.Less(v[0]-1.1, 1e-5),
			it.Less(v[1]-2.2, 1e-5),
			it.Less(v[2]-3.3, 1e-5),
			it.Less(v[3]-4.4, 1e-5),
			it.Less(v[4]-5.5, 1e-5),
			it.Less(v[5]-6.6, 1e-5),
			it.Less(v[6]-7.7, 1e-5),
			it.Less(v[7]-8.8, 1e-5),
			it.Less(v[8]-9.9, 1e-5),
		)
	}
}

var (
	dbvecs = fvecs.NewDecoder[byte](mock{tbvecs})
	divecs = fvecs.NewDecoder[uint32](mock{tivecs})
	dfvecs = fvecs.NewDecoder[float32](mock{tfvecs})
)

func BenchmarkBVecs(b *testing.B) {
	b.ReportAllocs()

	for n := b.N; n > 0; n-- {
		if _, err := dbvecs.Read(); err != nil {
			panic(err)
		}
	}
}

func BenchmarkIVecs(b *testing.B) {
	b.ReportAllocs()

	for n := b.N; n > 0; n-- {
		if _, err := divecs.Read(); err != nil {
			panic(err)
		}
	}
}

func BenchmarkFVecs(b *testing.B) {
	b.ReportAllocs()

	for n := b.N; n > 0; n-- {
		if _, err := dfvecs.Read(); err != nil {
			panic(err)
		}
	}
}

var (
	tbvecs = []byte{
		9, 0, 0, 0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}

	tivecs = []byte{
		9, 0, 0, 0,
		1, 0, 0, 0,
		2, 0, 0, 0,
		3, 0, 0, 0,
		4, 0, 0, 0,
		5, 0, 0, 0,
		6, 0, 0, 0,
		7, 0, 0, 0,
		8, 0, 0, 0,
		9, 0, 0, 0,
	}

	tfvecs = []byte{
		9, 0, 0, 0,
		0xcd, 0xcc, 0x8c, 0x3f, // 1.1
		0xcd, 0xcc, 0x0c, 0x40, // 2.2
		0x33, 0x33, 0x53, 0x40, // 3.3
		0xcd, 0xcc, 0x8c, 0x40, // 4.4
		0x00, 0x00, 0xb0, 0x40, // 5.5
		0x33, 0x33, 0xd3, 0x40, // 6.6
		0x66, 0x66, 0xf6, 0x40, // 7.7
		0xcd, 0xcc, 0x0c, 0x41, // 8.8
		0x66, 0x66, 0x1e, 0x41, // 9.9
	}
)

// io.Reader mock
type mock struct {
	b []byte
}

func (m mock) Read(p []byte) (int, error) {
	if len(p) == 4 {
		copy(p, m.b[0:4])
		return 4, nil
	}

	copy(p, m.b[4:])
	return len(m.b), nil
}
