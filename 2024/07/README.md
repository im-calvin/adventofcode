## Planning

### Part 1
- smells like dynamic programming, we need to branch out at each ` ` and then try either `+` or `*`. 
- we can early stop when we overshoot since the only operators we have are `+` and `*` (increasing) as long as the input is positive, which it is

let's make a more concrete plan!

1. take the input, parse each line. for each line, the left of the colon will be `result`, the right is a slice which will be split on ` `. 
2. we define a helper `multiply` and `add`
3. `curr = nums[0]`. we start at `i=1` and increase `while i < len(nums)`. we do this because it's easier for dp. for each number:
  - break early if the input is too big!
  - check the bounds of `i`
  - stop recursing if `curr == result`
  - try adding
  - try multiplying

  ### Part 2

  