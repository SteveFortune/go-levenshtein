# Levenshtein Edit Distance

A Golang implementation of the Levenshtein Edit Distance algorithm.

### Minimum Edit Distance

- How similar are 2 strings?
- Example:
  - User gives string "graffe"
  - Which of the following is closest?
    - "graf"
    - "graft"
    - "grail"
    - "giraffe"
  - Correct their spelling
- Minimum number of edit operations needed to transform one string into another:
  - Insertions: "xyz" to "uxyz"
  - Deletions: "xyz" to "xz"
  - Substitutions: "xyz" to "ayz"
- Example: turning "intention" into "execution":

```
I   N   T   E   *   N   T   I   O   N
|   |   |   |   |   |   |   |   |   |
*   E   X   E   C   U   T   I   O   N
(d) (s) (s)     (i) (s)

```
  - 1 deletion
  - 3 substitutions
  - 1 insertion
  - Levenshtein distance is 8 (substitutions cost 2)

### Algorithm

- Search for a path from the start string to the final string, where the path is a sequence of edit operations.
- The initial state is the word that we're transforming.
- The operators are: insert, delete and substitute.
- The end state is the word that we're aiming for.
- Path cost is the number of edits (what we want to reduce as much as possible).
- There are lots of possible sequence paths, so we can't navigate naively.
- Lots of the paths end up at the same state, so we don't care about most of them. Just the shortest paths to each revisited states
- Definitions:

  - For
    - `x` a string of length `n`
    - `y` a string of length `m`

  - `d(i,j)`
    - the edit distance between `x[1..i]` and `y[1..j]`
    - thus the edit distance between `x` and `y` is `d(n,m)`

- Initialization:

  - The difference between the first `i` characters of `x` and the null `y` string is the cost of deleting all of those characters: `d(i,0) = i`
  - The difference between the first `j` characters of `y` and the null `x` string is the cost of inserting all of those characters: `d(0,j) = j`
  - I.e. cost of those is the length of the strings.

- Recurrence relation:

```
For each i = 1...m
  For each j = 1...n
                  { d(i-1,j) + 1
    d(i,j) = min  { d(i,j-1) + 1
                  { d(i-1,j-1) +  1; { if x(i) ≠ y(j)
                                  0; { if x(i) = y(j)

```

### Applications

- Spelling checking
- Diff-ing sentences to determine how close they are (machine translation accuracy, paraphrasing, etc)
- Named entity extraction;
  - How likely is it that "IBM" and "IBM Inc" refer to the same entity?
  - How likely is it that "U.S. President Barack Hussein Obama" and "President Barack Obama" refer to the same person?

### References
