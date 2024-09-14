#include <algorithm>
#include <iostream>
#include <map>
#include <string>
#include <vector>
using namespace std;

int main(void) {
  ios::sync_with_stdio(0);
  cin.tie(0);
  cout.tie(0);

  int n, m;
  int cnt = 0;
  cin >> n >> m;
  string s;
  vector<string> result;
  map<string, bool> list;
  for (int i = 0; i < n; i++) {
    cin >> s;
    list.insert(make_pair(s, true));
  }
  for (int i = 0; i < m; i++) {
    cin >> s;
    if (list[s]) {
      result.push_back(s);
      cnt++;
    }
  }
  cout << cnt << '\n';
  sort(result.begin(), result.end());
  for (int i = 0; i < result.size(); i++) {
    cout << result[i] << '\n';
  }
  return 0;
}