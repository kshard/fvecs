//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/fvecs
//

package fvecs_test

import (
	"bytes"
	"testing"

	"github.com/fogfish/it/v2"
	"github.com/kshard/fvecs"
)

func TestEncodeBVecs(t *testing.T) {
	b := &bytes.Buffer{}

	e := fvecs.NewEncoder[byte](b)
	err := e.Write([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})

	it.Then(t).Should(
		it.Nil(err),
		it.Seq(b.Bytes()).Equal(tbvecs...),
	)
}

func TestEncodeIVecs(t *testing.T) {
	b := &bytes.Buffer{}

	e := fvecs.NewEncoder[uint32](b)
	err := e.Write([]uint32{1, 2, 3, 4, 5, 6, 7, 8, 9})

	it.Then(t).Should(
		it.Nil(err),
		it.Seq(b.Bytes()).Equal(tivecs...),
	)
}

func TestEncodeFVecs(t *testing.T) {
	b := &bytes.Buffer{}

	e := fvecs.NewEncoder[float32](b)
	err := e.Write([]float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9})

	it.Then(t).Should(
		it.Nil(err),
		it.Seq(b.Bytes()).Equal(tfvecs...),
	)
}
