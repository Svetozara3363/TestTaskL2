package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	slice := []string{"a", "b", "c"}
	reverseSlice(slice)
	require.Equal(t, slice, []string{"c", "b", "a"})
}

func TestDeleteSimilar(t *testing.T) {
	require.Equal(t, deleteSimilar([]string{"c", "b", "a", "b"}), []string{"c", "b", "a"})
}

func TestDeleteSimilar2(t *testing.T) {
	require.Equal(t, deleteSimilar([]string{"c", "b", "a"}), []string{"c", "b", "a"})
}

func TestSortByColumn(t *testing.T) {
	var lines []string
	lines = append(lines, "aaaa ggggg")
	lines = append(lines, "cccc aaaaa")
	lines = append(lines, "aaaa bbbbb")
	lines = append(lines, "hhhh ccccc")

	var res []string
	res = append(res, "cccc aaaaa")
	res = append(res, "aaaa bbbbb")
	res = append(res, "hhhh ccccc")
	res = append(res, "aaaa ggggg")

	require.Equal(t, sortByColumn(lines, 1, false), res)
}

func TestSortByColumnNumber(t *testing.T) {
	var lines []string
	lines = append(lines, "aaaa 33")
	lines = append(lines, "cccc 111")
	lines = append(lines, "aaaa 2")
	lines = append(lines, "hhhh 32")

	var res []string
	res = append(res, "aaaa 2")
	res = append(res, "hhhh 32")
	res = append(res, "aaaa 33")
	res = append(res, "cccc 111")

	require.Equal(t, sortByColumn(lines, 1, true), res)
}
