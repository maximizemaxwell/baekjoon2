#include <cmath>
#include <cstring>
#include <iostream>
#include <string>

using namespace std;

int N;
int cnt[10];

int LastIsNine(int num) {
  while (1) {
    if (num % 10 == 9)
      break;
    num--;
  }
  return num;
}
void A2B(int A, int B, int digit) {
  if (B < 10) {
    for (int i = A; i <= B; i++)
      cnt[i] += digit;
    return;
  }
  int temp = LastIsNine(B);
  for (int i = temp + 1; i <= B; i++) {
    string str = to_string(i);
    for (int j = 0; j < str.length(); j++) {
      int idx = str[j] - '0';
      cnt[idx] += digit;
    }
  }
  for (int i = 0; i < 10; i++) {
    cnt[i] += ((temp / 10 - A / 10 + 1) * digit);
  }
  A2B(A / 10, temp / 10, digit * 10);
  return;
}
int main(void) {
  ios::sync_with_stdio(0);
  cin.tie(0);
  cout.tie(0);
  cin >> N;
  memset(cnt, 0, sizeof(cnt));
  if (N < 10) {
    for (int i = 1; i <= N; i++)
      cnt[i] += 1;
    for (int i = 0; i < 10; i++)
      cout << cnt[i] << ' ';
    cout << '\n';
    return 0;
  }
  int len = to_string(N).size();
  int A = (int)pow(10, len - 1);
  if (A > LastIsNine(N)) {
    for (int i = LastIsNine(N) + 1; i <= N; i++) {
      string s = to_string(i);
      for (int j = 0; j < s.size(); j++) {
        int idx = s[j] - '0';
        cnt[idx] += 1;
      }
    }
    A /= 10;
    A2B(A, LastIsNine(N), 1);
    while (1) {
      if (A == 1)
        break;
      A2B(A / 10, A - 1, 1);
      A /= 10;
    }
  } else {
    A2B(A, N, 1);
    while (1) {
      if (A == 1)
        break;
      A2B(A / 10, A - 1, 1);
      A /= 10;
    }
  }
  for (int i = 0; i < 10; i++)
    cout << cnt[i] << ' ';
  cout << '\n';
  return 0;
}