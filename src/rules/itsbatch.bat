
for /l %%x in (986,-1,1) do (
	rules.exe -step %%x
	SLEEP 1
	)