package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckFirst(t *testing.T) {
	err := checkFist([]rune("2dsfdsf"))
	require.NotEqual(t, err, nil)
}

func TestCheckFirst2(t *testing.T) {
	err := checkFist([]rune("sfd2dsf"))
	require.Equal(t, err, nil)
}

func TestDeleteFromSlice(t *testing.T) {
	newSlice := deleteFromSlice([]rune{'a', 'b', 'c', 'd'}, 2)
	require.Equal(t, newSlice, []rune{'a', 'b', 'd'})
}

func TestAddSymbols(t *testing.T) {
	require.Equal(t, addSymbols('a', 5), []rune{'a', 'a', 'a', 'a', 'a'})
}

func TestUnpack(t *testing.T) {
	require.Equal(t, unpack("a2b2cc"), "aabbcc")
}
