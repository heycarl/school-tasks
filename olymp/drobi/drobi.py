import math
i = 0
drob = float(input())

znam_max = int(input())

for chislitel in range(1, znam_max+1):
    for znamenatel in range(1, chislitel):
        if math.gcd(chislitel, znamenatel) == 1 and znamenatel / chislitel <= drob:
            i += 1

print(i)