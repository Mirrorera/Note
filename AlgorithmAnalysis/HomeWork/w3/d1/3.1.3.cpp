
#include<bits/stdc++.h>

using namespace std;

const int MAXN = 100 + 7;

int graph[MAXN][MAXN];
int match[MAXN];
int vis[MAXN];

const char out[] = {'A', 'B', 'C', 'X', 'Y', 'Z'};
bool find(int p) {
    for(int i=0;i<6;++i) {
        if(!graph[p][i])
            continue;
        if(vis[i])
            continue;
        vis[i] = 1;
        if(match[i] == -1 || find(match[i])) {
            match[p] = i;
            match[i] = p;
            return true;
        } else {
            return false;
        }
    }
    return false;
}
int main() {
    memset(match, -1, sizeof(match));
    //cin >> N;
    for(int i=0;i<3;++i)
        for(int j=3;j<6;++j)
            graph[i][j] = graph[j][i] = 1;
    graph[0][3] = graph[3][0] = 0;
    graph[2][3] = graph[3][2] = 0;
    graph[2][5] = graph[5][2] = 0;

    for(int i=0;i<3;i++) {
        memset(vis, 0, sizeof(vis));
        find(i);
    }
    for(int i=0;i<3;++i)
        cout << out[i] << ' ' << out[match[i]] << endl;
    return 0;
}
