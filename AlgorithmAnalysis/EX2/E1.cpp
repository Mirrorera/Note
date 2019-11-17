
#include <bits/stdc++.h>

using namespace std;

const int MAXN = 1000000 + 233;//1e7

int N;
int pos[MAXN];

int quickQueryMid(int ll, int rr) {
    int l = ll;
    int r = rr - 1;
    int t = pos[ll];

//与快速排序核心操作类似，使t左边的数均比其小，右边的数均比其大
    while(l < r) {
        while(l < r && pos[r] <= t) {
            --r;
        }
        if(l == r) {
            break;
        }
        //替代swap(pos[l], pos[r]), 二者之中实际上一直有一个为t
        pos[l] = pos[r];
        pos[r] = t;

        while(l < r && pos[l] >= t) {
            ++l;
        }
        if(l == r) {
            break;
        }
        //替代swap(pos[l], pos[r]), 二者之中实际上一直有一个为t
        pos[r] = pos[l];
        pos[l] = t;
    }
    
    //缩小问题可能规模， 二分的寻找中位数
    if (l == N / 2) {
        return pos[l];
    } else {
        if (l < N / 2) {
            return quickQueryMid(l + 1, rr);
        } else {
            return quickQueryMid(ll, r);
        }
    }
}

int main() {
    scanf("%d", &N);
    for(int i=0;i<N;++i) {
        scanf("%d", &pos[i]);
    }
    int pipe = quickQueryMid(0, N);
    int ans(0);
    for(int i=0;i<N;++i) {
        ans += abs(pos[i] - pipe);
    }
    printf("TOTAL LENGTH is %d When Pipe's function is (y = %d)\n", ans, pipe);
    return 0;
}
/*
4 3 7 8 2
^       ^
l = r r = t
8 3 7 4 2
^     ^

8 4 7 3 2
  ^   ^

8 7 4 3 2
  ^ ^
*/
/*
- 原问题等价于求解对于一个整数数组a[i], 寻找一个x, 令$\sum_i |a[i] - x|$最小
- 有推论(*), 对于$a[k] \leq x \leq a[k + 1], Ans(x) = Const$
- 问题规约到 求$max_j\{\sum_i |a[i] - a[j]|\}$
- 且$\sum_i |a[i] - x|\ = (|a[0] - y| + |a[n - 1] - y|) + (|a[1] - y| + a[n - 2] - y) + ...$
- 可知管道的选择应该让两边的油井数量尽量相近, 再结合推论(*), 可知管道应通过中位数位置,或中位数位置两侧的任一个
- 求中位数算法时间复杂度为O(n)
    - 核心是快速排序中的子操作, 对于一个数t, 可以O(n)时间内, 使t左边的数均比其小，右边的数均比其大
    - 每次这样的操作后, 问题规模平均缩小一半, 计算次数f(n) = n * (1 + 1 / 2 + 1 / 4 + ...) < 2n
    - 所以可以在O(n)复杂度下求解管道位置
- 最后计算结果
*/