# Alpaca-Take-Home
Alpaca's Take Home Exam solution

### Problem :thinking_face:
This problem is related to the Hungarian lottery. In case you are not familiar with it: players pick 5 distinct numbers from 1 to 90. There is a weekly lottery picking event when the lotto organization picks 5 distinct numbers randomly between 1 and 90 – just like the players did. The player’s reward then depends on how many of the player’s numbers match with the ones selected at the lotto picking. A player wins if he/she has 2, 3, 4 or 5 matching numbers.

Now, the problem: at the lottery event, right after picking the numbers, a computer shall be able to report quickly that how many winners are in each category, for example:

| Numbers matching | Winners |
| ------------- |:-------------:
| 5             | 0             |
| 4             | 12            |
| 3             | 818           |
| 2             | 22613         |


This report shall be generated within a couple of seconds after picking the winner number. The player’s numbers are known in advance – at least 1 hour ahead of the show. In peak periods there are about 5 million players, but to be sure we can assume it is not exceeding 10 million.

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

In my solution, I implemented a boolMap `map[uint]*[]bool` and compared it against a bitMap implementation `map[uint]*[]byte`.
The boolMap implementation beats the bitMap implementation in terms of speed to query but the bitMap implementation is much more efficient
in terms of storing data since it is just storing the match column information using 1 bit meaning a single byte in a byte array which makes the
memory usage 8 times lower than the boolMap implementation.

<img width="1201" alt="Screen Shot 2021-08-02 at 9 55 44 PM" src="https://user-images.githubusercontent.com/22711718/127876492-252a186a-d63a-482d-b1ba-545d2521545d.png">

<img width="1032" alt="Screen Shot 2021-08-02 at 10 24 24 PM" src="https://user-images.githubusercontent.com/22711718/127877149-b23c2caa-ea76-4411-8467-0b8c45e29714.png">

**Memory used when using a bitMap**
<img width="942" alt="Screen Shot 2021-08-02 at 10 27 21 PM" src="https://user-images.githubusercontent.com/22711718/127877426-ab7965ad-dfe4-4396-8c08-42d734ff2099.png">

**Memory used when using a boolMap**
<img width="954" alt="Screen Shot 2021-08-02 at 10 28 49 PM" src="https://user-images.githubusercontent.com/22711718/127877578-4dfb451f-f556-425e-aa39-1189c798311d.png">

