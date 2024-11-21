import sys

file = open(sys.argv[1]).read().strip().split("\n")
points = 0
stack = {}
matches = {}

for i, line in enumerate(file, start=1):
    game, sep, cards = line.partition(":")
    result, sep, score = cards.partition("|")
    a = score.split(" ")
    b = result.split(" ")
    a = set([int(x) for x in a if x != ""])
    b = set([int(x) for x in b if x != ""])
    winning = a & b
    matches[i] = len(winning)
    stack[i] = 1
    if len(winning) != 0:
        points += 2 ** (len(winning) - 1)
             
for k in matches.keys():
        for m in range(k + 1, k + 1 + matches[k]):
            stack.update({m: stack[m] + stack[k]})
        
print(points)
print(sum(stack.values()))


 