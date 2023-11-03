package utils

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Draw() {
	cs := g.CellSize

	if g.CellSize >= g.InitCellSize {
		rl.ClearBackground(color.RGBA{20, 20, 20, 255})
	} else {
		rl.ClearBackground(color.RGBA{0, 0, 0, 255})
	}

	for x := range g.Grid.Cells {
		for y := range g.Grid.Cells[x] {
			c := g.GetCell(x, y)

			c.Color = color.RGBA{0, 0, 0, 255}
			if c.Alive {
				c.Color = color.RGBA{uint8(float64(g.GetNumberAliveNeighbors(c)) / 4 * 255), 127, 0, 255}
			}
			if rl.GetMouseX() >= int32(x*cs) && rl.GetMouseX() <= int32(x*cs+cs-1) && rl.GetMouseY() >= int32(y*cs) && rl.GetMouseY() <= int32(y*cs+cs-1) {
				if rl.IsMouseButtonDown(rl.MouseLeftButton) {
					if !c.JustChanged {
						c.Alive = !g.GetCell(x, y).Alive
						c.JustChanged = true
						c.Color = color.RGBA{0, 0, 0, 255}
						if c.Alive {
							c.Color = color.RGBA{uint8(float64(g.GetNumberAliveNeighbors(c)) / 4 * 255), 127, 0, 255}
						}
					}
				} else {
					c.JustChanged = false
				}
				c.Color = color.RGBA{255, 255, 255, 255}
			}
			if g.CellSize >= g.InitCellSize {
				rl.DrawRectangle(int32(x*cs+1), int32(y*cs+1), int32(cs-2), int32(cs-2), c.Color)
				continue
			}
			if c.Alive {
				rl.DrawRectangle(int32(x*cs), int32(y*cs), int32(cs), int32(cs), c.Color)
			}
		}
	}
}
