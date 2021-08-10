# Alpaca-Take-Home
Alpaca's Take Home Exam solution

### Problem
This problem is related to the Hungarian lottery. In case you are not familiar with it: players pick 5 distinct numbers from 1 to 90. There is a weekly lottery picking event when the lotto organization picks 5 distinct numbers randomly between 1 and 90 – just like the players did. The player’s reward then depends on how many of the player’s numbers match with the ones selected at the lotto picking. A player wins if he/she has 2, 3, 4 or 5 matching numbers.

Now, the problem: at the lottery event, right after picking the numbers, a computer shall be able to report quickly that how many winners are in each category, for example:

| Numbers matching | Winners |
| ------------- |:-------------:
| 5             | 0             |
| 4             | 12            |
| 3             | 818           |
| 2             | 22613         |


This report shall be generated within a couple of seconds after picking the winner number. The player’s numbers are known in advance – at least 1 hour ahead of the show. In peak periods there are about 5 million players, but to be sure we can assume it is not exceeding 10 million.

### Assumptions:
1. A bet that has non-unique values will be discarded: e.g. 22 32 34 78 22
2. A bet that has a value that is not within the 1 - 90 bound will be discarded: e.g. 22 -7 91 11 12

### Libraries used:
1. bitset: https://github.com/bits-and-blooms/bitset

   - Used the library for bitwise operations instead of using save space instead of using a `[]bool` which is 8x more expensive in terms of memory compared to using `math.bits` 

### My Solution

The algorithm for finding the matches is scanning the entries or bets and checking each number whether it has a match in 
the given winning picks. In relation to this, as you are scanning the bets, you will have to maintain a map pointer to an
integer value so that if there are 2 or more matching numbers, you can update the specific map key as you go along. 
The challenge is how one can achieve a performance that outputs the result in a couple of second or even faster.

The problem is similar as to how database works — given a local file, you will load it when needed into memory and give an
output based on the declarative command that the user wants. Declarative meaning the user is not concerned on how you get
the results, they are only concerned _with_ the results. 

A tried and tested way of improving query speeds is using a database index. A widely known data structure is a B-tree/B-tree+
wherein locations to disk memory are maintained using an M-nary tree to create a quick path of accessing files. This, however,
requires that the records have an order but since the algorithm for finding matches do not need an order (Combination vs Permutations),
one will have to look elsewhere and this leads me to a Bitmap. 

I decided to use a BitMap because of the low cardinality of possible values within a bet which fits it perfectly to the data structure. Another
reason that the operation that is needed to perform to return the values is fairly simple and will fit to bitwise operations. 

Time Complexity: O(X * N) where X is the width of the bets and N is the maximum number of bettors which is in our case 10 million

Space Complexity: O(M * N) where M is the cardinality which is 90 and N is the maximum number of better which is in our case is 10 million. 
In our case, we can say O(M * (N/8)) since we used bits to store the bet information instead of a whole byte.

### How to run?
1. Run `make run-alpaca-input-file` to use the alpaca provided file

### How to run tests?
1. For Benchmarks, run `make test-benchmark`

2. For Unit tests, run `make test-correctness`
