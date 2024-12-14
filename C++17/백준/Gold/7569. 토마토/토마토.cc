#include <algorithm>
#include <iostream>
#include <queue>

using namespace std;

int m, n, h;

int dx[] = {0, 0, 1, -1, 0, 0};
int dy[] = {1, -1, 0, 0, 0, 0};
int dz[] = {0, 0, 0, 0, 1, -1};

int box[101][101][101];
int vis[101][101][101];

struct Position {
  int x, y, z;
};

void bfs(queue<Position> &Q) {
  while (!Q.empty()) {
    Position cur = Q.front();
    Q.pop();
    int x = cur.x;
    int y = cur.y;
    int z = cur.z;

    for (int dir = 0; dir < 6; dir++) {
      int nx = x + dx[dir];
      int ny = y + dy[dir];
      int nz = z + dz[dir];

      if (nx < 0 || nx >= m || ny < 0 || ny >= n || nz < 0 || nz >= h)
        continue;
      if (vis[nz][ny][nx] || box[nz][ny][nx] != 0)
        continue;

      vis[nz][ny][nx] = vis[z][y][x] + 1;
      Q.push({nx, ny, nz});
    }
  }
}

int calculateDays() {
  int days = 0;
  for (int i = 0; i < h; i++) {
    for (int j = 0; j < n; j++) {
      for (int k = 0; k < m; k++) {
        if (box[i][j][k] == 0 && vis[i][j][k] == 0) {
          return -1; // 익지 않은 토마토가 있는 경우
        }
        days = max(days, vis[i][j][k]);
      }
    }
  }
  return days - 1; // 첫날이 1이므로 1을 뺌
}

int main(void) {
  ios::sync_with_stdio(0);
  cin.tie(0);

  cin >> m >> n >> h;
  queue<Position> Q;

  for (int i = 0; i < h; i++) {
    for (int j = 0; j < n; j++) {
      for (int k = 0; k < m; k++) {
        cin >> box[i][j][k];

        if (box[i][j][k] == 1) {
          Q.push({k, j, i}); // x=k, y=j, z=i
          vis[i][j][k] = 1;  // 방문 처리
        }
      }
    }
  }

  bfs(Q);
  cout << calculateDays() << '\n';
  return 0;
}