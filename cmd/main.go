package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func getSubject() ([]string, int) {
	n := []string{"quack", "quack"}
	return n, len(n)
}

func getNouns() ([]string, int) {
	n := []string{"quack", "quuack", "quaack"}
	return n, len(n)
}

func getVerbs() ([]string, int) {
	n := []string{"qua", "ququ", "quak"}
	return n, len(n)
}

func getAdv() ([]string, int) {
	n := []string{"quack", "quack", "quackquack", "qwack"}
	return n, len(n)
}

func getAdj() ([]string, int) {
	n := []string{"quack", "quacquacquac", "quack"}
	return n, len(n)
}

func getConj() ([]string, int) {
	n := []string{"qua", "ack", "quack"}
	return n, len(n)
}

func getArt() ([]string, int) {
	n := []string{"qua", "ack"}
	return n, len(n)
}

func getEndPunc() ([]string, int) {
	n := []string{".", "?", "!", ";", ":"}
	return n, len(n)
}

func getSplitPunc() ([]string, int) {
	n := []string{",", ";", ":"}
	return n, len(n)
}

func coinFlip(r *rand.Rand) bool {
	return r.Intn(1) != 0
}

func makeSentence(r *rand.Rand) string {
	s, sl := getSubject()
	n, nl := getNouns()
	v, nv := getVerbs()
	av, avl := getAdv()
	ad, adl := getAdj()
	c, cl := getConj()
	ar, arl := getArt()
	ep, epl := getEndPunc()
	sp, spl := getSplitPunc()

	sLen := r.Intn(20)
	if sLen == 0 {
		sLen = r.Intn(1) + 1
	}

	for v := range sLen {
		if sLen < 2 && sLen > 1 {
		}
	}

	// sentence = n[r.Intn(nl)] + ' '

	// nouns[num] + ' ' + verbs[num] + ' ' + adv[num] + ' ' + adj[num]
}

func main() {
	prompt := flag.String("prompt", "", "Question or prompt to ask the rubber ducky")
	flag.Parse()

	apiKey := os.Getenv("YOUR_AI_KEY")
	if apiKey == "" {
		panic("Can't quack, no API Key")
	}

	seed, err := strconv.Atoi(apiKey)
	if err != nil {
		panic("Can't quack :<")
	}

	r := rand.New(rand.NewSource(int64(seed)))

	fmt.Printf("%s", *prompt)
}
