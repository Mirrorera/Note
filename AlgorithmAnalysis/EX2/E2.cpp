#include <bits/stdc++.h>

using namespace std;

const int MAXN = 10000 + 233;
int N;
int data[MAXN][MAXN];

int NextMat(int n) {
    int ANS(0);
    /*if (n == 1) {
        return data[0][0];
    }*/
    int begr, begc;
    if (n%3) {
        begr = begc = 0;
    } else {
        begr = begc = 1;
    }
    for(int i=begr;i<n;i+=3) {
        for(int j=begc;j<n;j+=3) {
            ANS += data[i][j];
        }
    }
    return ANS;
}
int main() {
    scanf("%d", &N);
    for(int i=0;i<N;++i) {
        for(int j=0;j<N;++j) {
            scanf("%d", &data[i][j]);
        }
    }
    printf("%d\n", NextMat(N));
    return 0;
}
/*
- 为了将矩阵中地雷总数求出, 将矩阵划分为不相交的3*3子矩阵, 每个子矩阵中心的值就是这个子矩阵中的地雷数。
- 考虑到当边长不是3的倍数时, 不能做到均分为3*3子矩阵, 可以为其加上一圈数值和地雷均为0的格子, 使新矩阵边长为3的倍数的同时, 分割出的子矩阵中心值均在原矩阵中出现
- 实际构造过程不需要添加格子, 只需要在矩阵边长不能被3整除时, 统计时从矩阵(0, 0)位置开始,能被整除时, 统计时从矩阵(1, 1)位置开始, 即可统一三种情况
*/