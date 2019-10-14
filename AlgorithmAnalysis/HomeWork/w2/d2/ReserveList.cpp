#include<bits/stdc++.h>
using namespace std;
struct Node{
    int data;
    Node *Next;
    void Ins(int n);
    void Del();
    Node(int n) {
        data = n;
        Next = NULL;
    }
};

void Node::Ins(int n) {
    Node *it = new Node(n);
    it->Next = this->Next;
    this->Next = it;
}
void Node::Del() {
    Node *it = this->Next;
    this->Next = this->Next->Next;
    delete it;
}

int N, K;
Node *head;

pair<Node*, Node*> reserve(Node *oper, int n) {
    if(n == K) {
        return make_pair(oper->Next, oper);
    }
    pair<Node*, Node*> rt = reserve(oper->Next, n + 1);
    oper->Next->Next = oper;
    return rt;
}

int Reserve(Node *oper) {
    if(oper->Next == NULL) {
        return K - 1;
    }
    int t = Reserve(oper->Next);
    if(t) {
        return
            t - 1;
    } else {
        pair<Node*, Node*> rt;
        rt = reserve(oper->Next, 1);
        oper->Next->Next = rt.first;
        oper->Next = rt.second;
        return K - 1;
    }
    return -1;
}


int main() {
    head = new Node(-1);
    scanf("%d %d", &N, &K);

    Node *p = head;
    for(int i=0;i<N;++i) {
        int t;
        scanf("%d", &t);
        p->Ins(t);
        p = p->Next;
    }
    
    int left = Reserve(head->Next);
    if(!left) {
        pair<Node*, Node*> rt;
        rt = reserve(head->Next, 1);
        head->Next->Next = rt.first;
        head->Next = rt.second;
    }
    p = head->Next;
    for(int i=0;i<N;++i) {
        if(p->Next == NULL)
            break;
        printf("%d -> ", p->data);
        p = p->Next;
    }
    printf("%d\n", p->data);
    return 0;
}
