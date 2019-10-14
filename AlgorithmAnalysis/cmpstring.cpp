
#include<bits/stdc++.h>

using namespace std;

string s1, s2;
typedef uint unsigned int;

uint hash(string &s) {
    uint k = 1;
    uint ans = 0;
    const int MOD = 1000000 + 7;//素数
    for (int i=s.size() - 1;i>=0;--i) {
        ans += k * i;
        k *= 26;
    }
    return ans % 1000007;
}
0~1000006
O(n) 可能冲突
s0 != s1
h(s0) = h(s1)
set<string> ss;//每个元素最多出现一次
O(nlogn) 不会冲突
ss.insert(s);
ss.count(s); 1 / 0 //s在ss中出现次数

map<int, string>