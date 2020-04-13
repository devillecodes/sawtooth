package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// All the possible moves.
const (
	Up    string = "up"
	Down  string = "down"
	Left  string = "left"
	Right string = "right"
)

// Coord represents the X and Y coordinates of a single cell on the board.
type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Snake represents a snake on the board, yours or an opponents.
type Snake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int     `json:"health"`
	Body   []Coord `json:"body"`
}

// Board represents the battlefield where many a noble snake have competed.
type Board struct {
	Height int     `json:"height"`
	Width  int     `json:"width"`
	Food   []Coord `json:"food"`
	Snakes []Snake `json:"snakes"`
}

// Game contains the ID of a game being played.
type Game struct {
	ID string `json:"id"`
}

// StartRequest contains all the particulars to start a battle.
type StartRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

// StartResponse contains the physical characteristics of your snake.
type StartResponse struct {
	Color    string `json:"color,omitempty"`
	HeadType string `json:"headType,omitempty"`
	TailType string `json:"tailType,omitempty"`
}

// MoveRequest contains all the details you need to decide where to move next.
type MoveRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

// MoveResponse contains the details the server needs for your next move.
type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}

// EndRequest is the final step in a game. Hopefully you were victorius!
type EndRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

// HandleIndex is the root handler of this Battlesnake client.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sawtooth lives! Tremble mortals and despair!")
}

// HandlePing tells the server your snake is in the saddle and ready for battle.
func HandlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

// HandleStart is called at the start of each game your Battlesnake is playing.
// The StartRequest object contains information about the game that's about to start.
func HandleStart(w http.ResponseWriter, r *http.Request) {
	request := StartRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	response := StartResponse{
		Color:    "#600AAA",
		HeadType: "fang",
		TailType: "sharp",
	}

	fmt.Print("START\n")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HandleMove is called for each turn of each game.
// Valid responses are "up", "down", "left", or "right".
func HandleMove(w http.ResponseWriter, r *http.Request) {
	request := MoveRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	move := NextMove(request)

	response := MoveResponse{
		Move: move,
	}

	fmt.Printf("MOVE: %s\n", response.Move)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HandleEnd is called when a game your Battlesnake was playing has ended.
// It's purely for informational purposes, no response required.
func HandleEnd(w http.ResponseWriter, r *http.Request) {
	request := EndRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Nothing to respond with here
	fmt.Print("END\n")
}

// NextMove determines what our next move should be.
func NextMove(r MoveRequest) (move string) {
	b := r.Board
	s := r.You

	if !AgainstWall(s, b) {
		return TowardNearestWall(s, b)
	}

	return HugWall(s, b)
}

// AgainstWall determines whether a snake's head is against the edge of the
// board.
func AgainstWall(s Snake, b Board) bool {
	hix, wix := b.Height-1, b.Width-1
	head := s.Body[0]
	if head.X == 0 || head.X == wix {
		return true
	}
	if head.Y == 0 || head.Y == hix {
		return true
	}
	return false
}

// TowardNearestWall returns direction towards nearest wall.
// Left and right are favoured over up and down when distances are equal.
func TowardNearestWall(s Snake, b Board) string {
	w, h := b.Width, b.Height
	head := s.Body[0]

	fromLeft := head.X + 1
	fromRight := w - head.X
	fromTop := head.Y + 1
	fromBotton := h - head.Y

	switch {
	case fromRight < fromLeft && fromRight <= fromTop && fromRight <= fromBotton:
		return Right
	case fromTop < fromLeft && fromTop < fromBotton && fromTop < fromRight:
		return Up
	case fromBotton < fromLeft && fromBotton < fromTop && fromBotton < fromRight:
		return Down
	default:
		return Left
	}
}

// HugWall ensures we keep hugging the wall in a clockwise direction.
func HugWall(s Snake, b Board) string {
	wix, hix := b.Width-1, b.Height-1
	head := s.Body[0]

	switch {
	case head.Y == 0 && head.X < wix:
		return Right
	case head.X == wix && head.Y < hix:
		return Down
	case head.Y == hix && head.X > 0 && head.X <= wix:
		return Left
	case head.X == 0 && head.Y > 0:
		fallthrough
	default:
		return Up
	}
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/ping", HandlePing)

	http.HandleFunc("/start", HandleStart)
	http.HandleFunc("/move", HandleMove)
	http.HandleFunc("/end", HandleEnd)

	fmt.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
