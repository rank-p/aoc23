import itertools
import math
from functools import reduce
with open("aoc8.txt") as f:
	text = f.read()
	instructions , map = text.split("\n\n")

parsed_map = {key: tuple(value[1:-1].split(", ")) for key, value in [x.strip().split(" = ") for x in map.split("\n")][:-1] }	
#start = map.split("\n\n")[0].split(" = ")[0]
start = [key for key in parsed_map.keys() if key[-1]=='A'] 
print(instructions)
print(start)


def findCycle(key, instructions, map):
	steps=0
	for char in itertools.cycle(instructions):
		idx = 0 if char == 'L' else 1
		key = map[key][idx]
		steps += 1
		if key[-1]=='Z' : break
	return steps, key 
	

def lcm(a, b):
	return abs(a*b) // math.gcd(a,b)

cycle = []
for key in start:
	cycle.append(findCycle(key,instructions,parsed_map)[0])

print(reduce(lambda x, y: lcm(x,y), cycle))
	

