import tensorflow.keras.models
import json

save_path = './model/dnn/'

def __main__():
	model = Sequential()

	model.add(Dense(318, 
		input_shape=(10, 53, 3),
		activation='softmax'))

	model.add(Dense(159))

	model.add(Dense(52))

	model.add(Dropout(0.5))

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
	