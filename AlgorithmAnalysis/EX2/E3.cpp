
#include<bits/stdc++.h>

using namespace std;

const int MAXN = 10000 + 233;

vector<int> graph[MAXN];
bool incircle[MAXN];

int N;

bool vis[MAXN];
int dp[MAXN][2];

inline int mmax(const int a, const int b) {
    return a > b ? a : b;
}
int findCircle(int pre, int p) {
    vis[p] = true;
    for (vector<int>::iterator nex = graph[p].begin();nex < graph[p].end();++nex) {
        if (*nex == pre) {
            continue;
        }
        if (vis[*nex]) {
            incircle[p] = true;
            return *nex;
        }
        int t;
        //vis[p] = true;
        t = findCircle(p, *nex);
        if (t >= 0) {
            incircle[p] = true;
            if(t == p) {
                return -2;//已找到环
            } else {
                return t;
            }
        } else {
            if (t == -2) {
                return -2;
            }
        }
    }
    //vis[p] = false;
    return -1;
}

int clu(int root, int prestat) {
    if (dp[root][prestat]) {
        return dp[root][prestat];
    }
    int tot(0);
    vis[root] = true;
    for (vector<int>::iterator ch = graph[root].begin();ch < graph[root].end();++ch) {
        if (vis[*ch]) {
            continue;
        }
        tot += prestat ? clu(*ch, 1 ^ prestat) :
                         mmax(clu(*ch, prestat), clu(*ch, 1 ^ prestat) + 1);
    }
    return dp[root][prestat] = tot;
}
int main() {
    memset(incircle, 0, sizeof(incircle));
    scanf("%d", &N);
    for(int i=0;i<N;++i) {
        int u, v;
        scanf("%d %d", &u, &v);
        graph[u].push_back(v);
        graph[v].push_back(u);
    }

    memset(vis, 0, sizeof(vis));    
    findCircle(-1, 0);
    
    int root1;
    root1 = -1;
    while(!incircle[++root1]);
    vector<int>::iterator root2 = graph[root1].end();
    while(!incircle[*(--root2)]);

    int ans1, ans2;
    memset(vis, 0, sizeof(vis));
    memset(dp, 0, sizeof(dp));
    ans1 = clu(root1, 1);
    memset(vis, 0, sizeof(vis));
    memset(dp, 0, sizeof(dp));
    ans2 = clu(*root2, 1);
    
    printf("ANS = %d\n", mmax(ans1, ans2));
    return 0;
}