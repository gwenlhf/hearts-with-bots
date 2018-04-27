for /l %%x in (1000,-1,1) do (
	rules.exe -step %%x
	SLEEP 1
	)