#include <iostream>
#include <string>

using namespace std;


long long getHash(string s)
{
	const int k = 31, mod = 1e9+7;
	long long h = 0, m = 1; 
	for(char c : s)
	{
		int x = (int)(c - 'a' + 1);
		h = (h + m * x) % mod;
		m = (m * k) % mod;
	}
	return h;
}



int main(int argc, char const *argv[])
{
	/* code */
	string s = "dima";
	cout << s[0] - 'a' + 1 << endl;
	string s1 = "dima";
	cout << "Hash of str = " << getHash(s) << endl;
	cout << "Hash of str = " << getHash(s1) << endl;
	return 0;
}