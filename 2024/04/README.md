## Planning

### Part 1

- need to find all instances of `XMAS` 
- can appear left to right, right to left, top to bottom, bottom to top, diagonal top left to bottom right, diagonal top right to bottom left, diagonal bottom left to top right, diagonal bottom right to top left

- I can take the input in and put it as a 2d array, and then iterate through it. whenever I see an `X`, I can do a search. This is pretty slow, and exploits that the word `XMAS` contains no duplicates of `X`
