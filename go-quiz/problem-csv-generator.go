package main

import (
	"encoding/csv"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	questions [][]string
	operation = map[string]string{"a": "+", "s": "-", "m": "x", "d": "/"}
	num_limit = 10
)

func generateRandomQuestion(topic string) []string {
	var (
		question []string
		output   int
	)

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	op1 := r.Intn(num_limit)
	op2 := r.Intn(num_limit)

	if op1 < op2 {
		tmp := op1
		op1 = op2
		op2 = tmp
	}

	switch topic {
	case "a":
		output = op1 + op2
	case "s":
		output = op1 - op2
	case "m":
		output = op1 * op2
	case "d":
		output = op1 / op2
		output = int(math.Round(float64(output)))
	}

	q := strconv.Itoa(op1) + operation[topic] + strconv.Itoa(op2)
	return append(question, q, strconv.Itoa(output))
}

func GenerateCsv(topic string, total_questions int) {
	f, err := os.Create(topic + string(".csv"))
	Check(err)
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for i := 0; i < total_questions; i++ {
		w.Write(generateRandomQuestion(topic))
	}
}
