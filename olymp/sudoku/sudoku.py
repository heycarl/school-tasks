f = open("input.txt", "r")
source_array = f.read().split("\n")
n = int(source_array[0])
arrays = []
answers = []
avaliable_syms = "987654321"

def check_row_unieq(row):
    used_syms = ""
    for sym in row:
        if used_syms.find(sym) == -1:
            used_syms += sym
        else:
            return 0
    return 1

def check_colomn_unieq(sol, index):
    used_syms = ""
    for i in range(9):
        if used_syms.find(sol[i][index]) == -1:
            used_syms += sol[i][index]
        else:
            return 0
    return 1


solution = source_array[10:10+9*n]
condition = source_array[1:10]
print(condition)
for i in range(n):
    arrays.append(solution[i*9:(i+1)*9])

for _ in range(n):
    answers.append(True)

for solution in arrays:
    for row in solution:
        if check_row_unieq(row) == 0:
            answers[arrays.index(solution)] = False

for solution in arrays:
    for row in range(len(solution)):
        for letter in range(len(solution[row])):
            if avaliable_syms.find(solution[row][letter]) == -1 :
                answers[arrays.index(solution)] = False
            if condition[row][letter] != "*" and solution[row][letter] != condition[row][letter]:
                answers[arrays.index(solution)] = False

for solution in arrays:
    for i in range(9):
        if check_colomn_unieq(solution, i) == 0:
            answers[arrays.index(solution)] = False

for solution in arrays:
    # each row
    for m in range(3):
        # each colomn
        for i in range(3):
            # each colomn in 3x3
            socket = ""
            for j in range(3):
                # each row in 3x3
                for k in range(3):
                    socket += solution[m*3+j][3*i+k]
            for letter in socket:
                if socket.count(letter) > 1:
                    answers[arrays.index(solution)] = False

fi = open("output.txt", "w")
for answer in answers:
    if answers.index(answer) > n:
        break
    if answer == True:
        fi.write("Y")
    else: 
        fi.write("N")
f.close()
fi.close()