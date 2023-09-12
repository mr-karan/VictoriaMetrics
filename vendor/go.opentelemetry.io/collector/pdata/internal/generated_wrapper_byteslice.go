// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package internal

type ByteSlice struct {
	orig *[]byte
}

func GetOrigByteSlice(ms ByteSlice) *[]byte {
	return ms.orig
}

func NewByteSlice(orig *[]byte) ByteSlice {
	return ByteSlice{orig: orig}
}
