# game-of-life-interview

https://playgameoflife.com/

## Expectation
* Create a data structure to hold your world
* Implement function to randomly generate a world
* Implement function to count the number of neighbors around a cell
* Implement function to determine if a cell will live or die next generation
* Display world into standard output
* Implement testing and adhere to good programming principles
* Feel free to create new functions or change existing functions as you see fit!
* Feel free to implement these functions in any programming language.
* Don't worry about if the program can run or not. This is  not a syntax quiz. 


## Conway's Game of Life Rules
1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
2. Any live cell with two or three live neighbours lives on to the next generation.
3. Any live cell with more than three live neighbours dies, as if by overpopulation.
4. Any dead cell with exactly 3 live neighbors becomes a live cell, as if by reproduction.

Example Output Generation

```text
Starting 3x3 World
|0|1|0|
|0|1|0|
|0|1|0|

Next Generation 

|0|0|0|
|1|1|1|
|0|0|0|

Next Generation

|0|1|0|
|0|1|0|
|0|1|0|
```
