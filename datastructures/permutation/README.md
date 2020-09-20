# This is how the recursion order works

For array [A B C D]

```
HOLD D -- ITERATION 0
  HOLD C
    A B C D
    B A C D
      SWITCH B - C
  HOLD B
    C A B D
    A C B D
      SWITCH A - B
  HOLD A
    B C A D
    C B A D
      FINALLY SWITCH D - C
HOLD C -- ITERATION 1
  HOLD A
    D B A C
    B D A C
      SWITCH B - A
  HOLD B
    A D B C
    D A B C
      SWITCH D - B
  HOLD D
    B A D C
    A B D C
      FINALLY SWITCH B - C
HOLD B -- ITERATION 2
  HOLD D
    A C D B
    C A D B
      SWITCH C - D
  HOLD C
    D A C B
    A D C B
      SWITCH A - C
  HOLD A
    C D A B
    D C A B
      FINALLY SWITCH A - B
HOLD A -- ITERATION 3
  HOLD B
    D C B A
    C D B A
      SWITCH C - B
  HOLD C
    B D C A
    D B C A
      SWITCH D - C
  HOLD D
    C B D A
    B C D A
```
