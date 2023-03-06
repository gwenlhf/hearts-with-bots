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