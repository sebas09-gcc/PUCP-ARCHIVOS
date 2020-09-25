#include<iostream>
using namespace std;

int main()
{
	long long int n,k,x;
	cin>>n;
	for(int i=0;i<n;i++){
		cin>>k>>x;
		cout<<9*(k-1)+x<<endl;
	}
	return 0;
}
