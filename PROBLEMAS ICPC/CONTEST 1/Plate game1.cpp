#include<bits/stdc++.h>
using namespace std;
long long int a,b,r;

int main(){
	cin>>a;
	cin>>b;
	cin>>r;
	cout<<r*2;
	if(a<r*2 || b<r*2){
		cout<<"Second"<<endl;
	}
	else{
		cout<<"First"<<endl;
	}
	return 0;
}
