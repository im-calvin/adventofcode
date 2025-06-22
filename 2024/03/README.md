## Planning

### Part 1

- need to find strings of the form: mul(xxx, xxx). that is "mul(" followed by 1 to 3 nums, followed by ", " followed by 1 to nums, followed by ")"

- regex: `(mul\()(\d{1,3},\d{1,3}\))`

- then from this find all the digits
- regex: `(\d{1,3}),(\d{1,3})` -> copy "843,597" -> split the string on `,`

### Part 2

- regex: mul\(\d{1,3},\d{1,3}\)|don\'t\(\)|do\(\)