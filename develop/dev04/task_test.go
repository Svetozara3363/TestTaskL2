package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToLower(t *testing.T) {
	var strings []string
	strings = append(strings, "ABCD")
	strings = append(strings, "dfEr")
	require.Equal(t, toLower(strings), []string{"abcd", "dfer"})
}

func TestDeleteRepeat(t *testing.T) {
	var strings []string
	strings = append(strings, "one")
	strings = append(strings, "two")
	strings = append(strings, "one")
	require.Equal(t, deleteRepeat(strings), []string{"one", "two"})
}
