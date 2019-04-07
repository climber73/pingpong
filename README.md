# Ping-Pong
## Description
Simple application built upon blockchain [tendermint](https://github.com/tendermint/tendermint) designed to store ping-pong scores.

## Why
In our company we often play ping-pong in the office.
When we want to compare our skills in some tournament we use a whiteboard and markers to save the game state.
As programmers interested in blockchain we decided to use blockchain to store scores and evaluate who is current winner.

## Blockchain
When I was deciding which blockchain to use there were three known options for me: bitcoin, etherium and tendermint.
Bitcoin was rejected because of I didn't find any framework ready to use
which was appropriate for my level of awareness in the field of blockchain.
Etherium and tendermint were quite good.
So as a first step I chose tendermint as blockchain platform because it looks much easier and well described ABCI layer.
Then I might build the application using etherium.

### Transaction
Each transaction corresponds to one played game. So it's a JSON string as follows:
```
{
    name1: p1,
    name2: p2
}
```
where `name1` - name of player1 (string),
`name2` - name of player2 (string),
`p1` and `p2` are integer points of player1 and player2 correspondingly.

## Application logic
When a pair of player finished their game the result is committed to the blockchain.
From this moment the application begins to know abount two players and their points.

For example: player1 won player2 in 8 from 10 their games.
So the score of player1 is 0.8, and the player2's score is 0.2.
The state of the game can be presented as table below:

_scores:_|player1|player2
---|---|---
**player1**|-|0.8
**player2**|0.2|-

And player1 is a winner!

If player3 desides to play then each player gets one more score for player3. If player3 wins player1 then player3's score against player1 will become 1.0 and player1's score against player3 will become 0.0 corespondingly.
Now the state of the game will be:

_scores:_|player1|player2|player3
---|---|---|---
**player1**|-|0.8|0.0
**player2**|0.2|-|0.0|
**player3**|1.0|0.0|-|

Congrats to payer3 now :)
