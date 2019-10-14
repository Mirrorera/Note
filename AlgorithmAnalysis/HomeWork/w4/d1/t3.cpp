
#include<bits/stdc++.h>

using namespace std;

const int MAXN = 1000 + 233;

int N;
int data[MAXN][MAXN];
int aim;

bool sqrSearch(int lr, int lc, int rr, int rc) {
    bool exist = false;
    if (lr < 0 || lr >= N || lc < 0 || lc > N || rr < 0 || rr >= N || rc < 0 || rc >= N) {
		return false;
	}
	if (lr > rr || lc > rc) {
		return false;
	}
	if (lr == rr && lc == rc) {
		if (data[lr][lc] == aim) {
			return true;
		} else {
			return false;
		}
	}
	int mr = (lr + rr) >> 1;
	int mc = (lc + rc) >> 1;
	int d = data[mr][mc];
	if (aim < d) {
		exist = sqrSearch(lr, lc, mr-1, mc-1);
		if (exist) {
			return exist;
		}
		exist = sqrSearch(mr, lc, rr, mc-1);
		if (exist) {
			return exist;
		}
		exist = sqrSearch(lr, mc, mr-1, rc);
		if (exist) {
			return exist;
		}
	} else if (aim > d) {
		exist = sqrSearch(mr+1, mc+1, rr, rc);
		if (exist) {
			return exist;
		}
		exist = sqrSearch(mr+1, lc, rr, mc);
		if (exist) {
			return exist;
		}
		exist = sqrSearch(lr, mc+1, mr, rc);
		if (exist) {
			return exist;
		}
	} else {
		return true;
	}
	return false;
}

int main() {
    cin >> N;
    for(int i=0;i<N;++i)
        for(int j=0;j<N;++j)
            cin >> data[i][j];
    cin >> aim;
    bool ans = sqrSearch(0, 0, N-1, N-1);
    if (ans)
        cout << "Exist" << endl;
    else
        cout << "Not Exist" << endl;
    return 0;
}