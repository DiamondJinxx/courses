#!/usr/bin/python3

n = int(input())
count = 0

for i in range(1, n + 1):
    if(count == n):
        break
    for j in range(1, i+1):
        if(count != n):
            count += 1
            print(i)
        else:
            break
            

