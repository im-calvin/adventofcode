## Planning

### Part 1
- for each number, the first number represents the size of file, and then free space
- each file has a file Id (incrementing)

goal: move file blocks one at a time from the end into the empty spaces 
- calculate the checksum 

- we can recreate the string first to get the '0..111....22222'
- this smells like two pointer from here? start a pointer at the end and a pointer at the beginning. front pointer looks for `.`. back pointer looks for non-`.` until they meet in the middle. and then calculate the checksum afterwards
G
helpers: 
- checksum: calc until we reach a `.` 
- swap 2 nums j