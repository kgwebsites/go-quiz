# Quiz Game
[Project Idea from Gophercises](https://github.com/gophercises/quiz)

Clone this repo and run `$ go run quiz.go` to start the quiz with default options. This will use the problems.csv file for the quiz questions and answers and set the quiz timer to 30 seconds.

To set your own csv file for questions and answers, add in the `-f` flag with the name of the file like so:
```
$ go run quiz.go -f=custom_problems.csv
```

To set a custom timer, use the `-t` flag followed by the number of seconds you wish the timer to be:
```
$ go run quiz.go -t=60
```

To shuffle the questions, just add in the `-s` flag:
```
$ go run quiz.go -s
```