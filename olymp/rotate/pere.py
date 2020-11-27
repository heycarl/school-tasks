f = open("input.txt", "r")
j = f.read().split('\n')
f.close()

words = []

for i in range(1, int(j[0])+1):
    words.append(j[i])

translate = ["W9M6", "M6W9"]
equals = "HINOSXZ0"
i = 0
new_words = []

for word in words:
    new = ""
    for letter in word:
        if translate[0].find(letter) != -1:
            new += translate[1][translate[0].find(letter)]
            continue
        if equals.find(letter) != -1:
            new += letter
            continue
        new = ""
        break
    if new != "":
        i += 1
        new_words.append(new)

o = open("output.txt", "w")
o.write(str(i))

for word in new_words:
    o.write("\n")
    o.write(word[::-1])
o.close()