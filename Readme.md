# HeartsWithBots: a Neural Network experiment
HeartsWithBots is a neural network project gone awry, here for the world to see, and hopefully learn from.

## Quick-Start
To see a rendered example of a game and AI prediction in your browser (requires NodeJS), `cd` to `./src/front` and run:
```
npm install
ng serve --open
```
If you want to see the bot make a prediction live (requires Python with tensorflow + h5py), `cd` to `./src/ai` and run:
```
python3 bot_demo.py
```
If you want to see a new game state generated live (requires Golang), `cd` to `./src/rules`, run
```
go run Main.go Hearts.go Trainer.go
```
and then navigate your browser to `http://localhost:8080/new`.

## Architecture
HeartsWithBots combines an Angular frontend with JSON output from a Golang rules service and a Python AI service. Also included are the batch files used in conjunction with the Golang service to generate the target/training data, located in raw JSON form under `./src/rules/train/`.

The Python AI (`./src/ai/`) has methods for training the AI, as well as providing prediction data for a given game state.

The Golang service (`./src/rules/`) has a host of useful methods, including the game rules logic (`Hearts.go`) and the method used to generate training sets (`Trainer.go`).

### The Game Object
A normal, unflattened game state looks like this:
```
{
	Players: Player[],
	Trick: Card[],
	ToPlay: int,
	HeartsBroken: boolean,
}
```
### The Player Object

A player object looks like this:
```
{
	Side: int,
	Hand: Card[],
	Points: Card[],
}
```

### The Card Object

A card object looks like this:
```
{
	Suit: int,
	Value: int,
}
```
All of these objects are shared between Angular, Python, and Golang. For the neural network, the game object is flattened into a 2d-array of type `float[53][3]`, where the first 52 indices reference cards in Suit-major order and record the booleans `canMove, inTrick, inPlay`, and the remaining index is used for the extra boolean values `oppenentHasAllPoints, selfHasAllPoints, heartsBroken`.

The neural network output is a 1d-array of type `float[52]`, with higher values predicting fewer points taken for the player if the card is selected.

## The Abridged History of HeartsWithBots
The original idea behind HeartsWithBots was a website. Specifically, a website where you could play Hearts against an intelligent AI, yes, but also against a human opponent, and a bot that picked moves at random. After the game, the human players would be asked to pick which player was which, as a sort of mini-Turing test.

To make a long story short, that turned out to be harder than I anticipated.

Working towards this goal, development time was split between all three parts (the frontend, the rules engine, and the AI component), and the output quality was proportionally diminished. What sounded like a simple project on paper quickly turned into development hell -- splitting the logic into different components sounded great, but integrating them involved a lot of code duplication, which lead to more bugs, which lead to the project iterations taking far longer than expected, which lead to what you see here.

The end result is a neural network with a 57% accuracy rating against a set of training data that has flaws, a rules engine that hasn't been fully tested, and a frontend which, despite looking great and taking the least amount of time to develop, can only display static sample data.

I have spent enough time on this project to recognize that it needs to be rewritten, but it will be a long time before I'm ready to sit down and go through it all again. The project was, however, a wonderful introduction to neural networks, and I learned a great deal about the subject in the process of building it. Had I spent less time on the integration and architecture, and more time developing a cohesive set of training data, I have little doubt that this project would have achieved greater success.

To me, the largest takeaways here were as follows:
+ An untrained network with a good model with good training data is better than a well-trained network with an inferior model and training data
+ It's generally better to start with a much smaller project scope than you need and work up, than to start with a larger scope than you need and work down
+ Working with lots of different languages and services connected by APIs is not worth it unless you have the infrastructure in place already
