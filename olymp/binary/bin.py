f = open("input.txt", "r")
nums = f.read().split('\n')[1:][0].split(" ")
f.close()

nums = [bin(int(i))[2:] for i in nums]
final = []

def check_if_odinakov(s):
    sk = s[0]
    for sym in s:
        if sym == sk:
            pass
        else:
            return 0
    return 1

for bin_num in nums:
    a = list(bin_num)
    if check_if_odinakov(a):
        a[len(a)-1] = str(int(not bool(str(a[len(a)-1]))))
        final.append("".join(a))  
        continue
    first = a[0]
    flag = False
    for sym in a:
        if sym != first and flag == False:
            flag == True
            if bool(int(sym)) == True:
                sym = '0'
            else:
                a[a.index(sym)] = '1'
            break
    final.append("".join(a))   
         
final = [int('0b' + str(num), 2) for num in final]
o = open("output.txt", "w")
o.write(" ".join([str(i) for i in final]))
o.close()