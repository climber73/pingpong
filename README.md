# Ping-Pong
## Description
Simple application built upon blockchain [tendermint](https://github.com/tendermint/tendermint) designed to store ping-pong scores.

## Why
In our company we often play ping-pong in the office.
When we want to compare our skills in some tournament we use a whiteboard and markers to save the game state.
As programmers interested in blockchain we decided to use blockchain to store scores and evaluate who is current winner.

## Tendermint
### Transaction
Each transaction corresponds to one played game. So it's a string as follows:
`<name1> <p1>:<p2> <name2>`,
where name1 - name of player1,
name2 - name of player2,
p1 and p2 are points of player1 and player2 correspondingly.
