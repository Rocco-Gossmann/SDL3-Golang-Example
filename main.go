package main

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
)

func main() {

	if !sdl.SetHint(sdl.HintRenderVSync, "1") {
		panic(sdl.GetError())
	}
	defer sdl.Quit()

	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	if window = sdl.CreateWindow("Hello Go-SDL3", 320, 200, sdl.WindowResizable); window == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)

	var renderer *sdl.Renderer
	if renderer = sdl.CreateRenderer(window, ""); renderer == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyRenderer(renderer)

Outer:
	for {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				break Outer
			case sdl.EventKeyDown:
				if event.Key().Scancode == sdl.ScancodeEscape {
					break Outer
				}
			}
		}

		sdl.SetRenderDrawColor(renderer, 0, 0, 0, 255)
		sdl.RenderClear(renderer)

		sdl.SetRenderDrawColor(renderer, 255, 255, 255, 255)
		sdl.RenderDebugText(renderer, 8, 8, "Hello")

		sdl.RenderPresent(renderer)
	}

}
