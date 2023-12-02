import re

MAP = {
    "one": 1, "two":2, "three":3, "four": 4, "five": 5, "six":6, "seven": 7, "eight": 8, "nine": 9
        } 

pattern_start = r"(?:one|two|three|four|five|six|seven|eight|nine|\d)"
pattern_end = r"(?:one|two|three|four|five|six|seven|eight|nine|\d)(?!.*(?:one|two|three|four|five|six|seven|eight|nine|\d))"

def parse_line(line: str) -> int:
    digit1 = re.search(pattern_start, line.strip()).group(0)
    digit2 = re.search(pattern_end, line.strip()).group(0)
    return (MAP.get(digit1) or int(digit1)) * 10 + (MAP.get(digit2) or int(digit2))


answer=0
lines=[]
with open("aoc1_input.txt", "r") as file:
    for line in file:
        answer += parse_line(line.strip())

print(answer)
