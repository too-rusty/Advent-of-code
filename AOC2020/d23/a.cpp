#include<iostream>
#include<map>
#include<set>
#include <vector>
using namespace std;

//PART B

struct Node {
    Node* next;
    int val;
    Node(int v):val(v),next(nullptr) {
    }
};

Node* head = nullptr;
map<int,Node*> addr;
const int lim=1000000; //mx element

Node* newNode(int val) {
    Node* x = new Node(val);
    addr[val]=x;
    return x;
}

void insert(int val,Node** node) {
    if (*node == nullptr) {
        *node = newNode(val);
        // addr[val] = node;
        return;
    }
    Node* tmp = (*node)->next;
    (*node)->next = newNode(val);
    (*node)->next->next=tmp;
    // addr[val]=node;
}
void insert_after(int val,int val2) {
    Node* now=addr[val];
    Node* tmp=now->next;
    Node* neww = newNode(val2);
    neww->next=tmp;
    now->next=neww;
}

void make_circular(int val) {
    Node* now=addr[val];
    now->next=head;
}

void traverse(int x){

    Node* tmp = head;

    int cnt=0;
    while(tmp!=nullptr){
        if(cnt==x)break;
        cout<<tmp->val<<" -> ";
        tmp=tmp->next;
          cnt++;
    }
    cout<<endl;
}


int foo(int val) {
    //take next three values
    //link the next after the current
    const int N=3; //total to pick
    
    //choose 3 nodes to unlink
    Node* now=addr[val];
    Node* now2=now->next;
    vector<Node*>v;
    int np=val-1;
    set<int>s;
    for(int i = 0;i<N;i++){
        v.push_back(now2);
        s.insert(now2->val);
        now2=now2->next;
    }
    //decide the position after which the values need to be linked
    while(1){
        if (np==0)np=lim;
        if(s.find(np)!=s.end()){
            --np;
        }
        else break;
    }
    //link the three values
    now->next=now2;
    Node* tmp2=addr[np]->next;
    Node* now3=addr[np];
    for(int i=0;i<N;i++){
        now3->next=v[i];
        now3=now3->next;
    }
    now3->next=tmp2;
    //return the next position to start the process at
    return now2->val;
}

void bef_and_after(int x){
    Node* tmp = head;
    while(1){
        if(tmp->val==x) {
            cout<<tmp->next->val<<" - ";
            cout<<tmp->next->next->val<<"\n";
            cout<<1ll*(tmp->next->val) * (tmp->next->next->val) << endl;
            break;
        }
        tmp = tmp->next;
    } 
}


int main() {

    // vector<int> vv({3,8,9,1,2,5,4,6,7});
    vector<int> vv({5,3,8,9,1,4,7,6,2});
    for(int i=10;i<lim+1;i++)vv.push_back(i);

    insert(5,&head);

    for(int i=0;i<lim-1;i++){
        insert_after(vv[i],vv[i+1]);
    }
    make_circular(lim);

    int st=vv[0]; //START
    for(int i =0;i<10000000;i++){
        st= foo(st);
        if(i%10000==0)cout<<i<<endl;
    }
    bef_and_after(1);

    return 0;
}