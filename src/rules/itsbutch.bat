for /l %%x in (1000,-1,1) do (
	SLEEP 5
	taskkill /F /IM rules.exe
	)