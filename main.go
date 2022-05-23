package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// modifying code originally from:
// • https://ebiten.org/tour/image.html
// • https://ebiten.org/examples/tiles.html
// • https://github.com/chonlatee/spaceship
const (
	screenWidth  = 240
	screenHeight = 240

	// Background spritesheet
	tileSize = 16
	tileXNum = 25 // the number of 16px columns in the image width
)

var (
	tilesImage *ebiten.Image
)

type Game struct {
	keys       []ebiten.Key
	bgLayers   [][]int
	player     *ebiten.Image
	playerPosX float64
	playerPosY float64
}

// var g &Game
var g *Game

func init() {
	eImg, _, err := ebitenutil.NewImageFromFile("images/tiles--mailbox.png")
	if err != nil {
		log.Fatal(err)
	}
	g = &Game{
		player:     eImg,
		playerPosX: 0,
		playerPosY: 0,
		bgLayers: [][]int{
			{
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 244, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 244, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 219, 243, 243, 243, 219, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 244, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			},
			{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 26, 27, 28, 29, 30, 31, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 51, 52, 53, 54, 55, 56, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 76, 77, 78, 79, 80, 81, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 101, 102, 103, 104, 105, 106, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 126, 127, 128, 129, 130, 131, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 303, 303, 245, 242, 303, 303, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0,
			},
		},
	}
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	img, _, err := image.Decode(bytes.NewReader(images.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	for _, k := range g.keys {
		// TODO [bug]
		// Don't allow double speed when arrow + wasd keys are pressed at the same
		// time (in same direction); just accept the last used key. (chris)
		if k == ebiten.KeyRight || k == ebiten.KeyD {
			g.playerPosX += 3
		} else if k == ebiten.KeyLeft || k == ebiten.KeyA {
			g.playerPosX -= 3
		} else if k == ebiten.KeyUp || k == ebiten.KeyW {
			g.playerPosY -= 3
		} else if k == ebiten.KeyDown || k == ebiten.KeyS {
			g.playerPosY += 3
		}
	}
	// Create map boundaries (screen edges)
	// inspired by https://github.com/chonlatee/spaceship
	w, h := g.player.Size()
	// left edge
	if g.playerPosX <= 0 {
		g.playerPosX = 0
	}
	// right edge
	if g.playerPosX+float64(w) >= float64(screenWidth) {
		g.playerPosX = float64(screenWidth) - float64(w)
	}
	// top edge
	if g.playerPosY <= 0 {
		g.playerPosY = 0
	}
	// bottom edge
	if g.playerPosY+float64(h) >= float64(screenHeight) {
		g.playerPosY = float64(screenHeight) - float64(h)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Background image:
	// Draw each tile with each DrawImage call.
	// As the source images of all DrawImage calls are always same,
	// this rendering is done very efficiently.
	// For more detail, see https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Image.DrawImage
	const xNum = screenWidth / tileSize // 15
	for _, l := range g.bgLayers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

			sx := (t % tileXNum) * tileSize
			sy := (t / tileXNum) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}

	// Character image:
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.playerPosX), float64(g.playerPosY))
	screen.DrawImage(g.player, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// TODO: why x2?
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("character collision")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
