f = open("input.txt", "r")
j = f.read().split('\n')
f.close()
nums_unsort = []

#2d to 1d
for row in j[1:]:
    k = row.split(" ")
    for m in k:
        if m != " ":
            nums_unsort.append(int(m))

nums_sorted = nums_unsort.copy()

def get_n_entries(key):
    return nums_unsort.count(key)

nums_sorted.sort(key=get_n_entries)

o = open("output.txt", "w")
o.write(str(nums_sorted[0]) + ' ' + str(nums_sorted.count(nums_sorted[0])))
o.close()