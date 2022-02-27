# The game of HuaRongDao（华容道）

just decided to have a little fun solving the game of HuaRongDao.
couldn't figure out otherwise :(

## usage
```
go run cmd/main.go -board=0
```


## board layouts

* board 0 -- 116 moves

```
+----------+----------+----------+------------+
|    0     |    1     |    2     |     3      |
+----------+----------+----------+------------+
| MaChao   | CaoCao   | CaoCao   | HuangZhong |
+          +          +          +            +
|          |          |          |            |
+----------+----------+----------+------------+
| ZhangFei | GuanYu   | GuanYu   | ZhaoYun    |
+          +----------+----------+            +
|          | Soldier0 | Soldier1 |            |
+----------+----------+----------+------------+
| Soldier2 | Empty0   | Empty1   | Soldier3   |
+----------+----------+----------+------------+
```

* board 1 -- 82 moves

```
+----------+----------+----------+------------+
|    0     |    1     |    2     |     3      |
+----------+----------+----------+------------+
| MaChao   | CaoCao   | CaoCao   | HuangZhong |
+          +          +          +            +
|          |          |          |            |
+----------+----------+----------+------------+
| ZhangFei | Soldier0 | Soldier1 | ZhaoYun    |
+          +----------+----------+            +
|          | Soldier2 | Soldier3 |            |
+----------+----------+----------+------------+
| Empty0   | GuanYu   | GuanYu   | Empty1     |
+----------+----------+----------+------------+
```

* board 2 -- 43 moves

```
+------------+--------+----------+----------+
|     0      |   1    |    2     |    3     |
+------------+--------+----------+----------+
| HuangZhong | MaChao | ZhaoYun  | Soldier0 |
+            +        +          +----------+
|            |        |          | Soldier1 |
+------------+--------+----------+----------+
| CaoCao     | CaoCao | Empty0   | Empty1   |
+            +        +----------+----------+
|            |        | Soldier2 | ZhangFei |
+------------+--------+----------+          +
| GuanYu     | GuanYu | Soldier3 |          |
+------------+--------+----------+----------+

```
