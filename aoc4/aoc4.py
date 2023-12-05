with open("aoc4.txt", "r") as f:
    score = 0
    for line in f.readlines():
        scratchcard = line.split(":")[1].strip()
        winning_numbers = [int(x) for x in scratchcard.split(" | ")[0].split(" ") if x != ""]
        numbers = [int(x) for x in scratchcard.split(" | ")[1].split(" ") if x != ""]
        nr_of_wins = 0
        for n in numbers:
            if n in winning_numbers:
                nr_of_wins += 1
        
        if nr_of_wins > 0:
            score += 2**(nr_of_wins-1)

    print(score)
