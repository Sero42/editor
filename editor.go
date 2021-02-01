// Author: S. Rösch
// Date: 17.08.2020 + 14.09.2020 + 30.01.2021
// Purpose: Using the ADT Sequence implementing an basic editor !!

// TODO: Add other than alphanumerical symbols

package main

import (
	"fmt" 
	. "./seq"
	"math"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)


func show (s Sequence) string {
	var savedIndex uint = s.CurrentIndex()
	var res string
	for i:=uint(0); i<s.Length(); i++ {
		s.Position(i)
		e,_:= s.CurrentElement()
		res = res + fmt.Sprintf("%c", e)
	}
	s.Position(savedIndex)
	return res
}


func show_modified (s Sequence, n uint) string {
	var savedIndex uint = s.CurrentIndex()
	var res string
	for i:=uint(0); i<n; i++ {
		s.Position(i)
		e,_:= s.CurrentElement()
		res = res + fmt.Sprintf("%c", e)
	}
	s.Position(savedIndex)
	return res
}

func draw_text (s, t Sequence, starting_column, starting_row, 
		utfEightification uint, err error, font *ttf.Font, text, 
		surface *sdl.Surface, window *sdl.Window) {
	surface.FillRect(nil, 0xffffffff)
	if err = text.Blit(nil, surface, &sdl.Rect{X: 0, Y: 0, W: 0, H:
			0}); err != nil {
		return
	}
	if s.CurrentIndex() > 60 {
		starting_column = s.CurrentIndex() - 60
	}else{
		starting_column = 0
	}
	utfEightification = uint(len(
	show_modified(s,starting_column))) // 
	if t.CurrentIndex() > 16 {
		starting_row = t.CurrentIndex() - 16
	}else{
		starting_row = 0
	}		
	savedIndex := t.CurrentIndex()
	for i:=uint(0); i<t.Length(); i++  {
		t.Position(i)
		g,_ := t.CurrentElement()
		if float64(t.CurrentIndex()) - float64(starting_row) >= 0 && 
				t.CurrentIndex() - starting_row < 17 && g.(Sequence).
				Length() > starting_column {text, _ = font.RenderUTF8Solid(
				show(g.(Sequence))[utfEightification:] + " ", sdl.Color{0,
				0,0,0})				 
			text.Blit (nil, surface, &sdl.Rect{X: 0, Y: 16 + int32(
				math.Abs(float64(t.CurrentIndex()) - float64(
				starting_row))*32), W: 0, H: 0})
		}
	}
	t.Position(savedIndex)	
	text, _ = font.RenderUTF8Solid("Column:         ", sdl.Color{0,0,0,0})		
	text.Blit (nil, surface, &sdl.Rect{X: 0, Y: 18*32, W: 0, H: 0}) 
	text, _ = font.RenderUTF8Solid(fmt.Sprint(s.CurrentIndex()+1), 
		sdl.Color{0,0,0,0})
	text.Blit (nil, surface, &sdl.Rect{X: 105, Y: 18*32, W:0, H:0})
	text, _ = font.RenderUTF8Solid("Row:         ", sdl.Color{0,0,0,0})
	text.Blit (nil, surface, &sdl.Rect{X: 190, Y: 18*32, W: 0, H: 0})
	text, _ = font.RenderUTF8Solid(fmt.Sprint(t.CurrentIndex() + 1),
		sdl.Color{0,0,0,0})
	text.Blit (nil, surface, &sdl.Rect{X: 250, Y: 18*32, W: 0, H: 0})
	text,_ = font.RenderUTF8Solid("_", sdl.Color{0,0,0,0})
	text.Blit(nil, surface, &sdl.Rect{X: int32((s.CurrentIndex()-
		starting_column)*14), Y: 16 + int32(math.Abs(float64(
		t.CurrentIndex()) - float64(starting_row))*32), W: 0, H: 0})

	window.UpdateSurface()
}


func main () {
	var s Sequence = New (rune('a')) //new Sequence
	var t Sequence = New (s) // new Sequence of Sequences
	var starting_column uint
	var utfEightification uint
	var starting_row uint
	t.Insert(s)
	t.Back()
	
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("THE EDITOR!!", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, 900, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	surface, err := window.GetSurface()
		if err != nil {
			panic(err)
		}
	defer surface.Free()
	
	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()
	font, err := ttf.OpenFont("./LiberationMono-Regular.ttf",24)
	if err != nil {
		panic(err)
	}
	text, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	defer text.Free()
	
	e,_:= t.CurrentElement()
	s = e.(Sequence)
	draw_text (s, t, starting_column, starting_row, utfEightification,
		err, font, text, surface, window)

	running := true
A:	for running {
		e,_:= t.CurrentElement()
		s = e.(Sequence)
		draw_text (s, t, starting_column, starting_row, utfEightification,
			err, font, text, surface, window)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch v := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				keyCode := v.Keysym.Sym
				depth := v.Keysym.Mod				
				if v.State == sdl.PRESSED {
					fmt.Println(v.Keysym.Sym)
					switch v.Keysym.Sym {	// Symbol key switch
						case sdl.K_ESCAPE:
							println("Quit")
							running = false
							break A
						case sdl.K_LEFT:  // Right Arrow
							s.Back ()				
						case sdl.K_RIGHT:  // Left Arrow
							s.Forth ()
						case sdl.K_UP: // Up
							t.Back()
						case sdl.K_DOWN: // Down
							t.Forth()	
						case sdl.K_DELETE: // Del-Key
							s.Delete ()
						case sdl.K_END: // End-Key
							s.Position(s.Length())
						case sdl.K_HOME: // POS1-Key
							s.Position(0)
						case sdl.K_BACKSPACE: // Backspace-Key
							s.Back()
							s.Delete()
							s.Forth()
						case sdl.K_RETURN: // Return-Key
							t.Forth()
							t.Insert(New(rune(0)))
							t.Back()
						case sdl.K_EXCLAIM:
							s.Insert('!')
						case sdl.K_QUESTION:
							s.Insert(rune('?'))
						case sdl.K_LEFTPAREN:
							fmt.Println("Waht?=")
							
					// Other than alphanumerical Symbols: TODO					
						default:
							if keyCode < 10000 {
								if depth == 4097 || depth == 4098 || 
								   depth == 12288 ||depth == 1 || 
								   depth == 2 || depth == 8192 {
									   keyCode -= 32
								}
								s.Insert(rune(keyCode))
							}
					}
				}							
			}
		}
	sdl.Delay(100)	
	}
}
