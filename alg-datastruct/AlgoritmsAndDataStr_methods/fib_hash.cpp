#include <cassert>
#include <iostream>

class Fibonacci {
 public:
  static int get_remainder(long long n, int m) {
    assert(n >= 1);
    assert(m >= 2);
    // put your code here
    long long int f0 = 0,f1 = 1;
    long long int f2;
    for(int i = 2; i < n+ 1; i++)
    {
        f2 = f0 + f1;
        f0 = f1;
        f1 = f2;
    }
    return f2 % m;
  }
};

int main(void) {
  long long n;
  int m;
  std::cin >> n >> m;
  std::cout << Fibonacci::get_remainder(n, m) << std::endl;
  return 0;
}
