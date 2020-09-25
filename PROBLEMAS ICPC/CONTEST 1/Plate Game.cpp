#include <bits/stdc++.h>
using namespace std;

int main(){
	long long int a,b,r,diametro,anchopar,largopar;
	cin>>a>>b>>r;
	diametro=2*r;
	if(diametro>a||diametro>b){
		cout<<"Second"<<endl;
	}
	else{
		cout<<"First"<<endl;

	}
	return 0;
}
