def fib_mod(n, m):
    F = [0]*3
    F[0] = 0
    F[1] = 1
    for i in range(2,n+1):
        F[2] = F[1] + F[0]
        F[0] = F[1]
        F[1] = F[2]
        print(F[2])
    return F[2] % m


def main():
    n, m = map(int, input().split())
    print(fib_mod(n, m))


if __name__ == "__main__":
    main()
