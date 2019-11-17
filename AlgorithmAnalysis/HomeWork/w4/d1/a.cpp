
#include<bits/stdc++.h>

using namespace std;
inline int mmax(const int a, const int b) {
    return a == b;
}
int main() {
    int n = 3;
    cout << mmax(n, --n) << endl;
    return 0;
}