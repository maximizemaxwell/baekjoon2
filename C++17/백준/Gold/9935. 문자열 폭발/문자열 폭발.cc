#include <iostream>
using namespace std;

int main(void) {
  ios::sync_with_stdio(0);
  cin.tie(0);
  cout.tie(0);

  string str, boom, temp = "";
  cin >> str;
  cin >> boom;

  for (int i = 0; i < str.length(); i++) {
    temp += str[i];
    if (temp.back() == boom.back()) {
      bool explosion = true;
      if (temp.length() < boom.length())
        continue;
      for (int j = 0; j < boom.length(); j++) {
        if (temp[temp.length() - boom.length() + j] != boom[j]) {
          explosion = false;
          break;
        }
      }
      if (explosion) {
        for (int j = 0; j < boom.length(); j++)
          temp.pop_back();
      }
    }
  }
  if (temp.empty())
    cout << "FRULA" << '\n';
  else
    cout << temp << '\n';
  return 0;
}