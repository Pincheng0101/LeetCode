# 2156. Find Substring With Given Hash Value

## Ｍodulo Optimize
```
let val(s[0]) = v0, 在 乘法運算完後做個 mod，防止數值太大會溢位
hash(s, p, m) = (val(s[0]) * p0 + val(s[1]) * p1 + ... + val(s[k-1]) * pk-1) mod m
              = (v0 * p0 mod m + v1 * p1 mod m + ... + vk-1 * pk-1 * mod m) mod m

```

## Reverse Rolling Hash (Rabin-Karp algorithm)
由於題目的 power 是依序乘上去，如果正向平移比對的話會需要計算除數，會比較麻煩
```
NextHash = ((PrevHash - Vi-1) / power + Vk * power^k) mod m 
```

所以可以倒過來處理，做乘法處理比較容易
```
NextHash = ((PrevHash - Vk+1 * power^k ) * power + Vi) mod m 
```


