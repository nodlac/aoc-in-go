package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type deer struct {
	name       string
	speed      int
	flightTime int
	restTime   int
	resting    bool
	actionTime int
	distance   int
	points     int
}

func main() {

	filename := "./input-user.txt"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	racers := make([]deer, 0, 10)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		parts := strings.Split(line, " ")
		name := parts[0]
		speed, sErr := strconv.Atoi(parts[3])
		if sErr != nil {
			panic(sErr)
		}
		flightTime, fErr := strconv.Atoi(parts[6])
		if fErr != nil {
			panic(sErr)
		}
		restTime, rErr := strconv.Atoi(parts[13])
		if rErr != nil {
			panic(sErr)
		}
		temp := deer{name: name, speed: speed, flightTime: flightTime, restTime: restTime}
		racers = append(racers, temp)
	}

	furthest := 0

	for _, d := range racers {
		dist := 0
		raceTime := 2503
		for raceTime > 0 {
			if raceTime > d.flightTime {
				dist += d.flightTime * d.speed
				raceTime -= d.flightTime
			} else {
				dist += raceTime * d.speed
				raceTime = 0
			}
			raceTime -= d.restTime
		}
		if dist > furthest {
			furthest = dist
		}
	}

	fmt.Println(furthest)

	fmt.Println("new calc")

	raceTime := 2503
	for i := 0; i < raceTime; i++ {
		var leader []string
		leaderDist := 0
		for i := range racers {
			d := &racers[i]

			d.actionTime += 1
			if !d.resting {
				d.distance += d.speed
				if d.actionTime == d.flightTime {
					d.resting = !d.resting
					d.actionTime = 0
				}
			} else {
				if d.actionTime == d.restTime {
					d.resting = !d.resting
					d.actionTime = 0
				}
			}
			if d.distance > leaderDist {
				leaderDist = d.distance
				leader = []string{d.name}
			} else if d.distance == leaderDist {
				leader = append(leader, d.name)
			}
		}
		for i := range racers {
			d := &racers[i]
			if slices.Contains(leader, d.name) {
				d.points += 1
			}
		}
	}

	mostPoints := 0
	for _, d := range racers {
		if d.points > mostPoints {
			mostPoints = d.points
		}
		fmt.Println(d)
		fmt.Printf("%v has %d points\n", d.name, d.points)
	}
	fmt.Println(mostPoints)
}
