#include <iostream>
using namespace std;

#define ll long long

ll mini, maxi;
bool val[1000001];
int num;

int main(void) {
  ios::sync_with_stdio(0);
  cin.tie(0);
  cout.tie(0);

  cin >> mini >> maxi;
  for (ll i = 2; i * i <= maxi; i++) {
    ll cur = i * i;
    ll sum = mini - (mini % cur);
    while(sum <= maxi) {
      if(sum >= mini) {
        val[sum - mini] = true;
      }
      sum += cur;
    }
  }
  for (int i = 0; i <= maxi - mini; i++) {
    if (val[i] == false) num++;
  }
  cout << num << '\n';
  

}