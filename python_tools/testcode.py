import re
import os
import binascii
import binhex

bestregexyet = "\t(?:[0-9a-f]{2} *){1,7}\t"
start_of_main = "[0-9]*<main>:"
hexstring = []
asdf = []
shellcode = []
objdump_input = open("objdump.txt", "r")
for line_of_text in objdump_input:
    if line_of_text != None:
        hex_match = re.search(bestregexyet,line_of_text, re.I)
        if hex_match != None:
            shellcode.append(hex_match[0].replace("\t","").replace("\n", ""))
            #hexstring = shellcode + str(hex_match[0][0])

for each in shellcode:
    asdf.append(each.strip())#.replace("  ", ""))

for line in asdf:
    for hexval in line.split(" "):
        hexstring.append("/x" + hexval)

#"".join(asdf)#.replace(" ", "/x")
