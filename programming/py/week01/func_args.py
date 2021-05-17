

def print_ab(a=10,b=10):
	print(a)
	print(b)


print_ab(5)
print_ab(b=10,a=20)

lst = [10,20]

print_ab(*lst)

args = {'a':20, 'b':10} #развернется в key=value и получится print_ab(a=10,b=20)
print_ab(**args)


print_ab()