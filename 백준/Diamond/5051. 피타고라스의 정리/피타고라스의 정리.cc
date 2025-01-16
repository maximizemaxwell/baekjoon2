#include<bits/stdc++.h>

using namespace std;
using ull = unsigned long long;
using cdbl = complex<double>;

const double PI = acos(-1.0);

inline unsigned bitreverse(const unsigned n, const unsigned k) {
    unsigned r, i;
    for (r = 0, i = 0; i < k; ++i)
        r |= ((n >> i) & 1) << (k - i - 1);
    return r;
}

void fft(vector<cdbl> &a, bool is_reverse=false) {
    const unsigned n = a.size(), k = __builtin_ctz(n);
    unsigned s, i, j;
    for (i = 0; i < n; i++) {
        j = bitreverse(i, k);
        if (i < j)
            swap(a[i], a[j]);
    }
    for (s = 2; s <= n; s *= 2) {
        double t = 2*PI/s * (is_reverse? -1 : 1);
        cdbl ws(cos(t), sin(t));
        for (i = 0; i < n; i += s) {
            cdbl w(1);
            for (j = 0; j < s/2; j++) {
                cdbl tmp = a[i + j + s/2] * w;
                a[i + j + s/2] = a[i + j] - tmp;
                a[i + j] += tmp;
                w *= ws;
            }
        }
    }
    if (is_reverse)
        for (i = 0; i < n; i++)
            a[i] /= n;
}

int main(void) {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);

    unsigned n, np;
    cin >> n;
    for (np = 1; np < 2*n; np *= 2);

    vector<unsigned> f(np, 0);
    for (ull x = 1; x < n; x++)
        f[x * x % n]++;

    vector<cdbl> g(f.begin(), f.end());
    vector<cdbl> fp(np, 0);
    for (unsigned i = 0; i < n; i++)
        fp[i] = fp[np-n+i] = f[(n - i) % n];

    fft(g);
    fft(fp);
    for (unsigned i = 0; i < np; i++)
        g[i] *= fp[i];
    fft(g, true);

    ull res = 0;
    for (unsigned i = 0; i < n; i++)
        res += f[i] * ((ull)(g[i].real() + 0.5) + f[2*i % n]);

    cout << res / 2;

    return 0;
}
