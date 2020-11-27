f = open("input.txt", "r")
j = f.read().split('\n')
f.close()

n = int(j[0])
source_text = j[1]
source_text = [source_text[i:i+n] for i in range(0, len(source_text), n)]
output_text = ""

for fragment in source_text:
    # a = reverse_fragment(fragment)
    output_text += fragment[::-1]
        
o = open("output.txt", "w")
o.write(output_text)
o.close()