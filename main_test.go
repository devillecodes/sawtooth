package main

import (
	"fmt"
	"testing"
)

func TestAgainstWall(t *testing.T) {
	b3x3 := Board{3, 3, []Coord{}, []Snake{}}
	b5x5 := Board{5, 5, []Coord{}, []Snake{}}

	var tests = []struct {
		s    Snake
		b    Board
		want bool
	}{
		{Snake{Body: []Coord{{0, 0}}}, b3x3, true},
		{Snake{Body: []Coord{{1, 0}}}, b3x3, true},
		{Snake{Body: []Coord{{2, 0}}}, b3x3, true},
		{Snake{Body: []Coord{{0, 1}}}, b3x3, true},
		{Snake{Body: []Coord{{1, 1}}}, b3x3, false},
		{Snake{Body: []Coord{{2, 1}}}, b3x3, true},
		{Snake{Body: []Coord{{0, 2}}}, b3x3, true},
		{Snake{Body: []Coord{{1, 2}}}, b3x3, true},
		{Snake{Body: []Coord{{2, 2}}}, b3x3, true},

		{Snake{Body: []Coord{{0, 0}}}, b5x5, true},
		{Snake{Body: []Coord{{1, 1}}}, b5x5, false},
		{Snake{Body: []Coord{{2, 2}}}, b5x5, false},
		{Snake{Body: []Coord{{3, 3}}}, b5x5, false},
		{Snake{Body: []Coord{{4, 4}}}, b5x5, true},
		{Snake{Body: []Coord{{0, 4}}}, b5x5, true},
		{Snake{Body: []Coord{{4, 0}}}, b5x5, true},
		{Snake{Body: []Coord{{3, 1}}}, b5x5, false},
		{Snake{Body: []Coord{{1, 3}}}, b5x5, false},
	}

	for _, tt := range tests {
		Body := tt.s.Body[0]
		testname := fmt.Sprintf("[%v,%v],[%v,%v]", Body.X, Body.Y, tt.b.Width, tt.b.Height)
		t.Run(testname, func(t *testing.T) {
			result := AgainstWall(tt.s, tt.b)
			if result != tt.want {
				t.Errorf("got %v, want %v", result, tt.want)
			}
		})
	}
}

func TestTowardNearestWall(t *testing.T) {
	b7x7 := Board{7, 7, []Coord{}, []Snake{}}

	var tests = []struct {
		s    Snake
		b    Board
		want string
	}{
		{Snake{Body: []Coord{{1, 1}}}, b7x7, Left},
		{Snake{Body: []Coord{{1, 2}}}, b7x7, Left},
		{Snake{Body: []Coord{{1, 3}}}, b7x7, Left},
		{Snake{Body: []Coord{{1, 4}}}, b7x7, Left},
		{Snake{Body: []Coord{{1, 5}}}, b7x7, Left},

		{Snake{Body: []Coord{{2, 1}}}, b7x7, Up},
		{Snake{Body: []Coord{{2, 2}}}, b7x7, Left},
		{Snake{Body: []Coord{{2, 3}}}, b7x7, Left},
		{Snake{Body: []Coord{{2, 4}}}, b7x7, Left},
		{Snake{Body: []Coord{{2, 5}}}, b7x7, Down},

		{Snake{Body: []Coord{{3, 1}}}, b7x7, Up},
		{Snake{Body: []Coord{{3, 2}}}, b7x7, Up},
		{Snake{Body: []Coord{{3, 3}}}, b7x7, Left},
		{Snake{Body: []Coord{{3, 4}}}, b7x7, Down},
		{Snake{Body: []Coord{{3, 5}}}, b7x7, Down},

		{Snake{Body: []Coord{{4, 1}}}, b7x7, Up},
		{Snake{Body: []Coord{{4, 2}}}, b7x7, Right},
		{Snake{Body: []Coord{{4, 3}}}, b7x7, Right},
		{Snake{Body: []Coord{{4, 4}}}, b7x7, Right},
		{Snake{Body: []Coord{{4, 5}}}, b7x7, Down},

		{Snake{Body: []Coord{{5, 1}}}, b7x7, Right},
		{Snake{Body: []Coord{{5, 2}}}, b7x7, Right},
		{Snake{Body: []Coord{{5, 3}}}, b7x7, Right},
		{Snake{Body: []Coord{{5, 4}}}, b7x7, Right},
		{Snake{Body: []Coord{{5, 5}}}, b7x7, Right},
	}

	for _, tt := range tests {
		Body := tt.s.Body[0]
		testname := fmt.Sprintf("[%v,%v],[%v,%v]", Body.X, Body.Y, tt.b.Width, tt.b.Height)
		t.Run(testname, func(t *testing.T) {
			result := TowardNearestWall(tt.s, tt.b)
			if result != tt.want {
				t.Errorf("got %v, want %v", result, tt.want)
			}
		})
	}
}
