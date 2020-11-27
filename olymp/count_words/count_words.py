f = open("input.txt", "r")
j = f.read().split('\n')
f.close()

n = int(j[0])
words = j[1].split(",")
i = 0

for word in words:
    if len(word) == n:
        i += 1

o = open("output.txt", "w")
o.write(str(i))
o.close()