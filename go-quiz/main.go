package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func Check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseProblems(questions [][]string) []problem {
	p := make([]problem, len(questions))
	for i := range p {
		p[i].q = questions[i][0]
		p[i].a = questions[i][1]
	}
	return p
}

func main() {
	var (
		correct = 0
		file    = flag.String("csv", "sample.csv", "CSV file containing quiz questions in 'question,answer' format.")
		limit   = flag.Int("limit", 30, "Time to solve each question in seconds.")
		qno     = flag.Int("qno", 10, "Numbers of questions in the quiz.")
		topic   = flag.String("topic", "mix",
			`Type of operation to be quizzed.
Options:
	a -> Addition
	s -> Substraction
	m -> Multiplication
	d -> Division
`)
	)
	flag.Parse()

	if *topic != "mix" {
		GenerateCsv(*topic, *qno)
		*file = *topic + ".csv"
	}

	in, err := ioutil.ReadFile(*file)
	Check(err)

	q := csv.NewReader(strings.NewReader(string(in)))
	questions, err := q.ReadAll()
	Check(err)

	for _, p := range parseProblems(questions) {
		resp := make(chan string)
		fmt.Printf("Solve: %s\n", p.q)
		timer := time.NewTimer(time.Duration(*limit) * time.Second)

		go func() {
			ans := ""
			fmt.Scanln(&ans)
			resp <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou Scored %d out of %d\n", correct, len(questions))
			return

		case response := <-resp:
			if response == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You Scored %d out of %d\n", correct, len(questions))
}
