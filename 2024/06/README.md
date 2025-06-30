## Planning

### Part 1
- build the maze, track the guard, denote places he's been to with `X`
- if you see a `#` change direction
- use a `while(true)` loop?
- at the end grab all the `X`'s

### Part 2
- making a loop is the same as making sure the guard crosses a new `X` at some point
- we should still make that `X` map, and then every time there's a box then put an `O` there
  - a 'box' would mean that the guard will create a box
  - this means that we *shouldn't* make the `X` map, but calculate this inline.
- as we make our `X` map, if we were to cross over an `X` at any point, then that's a location we can add a possible box
- this won't work because we need to account for the orientation of the guard whenever he reaches an `X`. we could make an `equals()` func to check for orientation AND location
  - each square could have multiple orientations...

- new idea, for each square we travel, we try to place an obstacle down in front of us. you can only place one down if when you place one down you match the orientation and location.. (and also an obstacle wasn't placed there to begin with..)

- there's 2 types of obstacles that we can place
1. we've crossed where we've been through before and it forms a new box
2. we haven't visited it before but it'll stil form a box... 
  - we might be able to deal with this by looping through the visitedMap and then seeing if any of them form boxes afterwards
  - can't we condense #1 and #2 into one by doing this? 