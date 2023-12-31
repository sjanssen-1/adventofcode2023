package main

import (
	"adventofcode2023/util"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/Goldziher/go-utils/sliceutils"
)

// https://adventofcode.com/2023/day/22
func main() {
	snapshot := util.ScannerToStringSlice(*util.ReadFile("/Users/stefjanssens/git/adventofcode2023/day22/demo_input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	jengaTower := parseJengaTower(snapshot)

	log.Default().Printf("P1: %d", part1(jengaTower))
	log.Default().Printf("P2: %d", part2())
}

type Brick struct {
	x, y, z     int
	orientation string
	size        int
}

func NewBrick(brick string) Brick {
	split := strings.Split(brick, "~")

	from, to := strings.Split(split[0], ","), strings.Split(split[1], ",")
	fx, _ := strconv.Atoi(from[0])
	fy, _ := strconv.Atoi(from[1])
	fz, _ := strconv.Atoi(from[2])

	tx, _ := strconv.Atoi(to[0])
	ty, _ := strconv.Atoi(to[1])
	tz, _ := strconv.Atoi(to[2])

	if split[0] == split[1] {
		// single block brick
		return Brick{tx, ty, tz, "x", 1}
	} else if fx != tx {
		return Brick{tx, ty, tz, "x", tx - fx + 1}
	} else if fy != ty {
		return Brick{tx, ty, tz, "y", ty - fy + 1}
	} else { // fz != tz
		return Brick{tx, ty, tz, "z", tz - fz + 1}
	}
}

func parseJengaTower(snapshot []string) []Brick {
	jengaTower := make([]Brick, 0)
	for _, brick := range snapshot {
		jengaTower = append(jengaTower, NewBrick(brick))
	}
	slices.SortFunc(jengaTower, sortBricks)
	return jengaTower
}

func sortBricks(a, b Brick) int {
	if a == b {
		return 0
	} else if a.z < b.z {
		return -1
	} else {
		return 1
	}
}

func settleTower(jengaTower []Brick) []Brick {
	settled := false

loop:
	for !settled {
		for i, brick := range jengaTower {
			if getLowZ(brick) == 1 {
				// skip bricks already on ground level
				continue
			}

			collision := false
			for j := i - 1; j >= 0; j-- {
				// if jengaTower[j].z < getLowZ(brick)-1 {
				// 	// z diff should be max 1
				// 	break
				// }

				if isCollision(Brick{brick.x, brick.y, brick.z - 1, brick.orientation, brick.size}, jengaTower[j]) {
					collision = true
					break
				}
			}

			if !collision {
				jengaTower[i] = Brick{brick.x, brick.y, brick.z - 1, brick.orientation, brick.size}
				slices.SortFunc(jengaTower, sortBricks)
				continue loop
			}
		}
		settled = true
	}

	return jengaTower
}

func isCollision(a, b Brick) bool {
	pa := getPoints(a)
	pb := getPoints(b)
	return len(sliceutils.Intersection(pa, pb)) > 0
}

func getPoints(brick Brick) []Point {
	points := []Point{}
	switch brick.orientation {
	case "x":
		for i := 0; i < brick.size; i++ {
			points = append(points, Point{brick.x - i, brick.y, brick.z})
		}
	case "y":
		for i := 0; i < brick.size; i++ {
			points = append(points, Point{brick.x, brick.y - i, brick.z})
		}
	case "z":
		for i := 0; i < brick.size; i++ {
			points = append(points, Point{brick.x, brick.y, brick.z - i})
		}
	}
	return points
}

func getZs(brick Brick) []int {
	if brick.orientation != "z" {
		return []int{brick.z}
	} else {
		zs := []int{}
		for i := 0; i < brick.size; i++ {
			zs = append(zs, brick.z+i)
		}
		return zs
	}
}

func getLowZ(brick Brick) int {
	if brick.orientation != "z" {
		return brick.z
	}
	return brick.z - brick.size
}

type Point struct {
	x, y, z int
}

func part1(jengaTower []Brick) int {
	settledTower := settleTower(jengaTower)

	canDisintegrate := 1 // top block can always be disintegrated
	for i := 0; i < len(settledTower)-1; i++ {
		collisionCounter := 0
		for j := i + 1; j < len(settledTower); j++ {
			if isCollision(Brick{settledTower[j].x, settledTower[j].y, settledTower[j].z - 1, settledTower[j].orientation, settledTower[j].size}, settledTower[i]) {
				collisionCounter++
			}
		}
		if collisionCounter > 1 {
			canDisintegrate++
		}
	}
	// log.Print(settledTower)
	return canDisintegrate
}

func part2() int {
	return 0
}
