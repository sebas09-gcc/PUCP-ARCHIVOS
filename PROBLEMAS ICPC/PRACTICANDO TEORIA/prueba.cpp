int sumarCifras(int x){
	int paso,acu=0;
	while(1){
		paso=x%10;
		x=x/10;
		acu=paso+acu;
		if(x==0)break;
	}
	return acu;
}

#include <bits/stdc++.h>
using namespace std;

int main(){
	int x=156844,cifras;
	cifras=sumarCifras(x);
	cout<<cifras<<endl;
}
