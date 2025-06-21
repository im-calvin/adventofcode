## Planning

### Part 1

- Pair up the smallest number in the left list with the smallest number in the right list, then the second-smallest left number with the second-smallest right number
- Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those distances
- Calculate the total distance between the lists 

1. Take in the file
2. Split is between left column and right column (two lists)
3. Sort the list from smallest to largest
4. Map over both simultaneously (Python zip?) and then sum the distance

### Part 2

1. Create a frequency map for the right list
2. Iterate through left list and multiply by the frequency of the right list, adding it to the sum