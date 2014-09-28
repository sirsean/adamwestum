package service

import (
	"io"
	"bufio"
	"math/rand"
)

var values = []string{
	"",
}

func Load(f io.Reader) []string {
	scanner := bufio.NewScanner(f)
	values = make([]string, 0)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	return values
}

func Lines(number int) []string {
	numValues := len(values)
	indexes := rand.Perm(numValues)
	lines := make([]string, number)
	for i := 0; i < number; i++ {
		if i % len(values) == 0 {
			indexes = rand.Perm(numValues)
		}
		lines[i] = values[indexes[i % numValues]]
	}
	return lines
}
