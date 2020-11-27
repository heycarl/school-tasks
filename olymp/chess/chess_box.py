f = open("input.txt", "r")
j = f.read().split('\n')[1:]
f.close()

n = 0

black_complect = [0,0,0,0,0,0]
white_complect = [0,0,0,0,0,0]

white_set = [0,0,0,0,0,0] # king 1 ferz 1 ladiya 2 slon 2 kon 2 peshka 8
black_set = [0,0,0,0,0,0] # king 1 ferz 1 ladiya 2 slon 2 kon 2 peshka 8

for figure in j:
    if figure == "white king":
        white_set[0] += 1
        continue
    if figure == "white queen":
        white_set[1] += 1
        continue
    if figure == "white rook":
        white_set[2] += 1
        continue
    if figure == "white bishop":
        white_set[3] += 1
        continue
    if figure == "white knight":
        white_set[4] += 1
        continue
    if figure == "white pawn":
        white_set[5] += 1
        continue

    if figure == "black king":
        black_set[0] += 1
        continue
    if figure == "black queen":
        black_set[1] += 1
        continue
    if figure == "black rook":
        black_set[2] += 1
        continue
    if figure == "black bishop":
        black_set[3] += 1
        continue
    if figure == "black knight":
        black_set[4] += 1
        continue
    if figure == "black pawn":
        black_set[5] += 1
        continue


for i in range(6):
    if i == 0 or i == 1:
        white_complect[i] = white_set[i] // 1 
        black_complect[i] = black_set[i] // 1
    if i == 2 or i == 3 or i == 4:
        white_complect[i] = white_set[i] // 2
        black_complect[i] = black_set[i] // 2
    if i == 5:
        white_complect[i] = white_set[i] // 8
        black_complect[i] = black_set[i] // 8

comlect = white_complect+black_complect
comlect.sort()

o = open("output.txt", "w")
o.write(str(comlect[0]))
o.close()