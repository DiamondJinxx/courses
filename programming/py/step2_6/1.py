#!/usr/bin/python3


run = True
summ = 0
quad = 0
while run:
    c = int(input())
    summ += c
    quad += c ** 2
    if summ == 0:
        run = False
print(quad)
