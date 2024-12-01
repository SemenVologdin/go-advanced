package services

import (
	"hash/fnv"
	"strconv"
)

type Hash struct {
}

func newHashService() *Hash {
	return &Hash{}
}

func (service Hash) HashString(text string) string {
	h := fnv.New32a()
	h.Write([]byte(text))
	hash := h.Sum32()
	return strconv.Itoa(int(hash))
}
