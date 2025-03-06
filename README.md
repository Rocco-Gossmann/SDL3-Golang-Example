# SDL3 - GoLang - Example Project

this is a small "Hello World" to checkout 
SDL3-Purego.

A Version of SDL3, that lets Go load the Runtime-Binarys directly.
Instead of having to build them via CGO first.

the `main.go` is an example application. 
(Just a 320x200px black window, with "Hello" written into it)

## how to make it work
- I've put all runtimes I could test into this project as well.
because the SDL3 ones are hard to find otherwise.
(Source: https://www.lwjgl.org/browse/nightly)

- Copy the runtime for your Build-Target next to your build binary (or into your go modules root folder)
- Then remove the `amd64.` or `arm64.` Prefix. from it.

- as for the code, have a look at the `main.go` for a starting point.

## dependencies
The Project should just have 2 depencies
- github.com/ebitengine/purego  (indirect)
    will load the SDL3-Runtime Lib (`.dll`, `.so.0` or  `.dylib`)

- github.com/jupiterrider/purego-sdl3
    Provides all the SDL3-Bindings to Go

A BIG thank you to https://github.com/JupiterRider (PureGo-SDL3)  and https://github.com/ebitengine (PureGo)
I can't wait to play around with this.
