import tensorflow as tf
import json
import numpy as np
from example_game import sample_game_json
from collections import defaultdict
save_path = './model/dnn/'

def main():
	# print(json.loads(sample_game_json))
	game = processGame(json.loads(sample_game_json))
	model = tf.keras.Sequential()

	model.add(tf.keras.layers.Dense(318, 
		input_shape=(10, 53, 3),
		activation='softmax'))

	model.add(tf.keras.layers.Dense(159))

	model.add(tf.keras.layers.Dense(52))

	model.add(tf.keras.layers.Dropout(0.5))

	model.compile(
		optimizer='rmsprop',
		loss='binary_crossentropy',
		metrics=['accuracy'])

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
		out.append(card)
	for i in range(4):
		if i == obj["ToPlay"]: continue
		for card in obj["Players"][i]["Hand"]:
			out.append(card)

	for card in out:
		card["inPlay"] = True

	for i in range(4):
		for card in obj["Players"][i]["Points"]:
			out.append(card)

	out = sorted(out, key=lambda v: (v["Suit"], v["Value"]))

	return np.array(list(map(flattenCard, out)))

def flattenCard(c):
	c = defaultdict(lambda:False, c)
	return (
		c["canMove"],
		c["inTrick"],
		c["inPlay"]
		)

def validMoves(hand, trick, heartsbroken):
	allClear = False
	if len(trick) == 0:
		if heartsbroken:
			allClear = True
		else:
			for card in hand:
			 if (card["Suit"] != 2 and not
			  	(card["Suit"] == 3 and
			     card["Value"] == 10)):
			 	card["canMove"] = True
	else:
		allClear = True
		for card in hand:
			if card["Suit"] == trick[0]["Suit"]:
				card["canMove"] = True
				allClear = False
	if allClear:
		for card in hand:
			card["canMove"] = True
	return hand

if __name__ == "__main__":
	main()