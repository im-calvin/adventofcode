## Planning

### Part 1

antinode: 
  1. two antennas must be perfectly in line.
  2. one antenna is twice as far away as the other 
  - for any pair of antennas, there are two antinodes, one on either side
    - we can do half and then reflect? 
  - for n >= 2 antennas of the same freq, we get:
  1. S(2) = 1
  2. S(n) = S(n-1) + (n-1) "lines"
  - where each "line" means 2 nodes

question: how many unique locations within the bounds of the map contain an antinode? 

approach: 
1. iterate through the map, whenever you find an antenna, add it to the frequency list, and run the helper
1a. helper(antennas, map) slice[int] (new antinodes)
- antennas[0:n-1] = old antennas
- antennas[-1] = new antenna

- distance calculation: for every old antenna, with the new antenna calculate where the antinodes would be:
1. (x1 - x2), (y1 - y2) = (z1, z2) * 2 
2. node1: (x1, y1) - (z1, z2) * 2
3. node2: (x2, y2) + (z1, z2) * 2
- check that they are in bounds and then add them to the result.
2. add the antinodes to a separate map

3. at the end check the number of antinodes in the antinode map

