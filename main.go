// Copyright 2017 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build example jsgo

package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten"
)

var count int

var (
	emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
)

func init() {
	emptyImage.Fill(color.White)
}

func main() {
	if err := ebiten.Run(update, 640, 480, 1, "Hello, World! holy shhh"); err != nil {
		log.Fatal(err)
	}
}

var (
	radius   float64 = (640 + 480) / 2
	rotSpeed         = 0.00275
	armLen   float64 = (640 + 480) / 10
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	count++
	count %= 360

	cf := float64(count)

	x0 := float32(radius/2 + armLen*math.Cos(2*math.Pi*cf*rotSpeed) - armLen*math.Sin(2*math.Pi*cf*rotSpeed))
	y0 := float32(radius/2 + armLen*math.Sin(2*math.Pi*cf*rotSpeed) + armLen*math.Cos(2*math.Pi*cf*rotSpeed))

	v, i := line(x0, y0, float32(radius/2), float32(radius/2), color.RGBA{0xf2, 0x12, 0x12, 0x80})
	screen.DrawTriangles(v, i, emptyImage, nil)

	ebitenutil.DebugPrint(screen, fmt.Sprint("x0:", x0, "\ny0:", y0, "\ncount:", count))
	return nil
}

func line(x0, y0, x1, y1 float32, clr color.RGBA) ([]ebiten.Vertex, []uint16) {
	const width = 1

	theta := math.Atan2(float64(y1-y0), float64(x1-x0))
	theta += math.Pi / 2
	dx := float32(math.Cos(theta))
	dy := float32(math.Sin(theta))

	r := float32(clr.R) / 0xff
	g := float32(clr.G) / 0xff
	b := float32(clr.B) / 0xff
	a := float32(clr.A) / 0xff

	return []ebiten.Vertex{
		{
			DstX:   x0 - width*dx/2,
			DstY:   y0 - width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x0 + width*dx/2,
			DstY:   y0 + width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1 - width*dx/2,
			DstY:   y1 - width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x1 + width*dx/2,
			DstY:   y1 + width*dy/2,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
	}, []uint16{0, 1, 2, 1, 2, 3}
}
