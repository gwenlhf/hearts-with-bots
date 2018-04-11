import tensorflow as tf
import json
import numpy as np
from example_game import sample_game_json
from collections import defaultdict
import urllib
from random import randint
import h5py
import time
from tempfile import TemporaryFile

save_path = './model/dnn/weights_3.h5'

training_path = './train/'

def main():
	# batchConvertTrainingJson('../rules/train')
	# trainModel()
	print(predictMove(json.loads(sample_game_json)))

def trainModel():
	model = loadModel()

	traindata = np.load("%s%sdata.npz" % (training_path, "mick"))["arr_0"]
	targedata = np.load("%s%sdata.npz" % (training_path, "rock"))["arr_0"]
	traindata = np.reshape(traindata, (10010, 53, 3))
	targedata = np.reshape(targedata, (10010, 52))

	model.fit(traindata, targedata, batch_size=10, epochs=50)
	model.save_weights(save_path)

def loadModel():
	model = tf.keras.Sequential()
	model.add(tf.keras.layers.InputLayer(batch_input_shape=(10,53,3)))
	model.add(tf.keras.layers.Flatten())
	model.add(tf.keras.layers.Dense(2756,
		activation='softmax'))
	model.add(tf.keras.layers.Dense(2756))
	model.add(tf.keras.layers.Dense(2756))
	model.add(tf.keras.layers.Dense(689))
	model.add(tf.keras.layers.Dense(52))
	model.add(tf.keras.layers.Dropout(0.5))
	model.compile(
		optimizer='rmsprop',
		loss='binary_crossentropy',
		metrics=['accuracy'])
	model.load_weights(save_path)
	return model

def signb(b):
	return 1 if b else -1

def batchConvertTrainingJson(traindir, num=1001):
	mick = np.empty((num, 10, 53, 3), np.int8)
	rock = np.empty((num, 10, 52), np.int8)
	for i in range(1, num):
		trstr = "%s%s%d" % (traindir, "/mick", i)
		tastr = "%s%s%d" % (traindir, "/rock", i)
		try:
			vs = np.vectorize(signb)
			traindata = open(trstr, 'r')
			targedata = open(tastr, 'r')
			trdata = json.load(traindata)[10:]
			tr = vs(np.array(trdata))
			mick[i-1] = tr
			tadata = json.load(targedata)[10:]
			ta = vs(np.array(tadata))
			rock[i-1] = ta
		except Exception as e:
			print(e)
			continue
		finally:
			traindata.close()
			targedata.close()
	np.savez("%s%s" % (training_path, "mickdata"), mick)
	np.savez("%s%s" % (training_path, "rockdata"), rock)

# this isn't working
# def createTrainingData(training, targets, batch_size=10, cap=1000):
# 	while cap > 0:
# 		batch_training = []
# 		batch_target = []
# 		for i in range(batch_size):
# 			game = newGame()
# 			while len(game["Trick"]) < 3:
# 				valid = validMoves(game)
# 				card = valid[np.random.randint(0,len(valid))]
# 				try:
# 					game = wrapMove(game, moveFromCard(card, game["ToPlay"]))
# 				except:
# 					continue
# 			valid = validMoves(game)
# 			bst = valid[0]
# 			hiscore = 26
# 			for card in valid:
# 				tmp = wrapMove(game, moveFromCard(card, 3))
# 				if tmp == -1:
# 					continue
# 				if tmp["Players"][game["ToPlay"]]["Total"] < hiscore:
# 					hiscore = tmp["Players"][game["ToPlay"]]["Total"]
# 					bst = card
# 			batch_training[i] = flattenGame(game)
# 			batch_target[i] = flattenMove(bst)
# 		training[cap-1] = batch_training
# 		targets[cap-1] = batch_target
# 		cap -= 1
# 		print("finished a batch! %d remaining to fill" % cap)



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
			card["Value"] = i % 13
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

def playGameDemo(ng=newGame, mv=wrapMove, model=loadModel()):
	game = ng()

def predictMove(game, model=loadModel()):
	s = game["ToPlay"]
	ia = np.empty((10,53,3))
	ia[0] = flattenGame(game)
	fm = model.predict(ia, batch_size=10)
	return unflattenMove(fm[0], s)

if __name__ == "__main__":
	main()
	pass
