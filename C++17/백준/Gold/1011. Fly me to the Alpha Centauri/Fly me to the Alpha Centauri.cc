#include <cmath>
#include <iostream>

using namespace std;

int main(void) {
  ios::sync_with_stdio(0);
  cin.tie(0);
  cout.tie(0);
  int t;
  long long x, y;
  long long move, max;
  cin >> t;
  for (int i = 0; i < t; i++) {
    cin >> x >> y;
    move = 0;
    max = 0;
    while (max * max <= y - x)
      max++;
    max--;
    move = 2 * max - 1;
    long long rem = y - x - max * max;
    rem = (long long)ceil((double)rem / (double)max);
    move += rem;
    cout << move << '\n';
  }
}