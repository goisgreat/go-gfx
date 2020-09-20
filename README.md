# Golang Game Engine
## What is the *Golang Game Engine* and how do I use it?
The golang game engine is a "physics" library for golang. Here is an example to help explain usage.
```golang
package gogfx

import (
    time "time"
    gogfx "github.com/goisgreat/go-gfx"
)

func main() {
    stream := gogfx.EventStream{
        Input: make(chan byte)
    }
    sprite := gogfx.Point{
        gogfx.Position{0,0},
        gogfx.Vector{1,0},
        gogfx.KeyboardController{
            Input: stream.Input,
            KeyboardMap: gogfx.WASD,
        },
    }
    sprite.KeyboardController.OnKeyboardControl = func(control gogfx.KeyboardConrol) {
        println(control)
    }

    go gogfx.GrabInput(stream)
    
    sprite.Init()
    
    for {
        gogfx.ClearTerminal()
        sprite.Update()
        sprite.Draw()
        time.Sleep(200 * time.Millisecond)
    }
}
```
This program waits for keyboard input asynchronously and loops, updating and drawing a point. When a WASD control is pressed, the callback `sprite.KeyboardController.OnKeyboardControl` is invoked and, as seen in the program, prints a given control on the terminal. In addition, `sprite` will move along the screen (since it has a vector with x velocity 1).
