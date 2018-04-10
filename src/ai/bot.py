import tensorflow as tf
import json
import numpy as np
from example_game import sample_game_json
from collections import defaultdict
import urllib
from random import randint
import h5py

save_path = './model/dnn/weights.h5'

training_path = './train/training_data.h5'

def main():
	f = h5py.File(training_path)
	traindata = f["mick"]
	targedata = f["rocky"]
	createTrainingData(traindata, targedata)
	# print(json.loads(sample_game_json))
	# game = flattenGame(json.loads(sample_game_json))
	# b_games = np.tile(game, (1,10,1,1))
	# print(b_games)
	# # urlopen("localhost:8080/")
	# model = tf.keras.Sequential()
	# model.add(tf.keras.layers.Dense(318, 
	# 	input_shape=(10, 53, 3),
	# 	activation='softmax'))
	# model.add(tf.keras.layers.Dense(159))
	# model.add(tf.keras.layers.Dense(52))
	# model.add(tf.keras.layers.Dropout(0.5))
	# model.compile(
	# 	optimizer='rmsprop',
	# 	loss='binary_crossentropy',
	# 	metrics=['accuracy'])
	# model.load_weights(save_path)
	# res = model.predict(b_games)


def createTrainingData(training, targets, batch_size=10, cap=1000):
	while cap > 0:
		batch_training = []
		batch_target = []
		for i in range(batch_size):
			game = newGame()
			while len(game["Trick"]) < 3:
				valid = validMoves(game)
				card = valid[np.random.randint(0,len(valid))]
				try:
					game = wrapMove(game, moveFromCard(card, game["ToPlay"]))
				except:
					continue
			valid = validMoves(game)
			bst = valid[0]
			hiscore = 26
			for card in valid:
				tmp = wrapMove(game, moveFromCard(card, 3))
				if tmp == -1:
					continue
				if tmp["Players"][game["ToPlay"]]["Total"] < hiscore:
					hiscore = tmp["Players"][game["ToPlay"]]["Total"]
					bst = card
			batch_training[i] = flattenGame(game)
			batch_target[i] = flattenMove(bst)
		training[cap-1] = batch_training
		targets[cap-1] = batch_target
		cap -= 1
		print("finished a batch! %d remaining to fill" % cap)



def moveFromCard(card, side):
	return {
		"Card":card,
		"Side":side
	}

def newGame(endpoint="http://localhost:8080/new"):
	by = urllib.request.urlopen(endpoint)
	return json.load(by)

def wrapMove(game, move):
	return sendMove({
		"Game":game,
		"Move":move
		})

def sendMove(wrapd, endpoint="http://localhost:8080/mov"):
	rq = urllib.request.Request(endpoint,
		data=bytes(json.dumps(wrapd), 'utf-8'),
		headers={
			"Content-Type":"application/json"
		})
	by = urllib.request.urlopen(rq)
	return json.load(by)

# flattenGame turns a Game state object
# into a flat 2d list with the following shape:
# [Card] (sorted by Suit/Value)
# 	[canMove,
# 	inTrick,
# 	inPlay]
# ,
# [OpponentHasAllPoints, SelfHasAllPoints, HeartsBroken]
def flattenGame(obj):
	out = []

	for card in obj["Trick"]:
		card["inTrick"] = True
		out.append(card)

	valid = validMoves(obj)

	for card in valid:
		out.append(card)
	for i in range(4):
		if i == obj["ToPlay"]: continue
		for card in obj["Players"][i]["Hand"]:
			out.append(card)

	for card in out:
		card["inPlay"] = True

	m = 0
	allPoints = -1
	for i in range(4):
		k = 0
		for card in obj["Players"][i]["Points"]:
			allPoints = -1
			k += 1
			m += 1
			out.append(card)
		if m > 0 and k == m:
			allPoints = i
			
	out = sorted(out, key=lambda v: (v["Suit"], v["Value"]))
	
	xtr = []
	if allPoints > -1:
		if obj["ToPlay"] == allPoints:
			xtr = (0,1,obj["HeartsBroken"])
		else:
			xtr = (1,0,obj["HeartsBroken"])
	out.append(xtr)
	return np.array(list(map(flattenCard, out)))

def unflattenMove(obj, side):
	card = { }
	top = np.sort(np.copy(obj))[0]
	for i in range(len(obj)):
		if obj[i] == top:
			card["Suit"] = i / 13
			card["Value"] = k % 13
	return {
		"Side":side,
		"Card":card
	}

def flattenMove(card):
	res = np.zeroes(52,3)
	res[card["Suit"] * 13 + card["Value"]] = 1
	return res

def flattenCard(c):
	c = defaultdict(lambda:False, c)
	return (
		c["canMove"],
		c["inTrick"],
		c["inPlay"]
		)

def validMoves(game, exclude=False):
	hand = game["Players"][game["ToPlay"]]["Hand"]
	trick = game["Trick"]
	heartsbroken = game["HeartsBroken"]
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
	if exclude:
		return [x for x in hand if x["canMove"]]
	return hand

if __name__ == "__main__":
	main()
	pass