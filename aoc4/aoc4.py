from functools import reduce

with open("aoc4.txt", "r") as f:
    score = 0
    score_map = {}
    for line in f.readlines():
        game_id = int(line.split(":")[0].split(" ")[-1].strip())
        scratchcard = line.split(":")[1].strip()
        winning_numbers = [int(x) for x in scratchcard.split(" | ")[0].split(" ") if x != ""]
        numbers = [int(x) for x in scratchcard.split(" | ")[1].split(" ") if x != ""]
        nr_of_wins = 0
        for n in numbers:
            if n in winning_numbers:
                nr_of_wins += 1
        
        score_map[game_id] = nr_of_wins
        
    count_map = {game_id: 1 for game_id in score_map.keys()}
    for game_id, score in score_map.items():
        for i in range(1, score+1):
            try:
                count_map[game_id+i] += count_map[game_id] 
            except:
                pass

    print(reduce(lambda x, key: x + count_map[key], count_map, 0))
