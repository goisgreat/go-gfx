A note
-
Short version
--
This project uses a right-down coordinate system.
For example, `X=0, Y=0` is in the *top left*. `X=1, Y=1` would be over 1 and down 1 from the top left.

Long version
--
The X-Y coordinates of this program may be confusing. You might be wondering, "What on earth! Why is X=0 Y=0 *not* in the middle?" Fear not! Basically, since when accessing a matrix with `m[0][0]` where m is this:
```golang
[][]bool{
    {0,0,0,0},
    {0,0,0,0},
    {0,0,0,0},
    {0,0,0,0},
}
```
would be in the top left (the first `0` you see). And this is *exactly* how this program is implemented! So, if an object as x `4` and y `4` (in a `10` by `10` matrix), it would be rendered in the middle. So if your code is:
```golang
package main

import (
    "github.com/goisgreat/go-physics"
)

func main() {
    point := physics.point{
        physics.Position{0,0},
        ...
    }
    point.Draw(...)
    ...
}
```
and you want `point` to be rendered in the middle of an `3` by `3` grid, simply change `point`'s position to `Position{1, 1}`