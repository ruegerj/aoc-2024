package day12

import (
	"fmt"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day12 struct{}

func (d Day12) Part1(input string) *common.Solution {
	tiles := parseTiles(input)
	gardens := make([]*garden, 0)
	visited := make(map[*gardenTile]bool)

	for _, row := range tiles {
		for _, tile := range row {
			if _, checked := visited[tile]; checked {
				continue
			}
			nextTiles := []*gardenTile{tile}
			garden := walkGarden(nextTiles, &garden{symbol: tile.symbol}, visited, tiles)
			gardens = append(gardens, garden)
		}
	}

	totalPrice := 0

	for _, garden := range gardens {
		totalPrice += garden.area * garden.perimeter
	}

	return common.NewSolution(1, totalPrice)
}

func (d Day12) Part2(input string) *common.Solution {
	tiles := parseTiles(input)
	gardens := make([]*garden, 0)
	visited := make(map[*gardenTile]bool)

	for _, row := range tiles {
		for _, tile := range row {
			if _, checked := visited[tile]; checked {
				continue
			}
			nextTiles := []*gardenTile{tile}
			garden := walkGarden(nextTiles, &garden{symbol: tile.symbol}, visited, tiles)
			gardens = append(gardens, garden)
		}
	}

	fmt.Println(countCorners(tiles[3][2], tiles))

	totalPrice := 0

	for _, garden := range gardens {
		totalPrice += garden.area * garden.corners
	}

	return common.NewSolution(2, totalPrice)
}

type garden struct {
	symbol                   string
	area, perimeter, corners int
}

type gardenTile struct {
	x, y   int
	symbol string
}

func walkGarden(nextTiles []*gardenTile, garden *garden, visited map[*gardenTile]bool, grid [][]*gardenTile) *garden {
	current := popHead(nextTiles)

	if current == nil {
		return garden
	}

	nextTiles = nextTiles[1:]

	if _, alreadyVisited := visited[current]; alreadyVisited {
		return walkGarden(nextTiles, garden, visited, grid)
	}

	visited[current] = true
	garden.area += 1
	garden.perimeter += 4
	garden.corners += countCorners(current, grid)

	if current.y-1 >= 0 && grid[current.y-1][current.x].symbol == current.symbol {
		nextTiles = append(nextTiles, grid[current.y-1][current.x])
		garden.perimeter--
	}
	if current.y+1 < len(grid) && grid[current.y+1][current.x].symbol == current.symbol {
		nextTiles = append(nextTiles, grid[current.y+1][current.x])
		garden.perimeter--
	}
	if current.x-1 >= 0 && grid[current.y][current.x-1].symbol == current.symbol {
		nextTiles = append(nextTiles, grid[current.y][current.x-1])
		garden.perimeter--
	}
	if current.x+1 < len(grid[0]) && grid[current.y][current.x+1].symbol == current.symbol {
		nextTiles = append(nextTiles, grid[current.y][current.x+1])
		garden.perimeter--
	}

	return walkGarden(nextTiles, garden, visited, grid)
}

func countCorners(tile *gardenTile, grid [][]*gardenTile) int {
	cornerCount := 0

	top := getTile(tile.x, tile.y-1, grid)
	topRight := getTile(tile.x+1, tile.y-1, grid)
	right := getTile(tile.x+1, tile.y, grid)
	bottomRight := getTile(tile.x+1, tile.y+1, grid)
	bottom := getTile(tile.x, tile.y+1, grid)
	bottomLeft := getTile(tile.x-1, tile.y+1, grid)
	left := getTile(tile.x-1, tile.y, grid)
	topLeft := getTile(tile.x-1, tile.y-1, grid)

	// top right (outwards)
	if (top == nil || top.symbol != tile.symbol) && (right == nil || right.symbol != tile.symbol) {
		cornerCount++
	}
	// bottom right (outwards)
	if (bottom == nil || bottom.symbol != tile.symbol) && (right == nil || right.symbol != tile.symbol) {
		cornerCount++
	}
	// bottom left (outwards)
	if (bottom == nil || bottom.symbol != tile.symbol) && (left == nil || left.symbol != tile.symbol) {
		cornerCount++
	}
	// top left (outwards)
	if (top == nil || top.symbol != tile.symbol) && (left == nil || left.symbol != tile.symbol) {
		cornerCount++
	}

	// top right (inwards)
	if (top != nil && top.symbol == tile.symbol) && (topRight == nil || topRight.symbol != tile.symbol) && (right != nil && right.symbol == tile.symbol) {
		cornerCount++
	}
	// bottom right (inwards)
	if (bottom != nil && bottom.symbol == tile.symbol) && (bottomRight == nil || bottomRight.symbol != tile.symbol) && (right != nil && right.symbol == tile.symbol) {
		cornerCount++
	}
	// bottom left (inwards)
	if (bottom != nil && bottom.symbol == tile.symbol) && (bottomLeft == nil || bottomLeft.symbol != tile.symbol) && (left != nil && left.symbol == tile.symbol) {
		cornerCount++
	}
	// top left (inwards)
	if (top != nil && top.symbol == tile.symbol) && (topLeft == nil || topLeft.symbol != tile.symbol) && (left != nil && left.symbol == tile.symbol) {
		cornerCount++
	}

	return cornerCount
}

func popHead(slice []*gardenTile) *gardenTile {
	if len(slice) == 0 {
		return nil
	}
	return slice[0]
}

func getTile(x, y int, grid [][]*gardenTile) *gardenTile {
	if x >= len(grid[0]) || x < 0 || y >= len(grid) || y < 0 {
		return nil
	}
	return grid[y][x]
}

func parseTiles(input string) [][]*gardenTile {
	grid := util.Matrix(input, "")
	tiles := make([][]*gardenTile, len(grid))
	for y := 0; y < len(grid); y++ {
		tiles[y] = make([]*gardenTile, len(grid[0]))
		for x := 0; x < len(grid[0]); x++ {
			tiles[y][x] = &gardenTile{
				x:      x,
				y:      y,
				symbol: grid[y][x],
			}
		}
	}
	return tiles
}
