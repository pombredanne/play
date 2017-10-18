// Play with rendering eX0 player with CanvasRenderingContext2D API.
package main

import (
	"math"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func main() {
	document.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
		go paint()
	})
}

func paint() {
	var canvas = document.GetElementByID("canvas").(*dom.HTMLCanvasElement)
	canvas.Width, canvas.Height = 320, 320

	var ctx = canvas.GetContext2d()
	ctx.SetTransform(10, 0, 0, 10, 160, 160)
	ctx.Rotate(-1.2)

	ctx.FillStyle = "red"
	ctx.FillRect(2, -1, 11, 2)

	ctx.StrokeStyle = "red"
	ctx.LineWidth = 2
	ctx.BeginPath()
	ctx.Ellipse(0, 0, 7, 7, Tau*1/12, 0, Tau*10/12, false)
	ctx.Stroke()
}

// Tau is the constant τ, which equals to 6.283185... or 2π.
// Reference: https://oeis.org/A019692
const Tau = 2 * math.Pi