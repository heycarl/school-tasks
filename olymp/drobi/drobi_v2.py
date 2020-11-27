import math
i = 0
drob = float(input())
znam_max = int(input())

for znamen in range(1, znam_max+1):
    for chislit in range(1, int(drob*znam_max)+1):
        if math.gcd(chislit, znamen) == 1 and chislit / znamen <= drob:
            i += 1
print(i)
