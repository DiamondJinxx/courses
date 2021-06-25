#!/usr/bin/python3

def get_compress_genom(genom):
    index = 0
    count = 0
    ret_genom = ''
    while index !=len(genom):
        g = genom[index]
        while (index != len(genom)) and (genom[index] == g):
            count += 1
            index += 1
        ret_genom += g + str(count)
        count = 0
    return ret_genom

string = input()
new_str = get_compress_genom(string)
print(new_str)
