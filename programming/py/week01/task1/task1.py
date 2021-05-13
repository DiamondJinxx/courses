import math
#this func return a closest integer number mod 5
def closest_mod_5(x):
	r = math.ceil(x / 5)
	return r * 5
