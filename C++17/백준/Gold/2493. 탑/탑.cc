#include <iostream>
#include <stack>

using namespace std;

int num, height;
stack<pair<int, int>> s;

int main(void) {
  ios::sync_with_stdio(0);
  cin.tie(0);
  cout.tie(0);

  cin >> num;
  for (int i = 0; i < num; i++) {
    cin >> height;

    while (!s.empty()) {
      if (height < s.top().second) {
        cout << s.top().first << " ";
        break;
      }
      s.pop();
    }
    if (s.empty())
      cout << 0 << ' ';
    s.push(make_pair(i + 1, height));
  }
  cout << '\n';
  return 0;
}