import sys
import re

def main():
    text = open(sys.argv[1]).read().split("\n")
    print(count_red_green_blue(text))

def count_red_green_blue(text):   
    p1 = 0
    p2 = 0
    for line in text:
        idx = text.index(line) + 1
        red = re.findall(r'[0-9]+ red', line)
        blue = re.findall(r'[0-9]+ blue', line)
        green = re.findall(r'[0-9]+ green', line)   
        r = " ".join(red)
        b = " ".join(blue)
        g = " ".join(green)
        red_dig = re.findall(r'[0-9]+', r)
        blue_dig = re.findall(r'[0-9]+', b)
        green_dig = re.findall(r'[0-9]+', g)
        red_num = [int(i) for i in red_dig] 
        blue_num = [int(i) for i in blue_dig]
        green_num = [int(i) for i in green_dig]
        if max(red_num) <= 12 and max(blue_num) <= 14 and max(green_num) <= 13:
            p1 += idx
        power = max(red_num) * max(blue_num) * max(green_num)
        p2 += power
    return p1, p2


