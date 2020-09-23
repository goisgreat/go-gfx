# Golang Game Engine
## What is the *Golang Game Engine* and how do I use it?
The golang game engine is a "physics" library for golang. Here is an example to help explain usage.
```golang
package physics

import (
    time "time"
    physics "github.com/goisgreat/go-gfx"
)

func main() {
    stream := physics.EventStream{
        Input: make(chan byte)
    }
    sprite := physics.Point{
        physics.Position{0,0},
        physics.Vector{1,0},
        physics.KeyboardController{
            Input: stream.Input,
            KeyboardMap: physics.WASD,
        },
    }
    sprite.KeyboardController.OnKeyboardControl = func(control physics.KeyboardConrol) {
        println(control)
    }

    go physics.GrabInput(stream)
    
    sprite.Init()
    
    for {
        physics.ClearTerminal()
        sprite.Update()
        sprite.Draw()
        time.Sleep(200 * time.Millisecond)
    }
}
```
This program waits for keyboard input asynchronously and loops, updating and drawing a point. When a WASD control is pressed, the callback `sprite.KeyboardController.OnKeyboardControl` is invoked and, as seen in the program, prints a given control on the terminal. In addition, `sprite` will move along the screen (since it has a vector with x velocity 1).
## Where is the documentation?
The documentation for this package has yet to be implemented. Sorry about that!