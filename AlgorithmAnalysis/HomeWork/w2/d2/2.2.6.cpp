#include<bits/stdc++.h>

using namespace std;

struct Node {
    int Data;
    Node *Next;
    void Add(Node *p);
    int IncSet();
    //Node* Plus(Node *lhs, Node *rhs, Node *ans);
    Node() = default;
    Node(int n) {
        Data = n;
        Next = NULL;
    }
};

void Node::Add(Node *p) {
    p->Next = this->Next;
    this->Next = p;
}

int Node::IncSet() {
    if(this->Next == NULL) {
        int c = this->Data / 10;
        this->Data %= 10;
        return c;
    }
    this->Data += this->Next->IncSet();
    int c = this->Data / 10;
    this->Data %= 10;
    return c;
}

int main() {
    Node *A = new Node(), *B = new Node();
    string sa, sb;
    int la, lb;
    
    cin >> sa >> sb;
    la = sa.size();
    lb = sb.size();

    Node *p = A;
    for(int i=0;i<la;++i) {
        p->Add(new Node(sa[i] - '0'));
        p = p->Next;
    }
    p = B;
    for(int i=0;i<lb;++i) {
        p->Add(new Node(sb[i] - '0'));
        p = p->Next;
    }

    if(la < lb) {
        Node *t = A;
        A = B;
        B = t;
        la = sb.size();
        lb = sa.size();
    }

    Node *Ans = new Node();
    p = Ans;
    Node *q = A;
    for(int i=0;i<la - lb;++i) {
        p->Add(new Node(q->Next->Data));
        p = p->Next;
        q = q->Next;
    }
    Node *w = B;
    for(int i=0;i<lb;++i) {
        //printf("%d %d\n", q->Next->Data, w->Next->Data);
        p->Add(new Node(q->Next->Data + w->Next->Data));
        p = p->Next;
        q = q->Next;
        w = w->Next;
    }
    
    int c = Ans->Next->IncSet();
    if(c) {
        Ans->Add(new Node(c));
    }
    
    p = Ans;
    if(p->Next != NULL) {
        printf("%d", p->Next->Data);
    }
    p = p->Next;
    while(p->Next != NULL) {
        printf(" -> %d", p->Next->Data);
        p = p->Next;
    }
    printf("\n");
    return 0;
}

