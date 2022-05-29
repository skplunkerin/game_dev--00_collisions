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
	"github.com/jakecoffman/cp"
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
	keys     []ebiten.Key
	bgLayers [][]int
	// NOTE: this tutorial references meta data from a tilemap to determine what
	// tiles should cause a collision, that sounds like a better option than me
	// manually tracking tile ID's:
	//   - tutorial: http://chipmunk-physics.net/tutorials/ChipmunkTileDemo/
	//               https://www.raywenderlich.com/2779-collisions-and-collectables-how-to-make-a-tile-based-game-with-cocos2d-2-x-part-2#toc-anchor-001
	// bgCollisionTileIds []int
	// bgCollisions [][]int
	bgCollisions []int
	player       *ebiten.Image
	playerPosX   float64
	playerPosY   float64
	space        *cp.Space
}

var g *Game

func init() {
	eImg, _, err := ebitenutil.NewImageFromFile("images/tiles--mailbox.png")
	if err != nil {
		log.Fatal(err)
	}
	// Chipmunk creates a list of all possible collisions that it will need to
	// iterate over for each check; each check makes the physics/collisions more
	// accurate, but each iteration takes CPU; this will need to be balanced.
	// https://chipmunk-physics.net/release/ChipmunkLatest-Docs/#cpSpace-Iterations
	space := cp.NewSpace()
	space.Iterations = 1 // Default: 10

	// Don't think this applies to use, even though the tiles are the same size.
	// https://chipmunk-physics.net/release/ChipmunkLatest-Docs/#cpSpace-SpatialHash
	// // The space will contain a very large number of similarly sized objects.
	// // This is the perfect candidate for using the spatial hash.
	// // Generally you will never need to do this.
	// space.UseSpatialHash(2.0, 10000)

	g = &Game{
		space:      space,
		player:     eImg,
		playerPosX: 0,
		playerPosY: 0,
		// collisions idea:
		// 1. []int{}: tilemap id's for collide-able tiles
		//    - pro: very simple
		//    - con: don't have an easy way to identify the tile id's;
		//           (fortunately, this example is easy enough to grab from below)
		// 2. [][]int{}: tilemap of the position that is collide-able
		// bgCollisionTileIds: []int{
		// 	26, 27, 28, 29, 30, 31,
		// 	51, 52, 53, 54, 55, 56,
		// 	76, 77, 78, 79, 80, 81,
		// 	101, 102, 103, 104, 105, 106,
		// 	126, 127, 128, 129, 130, 131,
		// 	303,
		// },
		// bgCollisions: [][]int{
		bgCollisions: []int{
			// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,

			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
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
	g.space.Step(1.0 / float64(ebiten.MaxTPS()))
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

	g.checkCollisions()

	return nil
}

// checkCollisions will make sure the player's X,Y positions will collide
// against the screen boundaries (window edges), and certain image tiles that
// are setup to be a boundary.
// TODO: setup logic to signify which image tiles have a collision boundary.
//       (chris)
func (g *Game) checkCollisions() {
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

	// map tiles collisions should happen from Chipmunk2D (cp)
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
			// if layer == 1 {
			// 	if t != 0 {
			// 		fmt.Println("layer: ", layer)
			// 		fmt.Println("l: ", l)
			// 		fmt.Println("i: ", i)
			// 		fmt.Println("t: ", t)
			// 		fmt.Printf("bgCollisions[t]: %#v\n", g.bgCollisions[t])
			// 		fmt.Printf("bgCollisions[i]: %#v\n", g.bgCollisions[i])
			// 		fmt.Println()
			// 	}
			// }
			if g.bgCollisions[i] == 1 {
				// add collision for tile
				// https://www.reddit.com/r/ebiten/comments/mghl4k/using_the_go_port_of_chipmunk2d_in_a_tile_based/
				// NOTE: check 1st comment, it would be better to create a single object
				// with all collisions than several objects with their own collisions if I
				// can figure this out.
				// (it looks like autogeometry capability was added to the repo: https://github.com/jakecoffman/cp/pull/20)
				// add blocks to physics space
				body := cp.NewStaticBody()
				body.SetPosition(cp.Vector{X: float64(sx), Y: float64(sy)})
				shape := cp.NewBox(body, tileSize, tileSize, 0)
				shape.SetElasticity(0)
				shape.SetFriction(1)
				g.space.AddBody(shape.Body())
				g.space.AddShape(shape)
			}

			// add tile image to screen
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}

	// Character image:
	op := &ebiten.DrawImageOptions{}
	// TODO: is `playerPosX|Y` the point in the center of the avatar? I think it
	// is? (chris)
	op.GeoM.Translate(float64(g.playerPosX), float64(g.playerPosY))
	// // body := cp.NewStaticBody()
	// body := cp.NewBody()
	// body.SetPosition(cp.Vector{X: float64(sx), Y: float64(sy)})
	// shape := cp.NewBox(body, tileSize, tileSize, 0)
	// shape.SetElasticity(0)
	// shape.SetFriction(1)
	// g.space.AddBody(shape.Body())
	// g.space.AddShape(shape)
	// add collision to character
	body := cp.NewBody(1.0, cp.INFINITY)
	// is this needed? (chris)
	// body.SetVelocity()
	shape := cp.NewCircle(body, tileSize/2, cp.Vector{})
	shape.SetFriction(0)
	shape.SetElasticity(0)
	shape.SetCollisionType(1)
	body.SetPosition(cp.Vector{X: float64(g.playerPosX), Y: float64(g.playerPosY)})
	// add character to screen
	screen.DrawImage(g.player, op)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf(
			"TPS: %0.2f\n"+
				"PlayerX: %f\n"+
				"PlayerY: %f",
			ebiten.CurrentTPS(),
			g.playerPosX,
			g.playerPosY,
		),
	)
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
