with open("aoc6.txt") as f:
    time = int("".join([x for x in f.readline().split(":")[1].strip().split(" ") if x!= ""]))
    distance = int("".join([x for x in f.readline().split(":")[1].strip().split(" ") if x!= ""]))

nr_of_wins = 0
for i in range(1, time):
    if i * (time-i) > distance:
        nr_of_wins += 1
    
print(nr_of_wins)
