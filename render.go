package main

import (
	"math"
	"math/rand"
	"sort"
	"sync"

	"github.com/thebargaintenor/raytracer-go/engine"
)

// RenderContext describes props of entire render job
type RenderContext struct {
	Width      int
	Height     int
	SampleSize int
	Camera     *engine.Camera
	Scene      *engine.Scene
}

// Bitmap describes dimensions and color data for an image
type Bitmap struct {
	Width  int
	Height int
	Data   *[]engine.Color
}

// PixelRenderContext contains context for concurrent rendering of individual pixel
type PixelRenderContext struct {
	Index   int
	X       int
	Y       int
	Context *RenderContext
}

// RenderedPixel describes output for finished pixel render
type RenderedPixel struct {
	Index int
	Color engine.Color
}

func worker(id int, wg *sync.WaitGroup, jobs <-chan *PixelRenderContext, results chan<- *RenderedPixel) {
	for job := range jobs {
		results <- &RenderedPixel{
			Index: job.Index,
			Color: getPixel(job.Context, job.X, job.Y),
		}

		wg.Done()
	}
}

func renderParallel(context *RenderContext) *Bitmap {
	pixelCount := context.Width * context.Height
	jobs := make(chan *PixelRenderContext, pixelCount)
	results := make(chan *RenderedPixel, pixelCount)

	var wg sync.WaitGroup
	wg.Add(pixelCount)

	for w := 1; w <= 8; w++ {
		go worker(w, &wg, jobs, results)
	}

	idx := 0 // this should be independent because the scanlines are inverted for PPM
	for y := context.Height - 1; y >= 0; y-- {
		for x := 0; x < context.Width; x++ {
			jobs <- &PixelRenderContext{
				Index:   idx,
				X:       x,
				Y:       y,
				Context: context,
			}
			idx++
		}
	}
	close(jobs)

	wg.Wait()
	close(results)

	sortableResults := []*RenderedPixel{}
	for result := range results {
		sortableResults = append(sortableResults, result)
	}

	sort.Slice(sortableResults, func(i, j int) bool {
		return sortableResults[i].Index < sortableResults[j].Index
	})

	data := []engine.Color{}
	for _, result := range sortableResults {
		data = append(data, result.Color)
	}

	return &Bitmap{
		Width:  context.Width,
		Height: context.Height,
		Data:   &data,
	}
}

func renderSingleThread(context *RenderContext) *Bitmap {
	data := []engine.Color{}

	for y := context.Height - 1; y >= 0; y-- {
		for x := 0; x < context.Width; x++ {
			data = append(data, getPixel(context, x, y))
		}
	}

	return &Bitmap{
		Width:  context.Width,
		Height: context.Height,
		Data:   &data,
	}
}

func getPixel(context *RenderContext, x, y int) engine.Color {
	c := engine.Color{0.0, 0.0, 0.0}

	for s := 0; s < context.SampleSize; s++ {
		u := (float64(x) + rand.Float64()) / float64(context.Width)
		v := (float64(y) + rand.Float64()) / float64(context.Height)
		ray := context.Camera.GetRay(u, v)

		c = c.Add(color(&ray, context.Scene, 0))
	}

	return correctGamma(c, context.SampleSize)
}

func correctGamma(c engine.Color, sampleSize int) engine.Color {
	c = c.ScalarDiv(float64(sampleSize))
	return engine.Color{
		math.Sqrt(c[0]),
		math.Sqrt(c[1]),
		math.Sqrt(c[2]),
	}
}

func color(r *engine.Ray, scene *engine.Scene, depth uint8) engine.Color {
	if hit, success := scene.Hit(r, 0.001, math.MaxFloat64); success {
		attenuation, scatteredRay, scattered := hit.Material.Scatter(r, hit)
		if depth < 50 && scattered {
			return attenuation.Mult(color(scatteredRay, scene, depth+1))
		}

		return engine.Color{0.0, 0.0, 0.0}
	}

	// scene background
	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y() + 1.0)
	startColor := engine.Color{1.0, 1.0, 1.0}
	endColor := engine.Color{0.5, 0.7, 1.0}
	// linear interpolation being color stops
	return startColor.ScalarMult(1.0 - t).Add(endColor.ScalarMult(t))
}
