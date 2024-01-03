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
	demoSnapshot := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day22\\demo_input.txt"))
	snapshot := util.ScannerToStringSlice(*util.ReadFile("C:\\Users\\janss\\git\\adventofcode2023\\day22\\input.txt"))
	defer util.TimeTrack(time.Now(), "main")

	jengaTowerDemo := parseJengaTower(demoSnapshot)
	jengaTower := parseJengaTower(snapshot)

	log.Default().Printf("P1 demo: %d", part1(jengaTowerDemo))
	log.Default().Printf("P1: %d", part1(jengaTower))
	log.Default().Printf("P2 demo: %d", part2(jengaTowerDemo))
	log.Default().Printf("P2: %d", part2(jengaTower))
}

type Brick struct {
	id           int
	x, y, z      int
	orientation  string
	size         int
	supports     []int
	supported_by []int
}

func NewBrick(id int, brick string) Brick {
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
		return Brick{id, tx, ty, tz, "x", 1, []int{}, []int{}}
	} else if fx != tx {
		return Brick{id, tx, ty, tz, "x", tx - fx + 1, []int{}, []int{}}
	} else if fy != ty {
		return Brick{id, tx, ty, tz, "y", ty - fy + 1, []int{}, []int{}}
	} else { // fz != tz
		return Brick{id, tx, ty, tz, "z", tz - fz + 1, []int{}, []int{}}
	}
}

func parseJengaTower(snapshot []string) []Brick {
	jengaTower := make([]Brick, 0)
	for id, brick := range snapshot {
		jengaTower = append(jengaTower, NewBrick(id, brick))
	}
	slices.SortFunc(jengaTower, sortBricks)
	return settleTower(jengaTower)
}

func sortBricks(a, b Brick) int {
	if a.x == b.x && a.y == b.y && a.z == b.z {
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
				if jengaTower[j].z < getLowZ(brick)-1 {
					// z diff should be max 1
					break
				}

				if isCollision(Brick{brick.id, brick.x, brick.y, brick.z - 1, brick.orientation, brick.size, brick.supports, brick.supported_by}, jengaTower[j]) {
					collision = true
					break
				}
			}

			if !collision {
				log.Printf("Sorted %d", i)
				jengaTower[i] = Brick{brick.id, brick.x, brick.y, brick.z - 1, brick.orientation, brick.size, brick.supports, brick.supported_by}
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
	return brick.z - brick.size + 1
}

type Point struct {
	x, y, z int
}

func part1(settledTower []Brick) int {
	for i := len(settledTower) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			originalIBrick := settledTower[i]
			// move brick i one down to check for collisions
			iBrick := Brick{
				settledTower[i].id,
				settledTower[i].x,
				settledTower[i].y,
				settledTower[i].z - 1,
				settledTower[i].orientation,
				settledTower[i].size,
				settledTower[i].supports,
				settledTower[i].supported_by,
			}

			if isCollision(iBrick, settledTower[j]) {
				// j supports i
				jBrick := settledTower[j]
				jBrick.supports = append(jBrick.supports, originalIBrick.id)
				originalIBrick.supported_by = append(originalIBrick.supported_by, jBrick.id)
				settledTower[j] = jBrick
				settledTower[i] = originalIBrick
			}
		}
	}

	canDisintegrate := 0

	// TODO check with supported and supported_by sets
	// for i := range settledTower {
	// 	onlySupportingBrick := true
	// 	for _, j := range settledTower[i].supports {
	// 		if len(settledTower[j].supported_by) < 2 {
	// 			onlySupportingBrick = false
	// 			break
	// 		}
	// 	}

	// 	if onlySupportingBrick {
	// 		canDisintegrate++
	// 	}
	// }

	for _, brick := range settledTower {
		if len(brick.supports) == 0 {
			// bricks that support nothing can always go
			canDisintegrate++
			continue
		}

		onlySupportingBrick := false
		for _, supporting := range brick.supports {
			if !sliceutils.Some(settledTower, func(value Brick, index int, slice []Brick) bool {
				return slices.Contains(value.supports, supporting) && value.id != brick.id
			}) {
				onlySupportingBrick = true
			}
		}
		if !onlySupportingBrick {
			canDisintegrate++
		}

	}

	// log.Print(settledTower)
	return canDisintegrate
}

func part2(settledTower []Brick) int {
	totalFalling := 0

	for i := range settledTower {
		brick := settledTower[i]

		queue := []int{}
		for _, j := range brick.supports {
			if len(settledTower[j].supported_by) == 1 {
				queue = append(queue, j)
			}
		}

		falling := queue
		falling = append(falling, brick.id)

		for len(queue) > 0 {
			j := queue[0]
			queue = queue[1:]

			for _, k := range settledTower[j].supported_by {
				if !slices.Contains(falling, k) {
					// if everything that supports k is also in the slice of falling bricks
					if sliceutils.Every(settledTower, func(value Brick, index int, slice []Brick) bool {
						if slices.Contains(value.supports, k) && !slices.Contains(falling, value.id) {
							return false
						}
						return true
					}) {
						queue = append(queue, k)
						falling = append(falling, k)
					}
				}
			}

		}
		totalFalling += len(falling) - 1
	}

	// log.Print(settledTower)
	return totalFalling
}
