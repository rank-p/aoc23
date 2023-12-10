import itertools
with open("aoc8.txt") as f:
	text = f.read()
	instructions , map = text.split("\n\n")

parsed_map = {key: tuple(value[1:-1].split(", ")) for key, value in [x.strip().split(" = ") for x in map.split("\n")][:-1] }	
#start = map.split("\n\n")[0].split(" = ")[0]
start = "AAA"
print(instructions)
print(start)

steps=0
for char in itertools.cycle(instructions):
	if steps % 1000000 == 0: print(steps)
	if char == 'L':
		start = parsed_map[start][0]
	else:
		start = parsed_map[start][1]
	steps += 1
	if start == "ZZZ":
		break
	
print(steps)
