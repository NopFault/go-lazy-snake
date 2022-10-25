package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Food struct {
	x int32
	y int32
}

type Tail struct {
	x int32
	y int32
}
type Snake struct {
	x     int32
	y     int32
	size  int32
	color rl.Color
	pos   string
	tail  []Tail
}

type Game struct {
	loop  bool
	speed int
}

func View(snake Snake, food Food) {

	// Redraw
	rl.BeginDrawing()

	rl.DrawRectangle(food.x, food.y, snake.size, snake.size, rl.Red)

	rl.DrawRectangle(snake.x, snake.y, snake.size, snake.size, snake.color)
	for _, tail := range snake.tail {
		rl.DrawRectangle(tail.x, tail.y, snake.size, snake.size, snake.color)
	}

	rl.ClearBackground(rl.RayWhite)
	rl.EndDrawing()
}
func round(number int32, scale int32) int32 {
	x := (number / scale) * scale
	y := x + scale

	if number-x > y-number {
		return y
	} else {
		return x
	}
}
func (f *Food) new(snake *Snake) {
	f.x = round(int32(rand.Intn(rl.GetScreenWidth())), snake.size)
	f.y = round(int32(rand.Intn(rl.GetScreenHeight())), snake.size)
}

func (g *Game) over() {
	g.loop = false
	rl.CloseWindow()
}

func (s *Snake) collisions(game *Game, food *Food) {
	if s.x <= 0 {
		game.over()
	}
	if s.x >= int32(rl.GetScreenWidth()) {
		game.over()
	}
	if s.y <= 0 {
		game.over()
	}
	if s.y >= int32(rl.GetScreenHeight()) {
		game.over()
	}

	if food.x == s.x && food.y == s.y {
		food.new(s)
		s.addTail()
	}
}

func (s *Snake) addTail() {

	var lastTailX int32
	var lastTailY int32

	if len(s.tail) > 0 {
		lastTailX = s.tail[len(s.tail)-1].x
		lastTailY = s.tail[len(s.tail)-1].y
	} else {
		lastTailX = s.x
		lastTailY = s.y
	}

	if s.pos == "U" || s.pos == "D" {
		lastTailY += s.size
	} else {
		lastTailX += s.size
	}

	s.tail = append(s.tail, Tail{x: lastTailX, y: lastTailY})
}

func (s *Snake) relocate() {

	// New coords of tail
	var tmpx int32
	var tmpy int32
	var oldx int32
	var oldy int32
	for i := range s.tail {

		if i > 0 {
			tmpx = s.tail[i].x
			tmpy = s.tail[i].y

			s.tail[i].y = oldy
			s.tail[i].x = oldx

			oldx = tmpx
			oldy = tmpy
		} else {
			oldx = s.tail[i].x
			oldy = s.tail[i].y
			s.tail[i].y = s.y
			s.tail[i].x = s.x
		}
	}

	// New coords of head
	if s.pos == "L" {
		s.x -= s.size
	} else if s.pos == "R" {
		s.x += s.size
	} else if s.pos == "U" {
		s.y -= s.size
	} else if s.pos == "D" {
		s.y += s.size
	}
}

func main() {
	var game Game = Game{
		loop: true,
	}
	rl.InitWindow(800, 450, "Lazy Snake")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		var snake Snake = Snake{
			x:     100,
			y:     200,
			size:  10,
			color: rl.DarkBlue,
			pos:   "R",
			tail:  []Tail{},
		}
		var food Food
		food.new(&snake)

		for game.loop {

			switch rl.GetKeyPressed() {
			case rl.KeyEscape:
				game.over()
				rl.CloseWindow()
			case rl.KeyRight:
				if snake.pos != "L" {
					snake.pos = "R"
				}
			case rl.KeyLeft:
				if snake.pos != "R" {
					snake.pos = "L"
				}
			case rl.KeyDown:
				if snake.pos != "U" {
					snake.pos = "D"

				}
			case rl.KeyUp:
				if snake.pos != "D" {

					snake.pos = "U"
				}
			}

			snake.relocate()
			snake.collisions(&game, &food)

			View(snake, food)
			time.Sleep(100 * time.Millisecond)
		}
	}

	rl.CloseWindow()
}
