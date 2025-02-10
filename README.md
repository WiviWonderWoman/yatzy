# Y A T Z Y !
This is a project I'm doing just for fun and to learn [Gio UI](https://gioui.org/) a
cross-platform GUI for Go. If yore not familiar whit YATZY read the [gameplay](#Gameplay) instructions below.

## How to get set up and run
* First you need to have[GO](https://go.dev/doc/install) installed. 
* Clone this reposiory:
    ```bash
    git clone https://github.com/WiviWonderWoman/yatzy.git
    ```
* Navigate inside:
    ```bash
    cd yatzy
    ```
* Run application:
    ```bash
    go run cmd/main.go
    ```
___
___

## Gameplay
* Yatzy can be played solitaire or by any number of players. 
* The player rolls five dice. 
* After each roll, the player chooses which dice to keep, and which to reroll. 
* A player may reroll some or all of the dice up to two times on a turn. 
* The player must put a score or zero into a score box each turn. 
* The game ends when all score boxes are used. 
* The player with the highest total score wins the game.

### Scoring
The following combinations earn points:

#### Upper Section:

* Ones: The sum of all dice showing the number 1.
* Twos: The sum of all dice showing the number 2.
* Threes: The sum of all dice showing the number 3.
* Fours: The sum of all dice showing the number 4.
* Fives: The sum of all dice showing the number 5.
* Sixes: The sum of all dice showing the number 6.

If a player manages to score at least 63 points (an average of three of each number) in the upper section, they are awarded a bonus of 50 points.

#### Lower Section:

* *One Pair*: Two dice showing the same number. 
    * Score: Sum of those two dice.
* *Two Pairs*: Two **different** pairs of dice. 
    * Score: Sum of dice in those two pairs.
* *Three of a Kind*: Three dice showing the same number. 
    * Score: Sum of those three dice.
* *Four of a Kind*: Four dice with the same number. 
    * Score: Sum of those four dice.
* *Small Straight*: The combination [1] [2] [3] [4] [5]. 
    * Score: 15 points (sum of all the dice).
* *Large Straight*: The combination [2] [3] [4] [5] [6]. 
    * Score: 20 points (sum of all the dice).
* *Full House*: Any set of three combined with a **different** pair. 
    * Score: Sum of all the dice.
* *Chance*: Any combination of dice. 
    * Score: Sum of all the dice.
* *Yatzy*: All five dice with the same number. 
    * Score: 50 points.


Some combinations offer the player a choice as to which category to score them under.<br /> E.g. [2] [2] [5] [5] [5] would score :
* **19** in *Full House* or *Chance*
* **15** in *Fives* or *Three of a Kind*
* **14** in *Two Pairs*
* **10** in *One Pair*
* **4** in *Twos*
* **0** in any other category