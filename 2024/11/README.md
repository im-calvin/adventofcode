## Planning

### Part 1

Rules: 

  - If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
  - If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
  - If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.

- we probably need some sort of struct to keep track of the stones
  - it would hold the contents, the number of digits
- create helpers to act on this struct that does the rules calculation

