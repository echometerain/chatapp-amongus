package main

import (
	"io/ioutil"
	"strings"
)

func mapLoad() [][]rune {
	skeld, err := ioutil.ReadFile("skeld map")
	if err != nil {
		panic(err)
	}
	stringArr := strings.Split(string(skeld), "\n")
	var runeArr [][]rune
	for i, v := range stringArr {
		runeArr[i] = []rune(v)
	}
	return runeArr
}
