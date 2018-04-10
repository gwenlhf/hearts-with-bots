import tensorflow as tf
import json
import numpy as np
from example_game import sample_game_json
from collections import defaultdict
save_path = './model/dnn/'

def main():
	# print(json.loads(sample_game_json))
	game = processGame(json.loads(sample_game_json))
	print(len(list(game)))
	# model = Sequential()

	# model.add(Dense(318, 
	# 	input_shape=(10, 53, 3),
	# 	activation='softmax'))

	# model.add(Dense(159))

	# model.add(Dense(52))

	# model.add(Dropout(0.5))

	# model.compile(
	# 	optimizer='rmsprop',
	# 	loss='binary_crossentropy',
	# 	metrics=['accuracy'])

# processGame turns a Game state object
# into a flat 2d list with the following shape:
# [Card] (sorted by Suit/Value)
# 	[canMove,
# 	inTrick,
# 	inPlay]
# ,
# [OpponentHasAllPoints, SelfHasAllPoints, 0]
def processGame(obj):
	out = []

	for card in obj["Trick"]:
		card["inTrick"] = True
		out.append(card)

	hand = obj["Players"][obj["ToPlay"]]["Hand"]
	valid = validMoves(hand, obj["Trick"], obj["HeartsBroken"])

	for card in valid:
		card["canMove"] = True
		out.append(card)

	for i in range(3):
		if i == obj["ToPlay"]: continue
		for card in obj["Players"][i]["Hand"]:
			out.append(card)

	for card in out:
		card["inPlay"] = True

	for i in range(3):
		for card in obj["Players"][i]["Points"]:
			out.append(card)

	out = sorted(out, key=lambda v: (v["Suit"], v["Value"]))

	return map(flattenCard, out)

def flattenCard(c):
	c = defaultdict(lambda:False, c)
	return (
		c["canMove"],
		c["inTrick"],
		c["inPlay"]
		)

def validMoves(hand, trick, heartsbroken):
	out = []
	if len(trick) == 0:
		if heartsbroken:
			return hand[:]
		else:
			return [x for x in hand if 
				(x["Suit"] != 2 and not 
				(x["Suit"] == 3 and x["Value"] == 10))]
	for card in hand:
		if card["Suit"] == trick[0]["Suit"]:
			out.append(card)
	if len(out) == 0:
		return hand[:]
	return out

if __name__ == "__main__":
	main()