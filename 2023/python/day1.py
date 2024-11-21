import re
import sys
def main():
    file = open(sys.argv[1]).read().strip().split("\n")
    dict_nums = {"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
    dict_rev_nums = {"enin": 9, "thgie": 8, "neves": 7, "xis": 6, "evif": 5, "ruof": 4, "eerht": 3, "owt": 2 , "eno": 1}
    
    solution = day_1_solution(file, dict_nums, dict_rev_nums)
    print(solution[0])
    print(solution[1])
def day_1_solution(file, dict_nums, dict_rev_nums):
    count_p1 = 0
    count_p2 = 0
    for line in file:
        p1_list = []
        p2_list = []
        rev_p2_list = []
        rev_line = line[::-1] 
        regex = re.findall(r"([0-9]|one|two|three|four|five|six|seven|eight|nine)", line) 
        rev_regex = re.findall(r"([0-9]|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno)", rev_line)
        for item in regex:
            if item in dict_nums.keys():
                p2_list.append(dict_nums[item])
            elif item.isdigit():
                p1_list.append(item) 
                p2_list.append(int(item))
        for item in rev_regex:
            if item in dict_rev_nums.keys():
                rev_p2_list.append(dict_rev_nums[item])
            elif item.isdigit():
                rev_p2_list.append(int(item))
        count_p1 += int(p1_list[0] + p1_list[-1])
        count_p2 += p2_list[0] * 10 + rev_p2_list[0]
    return count_p1, count_p2
            
main()