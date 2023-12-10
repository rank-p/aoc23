with open("aoc9.txt") as f:
	lines = f.read().split("\n")[:-1]
	lines = [[int(x.strip()) for x in lines[i].split(" ")] for i in range(len(lines))]

def extrapolate(seq):
	n = [seq[i] - seq[i-1] for i in range(1, len(seq))]
	if all([x==0 for x in n]): return n[-1] 	
	return n[-1] + extrapolate(n)

print(sum([seq[-1] + extrapolate(seq) for seq in lines]))
