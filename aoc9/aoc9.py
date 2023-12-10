with open("aoc9.txt") as f:
	lines = f.read().split("\n")[:-1]
	lines = [[int(x.strip()) for x in lines[i].split(" ")] for i in range(len(lines))]

def extrapolate(seq):
	n = [seq[i] - seq[i-1] for i in range(1, len(seq))]
	if all([x==0 for x in n]): return 0 	
	return  n[0] - extrapolate(n) 

print([seq[0] - extrapolate(seq) for seq in lines] )
print(sum([seq[0] - extrapolate(seq) for seq in lines]))
