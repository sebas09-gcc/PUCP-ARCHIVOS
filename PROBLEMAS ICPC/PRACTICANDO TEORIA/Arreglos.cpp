#include <bits/stdc++.h>
using namespace std;

int sumarCifras(int x){
	int paso,i=0,acu=0;
	while(1){
		paso=x%10;
		x=x/10;
		acu=paso+acu;
		if(x==0)break;
		i++;
	}
	return acu;
}

int digital_root(int x){
	int cifras;
	if(x<10)return x;
	cifras=sumarCifras(x);
	digital_root(cifras);
}

int main(){
    int n,k,x,contador;
    long long int cal,numero,x_supuesto;
    cin>>n;
    for(int i=0;i<n;i++){
    	cin>>k;
    	cin>>x;
    	contador=0;
    	numero=0;
    	while(1){
			x_supuesto=digital_root(numero);
			if(x_supuesto==x)contador++;
			if(contador==k){
    			break;
			}
    		numero++;
		}
		cout<<numero<<endl;
	}
}
