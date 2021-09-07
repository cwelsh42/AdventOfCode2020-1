# Description
This solution was developed to solve the problem set here: https://adventofcode.com/2020/day/1

# Complexity
This solution tends to O(1). 

The two complex processes taking place are changing the given array of integers into a hashmap (insertion) and checking for the existence of a matching pair in the hashmap (lookup).

- Hash table insertions are theoretically O(1), although Golang makes no official remarks regarding their complexity. 
- Lookups in a hashtable are generally constant time, so O(1).

# Assumptions
- You may want to pass a different value for the pairs to sum up to, other than 2020.
- There is one result. (If there were multiple solutions you would want to store these and return the array of solutions.)
- If there are multiple results, you only care about the first. (currently returns on the first valid solution)
- No input parsing/validation has been completed. You may wish to include some sanitising function or a reader which reads the values from a file.
