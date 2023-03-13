package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string) ([]problem, error) {

	//read all the problems from the quiz.csv
	if fObj, err := os.Open(fileName); err == nil {
		csvR := csv.NewReader(fObj)
		if cLines, err := csvR.ReadAll(); err == nil {
			return parseProblem(cLines), nil
		} else {
			return nil, fmt.Errorf("error in reading data in csv"+"format from %s file; %s", fileName, err.Error())
		}
	} else {
		return nil, fmt.Errorf("error in opening %s file; %s", fileName, err.Error())
	}
}

func main() {
	//1.input the name of the file
	//2.set the duration of the timer
	//3. pull the problems from the file (calling our problem puller func)
	//4.handle the error
	//5.create a variable to count our correct answer
	//6.using the duration of the timer, we want to initialize the timer
	//7.loop throught the problem, print the questions, we'll accept the answer
	//8.we'll calculate and print out the result

	fName := flag.String("f", "quiz.csv", "path of csv file")

	timer := flag.Int("t", 30, "timer for the quiz")

	flag.Parse()
	problems, err := problemPuller(*fName)
	if err != nil {
		exit(fmt.Sprintf("Something went wrong:%s", err.Error()))
	}
	correctAns := 0
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)

	ansC := make(chan string)
problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("Problem %d: %s=", i+1, p.q)
		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <-ansC:
			if iAns == p.a {
				correctAns++
			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}
	}

	fmt.Printf("Your result is %d out of %d\n", correctAns, len(problems))
	fmt.Printf("Press enter to exit")
	<-ansC
}

func parseProblem(lines [][]string) []problem {
	// go over the lines and parse them,problem struct
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
