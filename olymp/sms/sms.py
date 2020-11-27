# mode = int(input("Enter mode: "))
# string = input("Etner string: ")
mode = 2
string = "an" # 3+3+4

current_key = [-1, -1]
current_caps_state = False
output = 0

wordlist = [
    {
    '.': 1,
    ',': 2,
    '1': 3,
    },
    {
    'a': 1,
    'b': 2,
    'c': 3,
    '2': 4,
    },
    {
    'd': 1,
    'e': 2,
    'f': 3,
    '3': 4,
    },
    {
    'g': 1,
    'h': 2,
    'i': 3,
    '4': 4,
    },
    {
    'j': 1,
    'k': 2,
    'l': 3,
    '5': 4,
    },
    {
    'm': 1,
    'n': 2,
    'o': 3,
    '6': 4,
    },
    {
    'p': 1,
    'q': 2,
    'r': 3,
    's': 4,
    '7': 5,
    },
    {
    't': 1,
    'u': 2,
    'v': 8,
    '8': 4,
    },
    {
    'w': 1,
    'x': 2,
    'y': 3,
    'z': 4,
    '9': 5,
    },
    {
    '+': 1,
    '0': 2,
    ' ': 1
    }
]

if mode == 1:
    current_caps_state = False
else:
    current_caps_state = True


for s in string:
    print(s.isupper(), current_caps_state)
    if  s.isupper() != current_caps_state:
        current_caps_state = not current_caps_state
        output += 3
        for i in range(len(wordlist)):
                if wordlist[i].get(s.lower()) != None:
                    current_key[1] = i
                    output += wordlist[i].get(s.lower())
        if current_key[0] != current_key[1]:
            output += 2
            current_key[0] = current_key[1]
        else:
            output += 3
    else:
        for i in range(len(wordlist)):
                if wordlist[i].get(s.lower()) != None:
                    current_key[1] = i
                    output += wordlist[i].get(s.lower())
        if current_key[0] != current_key[1]:
            output += 2
            current_key[0] = current_key[1]
        else:
            output += 3
    print("Out: ", output)

print("Total out: ", output)
        
def if_caps():
    return