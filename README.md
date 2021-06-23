# go-quiz

## Installation

- `git clone git@github.com:shashikarsiddharth/go-quiz.git`
- `cd go-quiz`
- `go build .`
- `./go-quiz` to play simple quiz

## Usage

```        
Usage of ./go-quiz:
  -csv string
    	CSV file containing quiz questions in 'question,answer' format. (default "sample.csv")
  -limit int
    	Time to solve each question in seconds. (default 30)
  -qno int
    	Numbers of questions in the quiz. (default 10)
  -topic string
    	Type of operation to be quizzed.
    	Options:
    		a -> Addition
    		s -> Substraction
    		m -> Multiplication
    		d -> Division
    	 (default "mix")
```