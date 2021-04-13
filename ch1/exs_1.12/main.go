package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		var (
			cycles  float64 = 5
			res     float64 = 0.001
			size    int     = 100
			nframes int     = 64
			delay   int     = 8
		)
		cycles = getFloat(query, "cycles", cycles)
		res = getFloat(query, "res", res)
		size = getInt(query, "size", size)
		nframes = getInt(query, "nframes", nframes)
		delay = getInt(query, "delay", delay)
		anim := lissajous(cycles, res, size, nframes, delay)
		gif.EncodeAll(w, anim)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getInt(query url.Values, key string, def int) int {
	if value := query.Get(key); value != "" {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return def
		}
		return intValue
	}
	return def
}

func getFloat(query url.Values, key string, def float64) float64 {
	if value := query.Get(key); value != "" {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return def
		}
		return floatValue
	}
	return def
}

func lissajous(cycles, res float64, size, nframes, delay int) *gif.GIF {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	return &anim
}
