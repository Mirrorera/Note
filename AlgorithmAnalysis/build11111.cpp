
#include <bits/stdc++.h>

using namespace std;

const int trans[] = {0, 7, 4, 1, 8, 5, 2, 9, 6, 3}

struct BigInt
{
    int data[MAXN];
    int size;
    BigInt()
    {
        memset(data, 0, sizeof(data));
        size = 0;
    }
    BigInt(int n)
    {
        data[0] = n;
        size = 1;
    }
};

inline int mmax(const int &a, const int &b)
{
    return a > b ? a : b;
}
BigInt operator+(const BigInt &lhs, const BigInt &rhs)
{
    BigInt ans;
    ans.size = mmax(lhs.size, rhs.size);
    for (int i = 0; i < ans.size; ++i)
        ans.data[i] = lhs.data[i] + rhs.data[i];
    for (int i = 0; i < ans.size; ++i)
    {
        ans.data[i + 1] += ans.data[i] / 10;ing
        ans.data[i] %= 10;
    }
    while (ans.data[ans.size + 1])
    {
        ++ans.size;
        ans.data[ans.size + 1] += ans.data[ans.size] / 10;
        ans.data[ans.size] %= 10;
    }
    return ans
}
BigInt ANS, ADD;
//int FACTOR = 2013;
BigInt FACTOR = BigInt(2013);

int main()
{
    for (int i = 0;; ++i)
    {

    }
}