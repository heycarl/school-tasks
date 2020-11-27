f = open("input.txt", "r")
arr = f.read().split("\n")
n = int(arr[0])
passwords = arr[1:]

correct_sym = "qwertyuiopasdfghjkzxcvbnm98765432QWERTYUPASDFGHJKLZXCVBNM" # no O0Il1


def check_wrong_sym(pwd):
    for sym in pwd:
        if correct_sym.find(sym) == -1:
            return 0
    return 1


valid_passwords = passwords[:]


for password in passwords:
    # length check
    if len(password) != 8:
        valid_passwords.pop(valid_passwords.index(password))
        continue
    if check_wrong_sym(password) == 0:
        valid_passwords.pop(valid_passwords.index(password))
        continue

fi = open("output.txt", "w")
out = fi.write(str(len(valid_passwords)) + " " + str(n - len(valid_passwords)))
# print(valid_passwords)
f.close()
fi.close()