package blocks

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	numRows       = 22
	numCols       = 10
	overflow      = 2
	tickerMillis  = 17
	levelSecs     = 30
	speedBump     = 1.1
	virtualWidth  = 640
	virtualHeight = 480
	cell          = 20
	fastSpeed     = 16
	slowSpeed     = 16 * 50
)

// Game ...
type Game struct {
	board      *Arr
	pixels     []*ebiten.Image
	curRow     int
	curCol     int
	curPiece   *Arr
	gameOver   bool
	score      int
	level      int
	levelSpeed int
	levelStart time.Time
	curSpeed   int
	elapsed    int
	lastTs     time.Time
	face       font.Face
}

// New ..
func New() *Game {
	rand.Seed(time.Now().UnixNano())
	palette := [...]color.Color{
		color.Black,
		color.RGBA{R: 0, G: 255, B: 255, A: 255}, // cyan
		color.RGBA{R: 255, G: 255, B: 0, A: 255}, // yellow
		color.RGBA{R: 128, G: 0, B: 128, A: 255}, // purple
		color.RGBA{R: 255, G: 165, B: 0, A: 255}, // orange
		color.RGBA{R: 0, G: 0, B: 128, A: 255},   // blue
		color.RGBA{R: 0, G: 128, B: 0, A: 255},   // green
		color.RGBA{R: 128, G: 0, B: 0, A: 255},   // red
	}
	pixels := make([]*ebiten.Image, len(palette))
	for i, col := range palette {
		pixels[i] = ebiten.NewImage(1, 1)
		pixels[i].Fill(col)
	}
	ttf, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	face, _ := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	g := &Game{
		pixels:     pixels,
		board:      NewArr(numRows, numCols),
		face:       face,
		lastTs:     time.Now(),
		curSpeed:   slowSpeed,
		levelSpeed: slowSpeed,
		levelStart: time.Now(),
		level:      1,
	}
	g.spawn()
	return g
}

// Update ...
func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.left()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.right()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.up()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.space()
	}
	ts := time.Now()
	if ts.Sub(g.levelStart).Seconds() > levelSecs {
		g.level++
		g.levelStart = ts
		g.levelSpeed = int(float64(g.levelSpeed) / speedBump)
		g.curSpeed = g.levelSpeed
	}
	delta := ts.Sub(g.lastTs).Milliseconds()
	g.lastTs = ts
	g.elapsed += int(delta)
	if g.elapsed >= g.curSpeed {
		g.elapsed = 0
		g.down()
	}
	return nil
}

// Draw ...
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 80, G: 80, B: 80, A: 255})
	sx := (virtualWidth - cell*numCols) / 2
	sy := (virtualHeight - cell*(numRows-overflow)) / 2
	for r := overflow; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			op := ebiten.DrawImageOptions{}
			op.GeoM.Scale(cell, cell)
			op.GeoM.Translate(float64(sx+c*cell), float64(sy+(r-overflow)*cell))
			val := g.board.get(r, c)
			screen.DrawImage(g.pixels[val], &op)
		}
	}
	text.Draw(screen, fmt.Sprintf("Level: %d", g.level), g.face, 10, 24, color.White)
	text.Draw(screen, fmt.Sprintf("Score: %d", g.score), g.face, 10, 48, color.White)
	if g.gameOver {
		text.Draw(screen, "Game Over", g.face, 300, 200, color.White)
	}
}

// Layout ...
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return virtualWidth, virtualHeight
}

func (g *Game) spawn() {
	which := rand.Intn(len(pieces))
	g.curPiece = pieces[which].Clone()
	g.curCol = 4
	g.curRow = 0
	if !g.board.CanPlace(g.curPiece, g.curRow, g.curCol) {
		g.gameOver = true
	}
}

func (g *Game) down() {
	clone := g.board.Clone()
	clone.Remove(g.curPiece, g.curRow, g.curCol)
	if clone.CanPlace(g.curPiece, g.curRow+1, g.curCol) {
		g.curRow++
		clone.Place(g.curPiece, g.curRow, g.curCol)
		g.board = clone
	} else {
		scoreMultipliers := [...]int{0, 40, 100, 300, 1200}
		removed := g.board.RemoveFullRows()
		g.score += scoreMultipliers[removed] * g.level
		g.spawn()
		g.curSpeed = g.levelSpeed
	}
}

func (g *Game) left() {
	clone := g.board.Clone()
	clone.Remove(g.curPiece, g.curRow, g.curCol)
	if clone.CanPlace(g.curPiece, g.curRow, g.curCol-1) {
		g.curCol--
		clone.Place(g.curPiece, g.curRow, g.curCol)
		g.board = clone
	}
}

func (g *Game) right() {
	clone := g.board.Clone()
	clone.Remove(g.curPiece, g.curRow, g.curCol)
	if clone.CanPlace(g.curPiece, g.curRow, g.curCol+1) {
		g.curCol++
		clone.Place(g.curPiece, g.curRow, g.curCol)
		g.board = clone
	}
}

func (g *Game) up() {
	clone := g.board.Clone()
	clone.Remove(g.curPiece, g.curRow, g.curCol)
	rotated := g.curPiece.RotateCounterClockwise()
	if clone.CanPlace(rotated, g.curRow, g.curCol) {
		clone.Place(rotated, g.curRow, g.curCol)
		g.curPiece = rotated
		g.board = clone
	}
}

func (g *Game) space() {
	g.curSpeed = fastSpeed
}
