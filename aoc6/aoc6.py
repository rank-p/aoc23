with open("aoc6.txt") as f:
    time = [int(x.strip()) for x in f.readline().split(":")[1].strip().split(" ") if x!= ""]
    distance = [int(x.strip()) for x in f.readline().split(":")[1].strip().split(" ") if x!= ""]

score = 1 
for i in range(len(time)):
    nr_of_wins = 0
    for y in range(1,time[i]):
        if y * (time[i]-y) > distance[i]:
            nr_of_wins += 1
    score *= nr_of_wins
    
print(score)
