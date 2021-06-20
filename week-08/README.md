# week 08 作业

1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

## 解答

### 1

```sh
# 获取 Redis 镜像 并启动
docker search redis
docker pull redis
docker images
docker ps
docker run -p 6379:6379 redis:latest redis-server

# 重命名为 redis-test 后台运行
docker run -itd --name redis-test -p 6379:6379 redis

# 登录服务端后启动客户端
docker exec -it redis-test /bin/bash
redis-cli

# 直接启动客户端
docker exec -it redis-test redis-cli
docker exec -it redis-test redis-benchmark -t get,set -d 10 -n 10000 -q

# 停止镜像
docker stop redis-test
docker rm redis-test
```

测试命令

```sh
# 简要信息
docker exec -it redis-test redis-benchmark -t get,set -d 10 -n 10000 -q
docker exec -it redis-test redis-benchmark -t get,set -d 20 -n 10000 -q
docker exec -it redis-test redis-benchmark -t get,set -d 50 -n 10000 -q
docker exec -it redis-test redis-benchmark -t get,set -d 100 -n 10000 -q
docker exec -it redis-test redis-benchmark -t get,set -d 200 -n 10000 -q
docker exec -it redis-test redis-benchmark -t get,set -d 1000 -n 10000 -q
docker exec -it redis-test redis-benchmark -t get,set -d 5000 -n 10000 -q

# 详细信息
docker exec -it redis-test redis-benchmark -t get,set -d 10 -n 10000
docker exec -it redis-test redis-benchmark -t get,set -d 20 -n 10000
docker exec -it redis-test redis-benchmark -t get,set -d 50 -n 10000
docker exec -it redis-test redis-benchmark -t get,set -d 100 -n 10000
docker exec -it redis-test redis-benchmark -t get,set -d 200 -n 10000
docker exec -it redis-test redis-benchmark -t get,set -d 1000 -n 10000
docker exec -it redis-test redis-benchmark -t get,set -d 5000 -n 10000
```

简要测试结果

```sh
➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 10 -n 10000 -q

SET: 20366.60 requests per second, p50=1.999 msec
GET: 20533.88 requests per second, p50=1.999 msec

➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 10 -n 10000 -q
SET: 20283.98 requests per second, p50=2.023 msec
GET: 20408.16 requests per second, p50=2.007 msec

➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 20 -n 10000 -q
SET: 20080.32 requests per second, p50=2.031 msec
GET: 20746.89 requests per second, p50=1.975 msec

➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 50 -n 10000 -q
SET: 20283.98 requests per second, p50=1.999 msec
GET: 20000.00 requests per second, p50=1.999 msec

➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 100 -n 10000 -q
SET: 16863.41 requests per second, p50=2.335 msec
GET: 19880.71 requests per second, p50=2.055 msec

➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 200 -n 10000 -q
SET: 20202.02 requests per second, p50=2.031 msec
GET: 20576.13 requests per second, p50=1.991 msec

➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 1000 -n 10000 -q
SET: 20202.02 requests per second, p50=2.031 msec
GET: 20366.60 requests per second, p50=2.023 msec

➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 5000 -n 10000 -q
SET: 19841.27 requests per second, p50=2.071 msec
GET: 19762.85 requests per second, p50=2.047 msec

➜  ~
```

详细测试结果

```sh
➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 10 -n 10000
====== SET ======
  10000 requests completed in 0.50 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.959 milliseconds (cumulative count 1)
50.000% <= 2.055 milliseconds (cumulative count 5062)
75.000% <= 2.431 milliseconds (cumulative count 7502)
87.500% <= 2.711 milliseconds (cumulative count 8766)
93.750% <= 2.863 milliseconds (cumulative count 9384)
96.875% <= 2.999 milliseconds (cumulative count 9693)
98.438% <= 3.223 milliseconds (cumulative count 9844)
99.219% <= 3.455 milliseconds (cumulative count 9924)
99.609% <= 3.647 milliseconds (cumulative count 9962)
99.805% <= 3.887 milliseconds (cumulative count 9981)
99.902% <= 4.055 milliseconds (cumulative count 9991)
99.951% <= 4.471 milliseconds (cumulative count 9996)
99.976% <= 4.623 milliseconds (cumulative count 9998)
99.988% <= 4.695 milliseconds (cumulative count 9999)
99.994% <= 4.799 milliseconds (cumulative count 10000)
100.000% <= 4.799 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.030% <= 1.007 milliseconds (cumulative count 3)
0.110% <= 1.103 milliseconds (cumulative count 11)
0.300% <= 1.207 milliseconds (cumulative count 30)
0.550% <= 1.303 milliseconds (cumulative count 55)
1.270% <= 1.407 milliseconds (cumulative count 127)
3.960% <= 1.503 milliseconds (cumulative count 396)
10.010% <= 1.607 milliseconds (cumulative count 1001)
15.790% <= 1.703 milliseconds (cumulative count 1579)
25.310% <= 1.807 milliseconds (cumulative count 2531)
37.760% <= 1.903 milliseconds (cumulative count 3776)
47.200% <= 2.007 milliseconds (cumulative count 4720)
53.900% <= 2.103 milliseconds (cumulative count 5390)
97.860% <= 3.103 milliseconds (cumulative count 9786)
99.920% <= 4.103 milliseconds (cumulative count 9992)
100.000% <= 5.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20080.32 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.132     0.952     2.055     2.903     3.367     4.799
====== GET ======
  10000 requests completed in 0.49 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.911 milliseconds (cumulative count 1)
50.000% <= 1.999 milliseconds (cumulative count 5055)
75.000% <= 2.375 milliseconds (cumulative count 7529)
87.500% <= 2.647 milliseconds (cumulative count 8769)
93.750% <= 2.791 milliseconds (cumulative count 9386)
96.875% <= 2.895 milliseconds (cumulative count 9699)
98.438% <= 3.023 milliseconds (cumulative count 9844)
99.219% <= 3.255 milliseconds (cumulative count 9922)
99.609% <= 3.535 milliseconds (cumulative count 9961)
99.805% <= 3.751 milliseconds (cumulative count 9981)
99.902% <= 3.879 milliseconds (cumulative count 9991)
99.951% <= 4.007 milliseconds (cumulative count 9996)
99.976% <= 4.039 milliseconds (cumulative count 9998)
99.988% <= 4.095 milliseconds (cumulative count 9999)
99.994% <= 4.111 milliseconds (cumulative count 10000)
100.000% <= 4.111 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.050% <= 1.007 milliseconds (cumulative count 5)
0.120% <= 1.103 milliseconds (cumulative count 12)
0.300% <= 1.207 milliseconds (cumulative count 30)
0.540% <= 1.303 milliseconds (cumulative count 54)
1.450% <= 1.407 milliseconds (cumulative count 145)
5.470% <= 1.503 milliseconds (cumulative count 547)
12.170% <= 1.607 milliseconds (cumulative count 1217)
18.890% <= 1.703 milliseconds (cumulative count 1889)
30.640% <= 1.807 milliseconds (cumulative count 3064)
41.850% <= 1.903 milliseconds (cumulative count 4185)
51.130% <= 2.007 milliseconds (cumulative count 5113)
57.840% <= 2.103 milliseconds (cumulative count 5784)
98.800% <= 3.103 milliseconds (cumulative count 9880)
99.990% <= 4.103 milliseconds (cumulative count 9999)
100.000% <= 5.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20576.13 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.080     0.904     1.999     2.823     3.175     4.111

➜  ~
```

```sh
➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 20 -n 10000
====== SET ======
  10000 requests completed in 0.49 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.007 milliseconds (cumulative count 1)
50.000% <= 2.007 milliseconds (cumulative count 5019)
75.000% <= 2.383 milliseconds (cumulative count 7505)
87.500% <= 2.663 milliseconds (cumulative count 8765)
93.750% <= 2.815 milliseconds (cumulative count 9387)
96.875% <= 2.951 milliseconds (cumulative count 9692)
98.438% <= 3.239 milliseconds (cumulative count 9846)
99.219% <= 3.607 milliseconds (cumulative count 9922)
99.609% <= 4.095 milliseconds (cumulative count 9961)
99.805% <= 4.503 milliseconds (cumulative count 9981)
99.902% <= 4.927 milliseconds (cumulative count 9991)
99.951% <= 5.239 milliseconds (cumulative count 9996)
99.976% <= 5.439 milliseconds (cumulative count 9998)
99.988% <= 5.479 milliseconds (cumulative count 9999)
99.994% <= 5.495 milliseconds (cumulative count 10000)
100.000% <= 5.495 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.010% <= 1.007 milliseconds (cumulative count 1)
0.070% <= 1.103 milliseconds (cumulative count 7)
0.150% <= 1.207 milliseconds (cumulative count 15)
0.350% <= 1.303 milliseconds (cumulative count 35)
1.330% <= 1.407 milliseconds (cumulative count 133)
5.220% <= 1.503 milliseconds (cumulative count 522)
11.610% <= 1.607 milliseconds (cumulative count 1161)
18.240% <= 1.703 milliseconds (cumulative count 1824)
29.970% <= 1.807 milliseconds (cumulative count 2997)
41.540% <= 1.903 milliseconds (cumulative count 4154)
50.190% <= 2.007 milliseconds (cumulative count 5019)
56.970% <= 2.103 milliseconds (cumulative count 5697)
97.870% <= 3.103 milliseconds (cumulative count 9787)
99.610% <= 4.103 milliseconds (cumulative count 9961)
99.930% <= 5.103 milliseconds (cumulative count 9993)
100.000% <= 6.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20325.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.100     1.000     2.007     2.855     3.455     5.495
====== GET ======
  10000 requests completed in 0.48 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.999 milliseconds (cumulative count 1)
50.000% <= 1.967 milliseconds (cumulative count 5032)
75.000% <= 2.319 milliseconds (cumulative count 7521)
87.500% <= 2.583 milliseconds (cumulative count 8757)
93.750% <= 2.727 milliseconds (cumulative count 9392)
96.875% <= 2.815 milliseconds (cumulative count 9719)
98.438% <= 2.879 milliseconds (cumulative count 9851)
99.219% <= 2.975 milliseconds (cumulative count 9923)
99.609% <= 3.055 milliseconds (cumulative count 9964)
99.805% <= 3.119 milliseconds (cumulative count 9981)
99.902% <= 3.223 milliseconds (cumulative count 9991)
99.951% <= 3.311 milliseconds (cumulative count 9996)
99.976% <= 3.343 milliseconds (cumulative count 9998)
99.988% <= 3.399 milliseconds (cumulative count 9999)
99.994% <= 3.447 milliseconds (cumulative count 10000)
100.000% <= 3.447 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.010% <= 1.007 milliseconds (cumulative count 1)
0.050% <= 1.103 milliseconds (cumulative count 5)
0.190% <= 1.207 milliseconds (cumulative count 19)
0.460% <= 1.303 milliseconds (cumulative count 46)
1.690% <= 1.407 milliseconds (cumulative count 169)
6.690% <= 1.503 milliseconds (cumulative count 669)
13.430% <= 1.607 milliseconds (cumulative count 1343)
21.060% <= 1.703 milliseconds (cumulative count 2106)
34.020% <= 1.807 milliseconds (cumulative count 3402)
45.150% <= 1.903 milliseconds (cumulative count 4515)
53.430% <= 2.007 milliseconds (cumulative count 5343)
60.440% <= 2.103 milliseconds (cumulative count 6044)
99.760% <= 3.103 milliseconds (cumulative count 9976)
100.000% <= 4.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 21008.40 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.037     0.992     1.967     2.759     2.943     3.447

➜  ~
```

```sh
➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 50 -n 10000
====== SET ======
  10000 requests completed in 0.50 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.983 milliseconds (cumulative count 2)
50.000% <= 2.031 milliseconds (cumulative count 5017)
75.000% <= 2.423 milliseconds (cumulative count 7531)
87.500% <= 2.695 milliseconds (cumulative count 8773)
93.750% <= 2.847 milliseconds (cumulative count 9400)
96.875% <= 2.983 milliseconds (cumulative count 9688)
98.438% <= 3.215 milliseconds (cumulative count 9844)
99.219% <= 3.503 milliseconds (cumulative count 9922)
99.609% <= 3.927 milliseconds (cumulative count 9962)
99.805% <= 4.215 milliseconds (cumulative count 9981)
99.902% <= 4.591 milliseconds (cumulative count 9991)
99.951% <= 4.807 milliseconds (cumulative count 9996)
99.976% <= 5.015 milliseconds (cumulative count 9998)
99.988% <= 5.039 milliseconds (cumulative count 9999)
99.994% <= 5.135 milliseconds (cumulative count 10000)
100.000% <= 5.135 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.020% <= 1.007 milliseconds (cumulative count 2)
0.090% <= 1.103 milliseconds (cumulative count 9)
0.280% <= 1.207 milliseconds (cumulative count 28)
0.430% <= 1.303 milliseconds (cumulative count 43)
1.130% <= 1.407 milliseconds (cumulative count 113)
4.460% <= 1.503 milliseconds (cumulative count 446)
10.650% <= 1.607 milliseconds (cumulative count 1065)
16.640% <= 1.703 milliseconds (cumulative count 1664)
27.320% <= 1.807 milliseconds (cumulative count 2732)
39.070% <= 1.903 milliseconds (cumulative count 3907)
48.330% <= 2.007 milliseconds (cumulative count 4833)
55.280% <= 2.103 milliseconds (cumulative count 5528)
97.880% <= 3.103 milliseconds (cumulative count 9788)
99.740% <= 4.103 milliseconds (cumulative count 9974)
99.990% <= 5.103 milliseconds (cumulative count 9999)
100.000% <= 6.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20202.02 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.120     0.976     2.031     2.887     3.407     5.135
====== GET ======
  10000 requests completed in 0.49 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.951 milliseconds (cumulative count 1)
50.000% <= 1.991 milliseconds (cumulative count 5042)
75.000% <= 2.359 milliseconds (cumulative count 7503)
87.500% <= 2.639 milliseconds (cumulative count 8775)
93.750% <= 2.791 milliseconds (cumulative count 9391)
96.875% <= 2.879 milliseconds (cumulative count 9700)
98.438% <= 3.007 milliseconds (cumulative count 9846)
99.219% <= 3.207 milliseconds (cumulative count 9922)
99.609% <= 3.631 milliseconds (cumulative count 9961)
99.805% <= 4.279 milliseconds (cumulative count 9981)
99.902% <= 4.839 milliseconds (cumulative count 9991)
99.951% <= 5.303 milliseconds (cumulative count 9996)
99.976% <= 5.423 milliseconds (cumulative count 9998)
99.988% <= 5.495 milliseconds (cumulative count 9999)
99.994% <= 5.559 milliseconds (cumulative count 10000)
100.000% <= 5.559 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.050% <= 1.007 milliseconds (cumulative count 5)
0.230% <= 1.103 milliseconds (cumulative count 23)
0.530% <= 1.207 milliseconds (cumulative count 53)
0.890% <= 1.303 milliseconds (cumulative count 89)
1.860% <= 1.407 milliseconds (cumulative count 186)
5.990% <= 1.503 milliseconds (cumulative count 599)
12.720% <= 1.607 milliseconds (cumulative count 1272)
18.880% <= 1.703 milliseconds (cumulative count 1888)
30.500% <= 1.807 milliseconds (cumulative count 3050)
42.540% <= 1.903 milliseconds (cumulative count 4254)
51.620% <= 2.007 milliseconds (cumulative count 5162)
58.260% <= 2.103 milliseconds (cumulative count 5826)
98.900% <= 3.103 milliseconds (cumulative count 9890)
99.790% <= 4.103 milliseconds (cumulative count 9979)
99.940% <= 5.103 milliseconds (cumulative count 9994)
100.000% <= 6.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20491.80 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.075     0.944     1.991     2.823     3.135     5.559

➜  ~
```

```sh
➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 100 -n 10000
====== SET ======
  10000 requests completed in 0.49 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.983 milliseconds (cumulative count 1)
50.000% <= 2.015 milliseconds (cumulative count 5029)
75.000% <= 2.399 milliseconds (cumulative count 7533)
87.500% <= 2.671 milliseconds (cumulative count 8782)
93.750% <= 2.823 milliseconds (cumulative count 9390)
96.875% <= 2.943 milliseconds (cumulative count 9694)
98.438% <= 3.175 milliseconds (cumulative count 9844)
99.219% <= 3.535 milliseconds (cumulative count 9923)
99.609% <= 3.839 milliseconds (cumulative count 9961)
99.805% <= 4.159 milliseconds (cumulative count 9981)
99.902% <= 4.423 milliseconds (cumulative count 9991)
99.951% <= 4.591 milliseconds (cumulative count 9996)
99.976% <= 4.951 milliseconds (cumulative count 9998)
99.988% <= 5.007 milliseconds (cumulative count 9999)
99.994% <= 5.127 milliseconds (cumulative count 10000)
100.000% <= 5.127 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.010% <= 1.007 milliseconds (cumulative count 1)
0.060% <= 1.103 milliseconds (cumulative count 6)
0.190% <= 1.207 milliseconds (cumulative count 19)
0.380% <= 1.303 milliseconds (cumulative count 38)
1.290% <= 1.407 milliseconds (cumulative count 129)
5.210% <= 1.503 milliseconds (cumulative count 521)
11.570% <= 1.607 milliseconds (cumulative count 1157)
17.930% <= 1.703 milliseconds (cumulative count 1793)
29.010% <= 1.807 milliseconds (cumulative count 2901)
41.070% <= 1.903 milliseconds (cumulative count 4107)
49.590% <= 2.007 milliseconds (cumulative count 4959)
56.710% <= 2.103 milliseconds (cumulative count 5671)
98.130% <= 3.103 milliseconds (cumulative count 9813)
99.790% <= 4.103 milliseconds (cumulative count 9979)
99.990% <= 5.103 milliseconds (cumulative count 9999)
100.000% <= 6.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20366.60 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.101     0.976     2.015     2.863     3.415     5.127
====== GET ======
  10000 requests completed in 0.49 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.887 milliseconds (cumulative count 1)
50.000% <= 2.007 milliseconds (cumulative count 5047)
75.000% <= 2.383 milliseconds (cumulative count 7519)
87.500% <= 2.655 milliseconds (cumulative count 8787)
93.750% <= 2.799 milliseconds (cumulative count 9377)
96.875% <= 2.927 milliseconds (cumulative count 9690)
98.438% <= 3.135 milliseconds (cumulative count 9846)
99.219% <= 3.367 milliseconds (cumulative count 9923)
99.609% <= 3.743 milliseconds (cumulative count 9961)
99.805% <= 4.023 milliseconds (cumulative count 9981)
99.902% <= 4.263 milliseconds (cumulative count 9991)
99.951% <= 4.439 milliseconds (cumulative count 9996)
99.976% <= 4.591 milliseconds (cumulative count 9998)
99.988% <= 4.607 milliseconds (cumulative count 9999)
99.994% <= 4.719 milliseconds (cumulative count 10000)
100.000% <= 4.719 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.010% <= 0.903 milliseconds (cumulative count 1)
0.030% <= 1.007 milliseconds (cumulative count 3)
0.090% <= 1.103 milliseconds (cumulative count 9)
0.260% <= 1.207 milliseconds (cumulative count 26)
0.530% <= 1.303 milliseconds (cumulative count 53)
1.600% <= 1.407 milliseconds (cumulative count 160)
5.910% <= 1.503 milliseconds (cumulative count 591)
12.270% <= 1.607 milliseconds (cumulative count 1227)
19.220% <= 1.703 milliseconds (cumulative count 1922)
30.780% <= 1.807 milliseconds (cumulative count 3078)
41.740% <= 1.903 milliseconds (cumulative count 4174)
50.470% <= 2.007 milliseconds (cumulative count 5047)
57.120% <= 2.103 milliseconds (cumulative count 5712)
98.260% <= 3.103 milliseconds (cumulative count 9826)
99.830% <= 4.103 milliseconds (cumulative count 9983)
100.000% <= 5.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20491.80 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.087     0.880     2.007     2.831     3.287     4.719

➜  ~
```

```sh
➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 200 -n 10000
====== SET ======
  10000 requests completed in 0.49 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.999 milliseconds (cumulative count 1)
50.000% <= 1.991 milliseconds (cumulative count 5032)
75.000% <= 2.367 milliseconds (cumulative count 7512)
87.500% <= 2.639 milliseconds (cumulative count 8773)
93.750% <= 2.791 milliseconds (cumulative count 9390)
96.875% <= 2.879 milliseconds (cumulative count 9690)
98.438% <= 2.975 milliseconds (cumulative count 9844)
99.219% <= 3.119 milliseconds (cumulative count 9923)
99.609% <= 3.335 milliseconds (cumulative count 9962)
99.805% <= 3.623 milliseconds (cumulative count 9981)
99.902% <= 3.895 milliseconds (cumulative count 9991)
99.951% <= 4.055 milliseconds (cumulative count 9996)
99.976% <= 4.183 milliseconds (cumulative count 9998)
99.988% <= 4.207 milliseconds (cumulative count 9999)
99.994% <= 4.399 milliseconds (cumulative count 10000)
100.000% <= 4.399 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.010% <= 1.007 milliseconds (cumulative count 1)
0.160% <= 1.103 milliseconds (cumulative count 16)
0.620% <= 1.207 milliseconds (cumulative count 62)
1.370% <= 1.303 milliseconds (cumulative count 137)
3.140% <= 1.407 milliseconds (cumulative count 314)
7.150% <= 1.503 milliseconds (cumulative count 715)
13.720% <= 1.607 milliseconds (cumulative count 1372)
20.380% <= 1.703 milliseconds (cumulative count 2038)
30.840% <= 1.807 milliseconds (cumulative count 3084)
41.990% <= 1.903 milliseconds (cumulative count 4199)
51.500% <= 2.007 milliseconds (cumulative count 5150)
58.370% <= 2.103 milliseconds (cumulative count 5837)
99.190% <= 3.103 milliseconds (cumulative count 9919)
99.960% <= 4.103 milliseconds (cumulative count 9996)
100.000% <= 5.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20576.13 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.066     0.992     1.991     2.823     3.063     4.399
====== GET ======
  10000 requests completed in 0.49 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.967 milliseconds (cumulative count 1)
50.000% <= 2.015 milliseconds (cumulative count 5041)
75.000% <= 2.391 milliseconds (cumulative count 7513)
87.500% <= 2.679 milliseconds (cumulative count 8782)
93.750% <= 2.831 milliseconds (cumulative count 9399)
96.875% <= 2.999 milliseconds (cumulative count 9688)
98.438% <= 3.311 milliseconds (cumulative count 9844)
99.219% <= 3.711 milliseconds (cumulative count 9924)
99.609% <= 4.151 milliseconds (cumulative count 9961)
99.805% <= 4.439 milliseconds (cumulative count 9981)
99.902% <= 4.575 milliseconds (cumulative count 9991)
99.951% <= 4.807 milliseconds (cumulative count 9996)
99.976% <= 4.959 milliseconds (cumulative count 9998)
99.988% <= 5.007 milliseconds (cumulative count 9999)
99.994% <= 5.071 milliseconds (cumulative count 10000)
100.000% <= 5.071 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.030% <= 1.007 milliseconds (cumulative count 3)
0.110% <= 1.103 milliseconds (cumulative count 11)
0.270% <= 1.207 milliseconds (cumulative count 27)
0.540% <= 1.303 milliseconds (cumulative count 54)
1.740% <= 1.407 milliseconds (cumulative count 174)
6.090% <= 1.503 milliseconds (cumulative count 609)
12.550% <= 1.607 milliseconds (cumulative count 1255)
19.070% <= 1.703 milliseconds (cumulative count 1907)
30.490% <= 1.807 milliseconds (cumulative count 3049)
41.380% <= 1.903 milliseconds (cumulative count 4138)
49.920% <= 2.007 milliseconds (cumulative count 4992)
56.590% <= 2.103 milliseconds (cumulative count 5659)
97.670% <= 3.103 milliseconds (cumulative count 9767)
99.570% <= 4.103 milliseconds (cumulative count 9957)
100.000% <= 5.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20366.60 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.100     0.960     2.015     2.871     3.535     5.071

➜  ~
```

```sh
➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 1000 -n 10000
====== SET ======
  10000 requests completed in 0.50 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.007 milliseconds (cumulative count 1)
50.000% <= 2.047 milliseconds (cumulative count 5015)
75.000% <= 2.439 milliseconds (cumulative count 7517)
87.500% <= 2.719 milliseconds (cumulative count 8752)
93.750% <= 2.879 milliseconds (cumulative count 9375)
96.875% <= 3.055 milliseconds (cumulative count 9690)
98.438% <= 3.455 milliseconds (cumulative count 9845)
99.219% <= 4.119 milliseconds (cumulative count 9924)
99.609% <= 4.599 milliseconds (cumulative count 9962)
99.805% <= 4.959 milliseconds (cumulative count 9982)
99.902% <= 5.303 milliseconds (cumulative count 9991)
99.951% <= 6.215 milliseconds (cumulative count 9996)
99.976% <= 6.335 milliseconds (cumulative count 9998)
99.988% <= 6.407 milliseconds (cumulative count 9999)
99.994% <= 6.487 milliseconds (cumulative count 10000)
100.000% <= 6.487 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.010% <= 1.007 milliseconds (cumulative count 1)
0.070% <= 1.103 milliseconds (cumulative count 7)
0.210% <= 1.207 milliseconds (cumulative count 21)
0.400% <= 1.303 milliseconds (cumulative count 40)
0.850% <= 1.407 milliseconds (cumulative count 85)
3.170% <= 1.503 milliseconds (cumulative count 317)
9.710% <= 1.607 milliseconds (cumulative count 971)
15.460% <= 1.703 milliseconds (cumulative count 1546)
24.400% <= 1.807 milliseconds (cumulative count 2440)
37.150% <= 1.903 milliseconds (cumulative count 3715)
47.130% <= 2.007 milliseconds (cumulative count 4713)
53.920% <= 2.103 milliseconds (cumulative count 5392)
97.110% <= 3.103 milliseconds (cumulative count 9711)
99.210% <= 4.103 milliseconds (cumulative count 9921)
99.850% <= 5.103 milliseconds (cumulative count 9985)
99.940% <= 6.103 milliseconds (cumulative count 9994)
100.000% <= 7.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 19920.32 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.153     1.000     2.047     2.919     3.903     6.487
====== GET ======
  10000 requests completed in 0.50 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.975 milliseconds (cumulative count 1)
50.000% <= 2.015 milliseconds (cumulative count 5016)
75.000% <= 2.407 milliseconds (cumulative count 7505)
87.500% <= 2.695 milliseconds (cumulative count 8760)
93.750% <= 2.871 milliseconds (cumulative count 9385)
96.875% <= 3.159 milliseconds (cumulative count 9688)
98.438% <= 3.711 milliseconds (cumulative count 9844)
99.219% <= 4.407 milliseconds (cumulative count 9922)
99.609% <= 4.863 milliseconds (cumulative count 9961)
99.805% <= 5.807 milliseconds (cumulative count 9981)
99.902% <= 6.455 milliseconds (cumulative count 9991)
99.951% <= 7.207 milliseconds (cumulative count 9996)
99.976% <= 7.359 milliseconds (cumulative count 9998)
99.988% <= 7.423 milliseconds (cumulative count 9999)
99.994% <= 7.511 milliseconds (cumulative count 10000)
100.000% <= 7.511 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.020% <= 1.007 milliseconds (cumulative count 2)
0.090% <= 1.103 milliseconds (cumulative count 9)
0.200% <= 1.207 milliseconds (cumulative count 20)
0.340% <= 1.303 milliseconds (cumulative count 34)
1.510% <= 1.407 milliseconds (cumulative count 151)
5.810% <= 1.503 milliseconds (cumulative count 581)
11.840% <= 1.607 milliseconds (cumulative count 1184)
18.710% <= 1.703 milliseconds (cumulative count 1871)
29.850% <= 1.807 milliseconds (cumulative count 2985)
40.970% <= 1.903 milliseconds (cumulative count 4097)
49.600% <= 2.007 milliseconds (cumulative count 4960)
56.410% <= 2.103 milliseconds (cumulative count 5641)
96.620% <= 3.103 milliseconds (cumulative count 9662)
98.940% <= 4.103 milliseconds (cumulative count 9894)
99.700% <= 5.103 milliseconds (cumulative count 9970)
99.870% <= 6.103 milliseconds (cumulative count 9987)
99.940% <= 7.103 milliseconds (cumulative count 9994)
100.000% <= 8.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20080.32 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.130     0.968     2.015     2.935     4.151     7.511

➜  ~
```

```sh
➜  ~ docker exec -it redis-test redis-benchmark -t get,set -d 5000 -n 10000
====== SET ======
  10000 requests completed in 0.51 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.007 milliseconds (cumulative count 1)
50.000% <= 2.103 milliseconds (cumulative count 5026)
75.000% <= 2.503 milliseconds (cumulative count 7531)
87.500% <= 2.783 milliseconds (cumulative count 8775)
93.750% <= 2.943 milliseconds (cumulative count 9393)
96.875% <= 3.087 milliseconds (cumulative count 9694)
98.438% <= 3.343 milliseconds (cumulative count 9849)
99.219% <= 3.631 milliseconds (cumulative count 9923)
99.609% <= 4.063 milliseconds (cumulative count 9961)
99.805% <= 4.479 milliseconds (cumulative count 9982)
99.902% <= 4.663 milliseconds (cumulative count 9991)
99.951% <= 4.767 milliseconds (cumulative count 9996)
99.976% <= 4.847 milliseconds (cumulative count 9998)
99.988% <= 4.903 milliseconds (cumulative count 9999)
99.994% <= 4.935 milliseconds (cumulative count 10000)
100.000% <= 4.935 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.010% <= 1.007 milliseconds (cumulative count 1)
0.090% <= 1.103 milliseconds (cumulative count 9)
0.330% <= 1.207 milliseconds (cumulative count 33)
0.710% <= 1.303 milliseconds (cumulative count 71)
1.240% <= 1.407 milliseconds (cumulative count 124)
3.140% <= 1.503 milliseconds (cumulative count 314)
8.430% <= 1.607 milliseconds (cumulative count 843)
14.010% <= 1.703 milliseconds (cumulative count 1401)
21.220% <= 1.807 milliseconds (cumulative count 2122)
32.170% <= 1.903 milliseconds (cumulative count 3217)
42.650% <= 2.007 milliseconds (cumulative count 4265)
50.260% <= 2.103 milliseconds (cumulative count 5026)
97.050% <= 3.103 milliseconds (cumulative count 9705)
99.630% <= 4.103 milliseconds (cumulative count 9963)
100.000% <= 5.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 19493.18 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.188     1.000     2.103     2.983     3.503     4.935
====== GET ======
  10000 requests completed in 0.49 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.935 milliseconds (cumulative count 1)
50.000% <= 2.007 milliseconds (cumulative count 5019)
75.000% <= 2.399 milliseconds (cumulative count 7523)
87.500% <= 2.671 milliseconds (cumulative count 8750)
93.750% <= 2.823 milliseconds (cumulative count 9390)
96.875% <= 2.919 milliseconds (cumulative count 9688)
98.438% <= 3.031 milliseconds (cumulative count 9849)
99.219% <= 3.231 milliseconds (cumulative count 9922)
99.609% <= 3.479 milliseconds (cumulative count 9963)
99.805% <= 3.815 milliseconds (cumulative count 9981)
99.902% <= 4.111 milliseconds (cumulative count 9991)
99.951% <= 4.391 milliseconds (cumulative count 9996)
99.976% <= 4.535 milliseconds (cumulative count 9998)
99.988% <= 4.599 milliseconds (cumulative count 9999)
99.994% <= 4.703 milliseconds (cumulative count 10000)
100.000% <= 4.703 milliseconds (cumulative count 10000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.020% <= 1.007 milliseconds (cumulative count 2)
0.100% <= 1.103 milliseconds (cumulative count 10)
0.270% <= 1.207 milliseconds (cumulative count 27)
0.500% <= 1.303 milliseconds (cumulative count 50)
2.000% <= 1.407 milliseconds (cumulative count 200)
6.330% <= 1.503 milliseconds (cumulative count 633)
12.120% <= 1.607 milliseconds (cumulative count 1212)
18.500% <= 1.703 milliseconds (cumulative count 1850)
28.380% <= 1.807 milliseconds (cumulative count 2838)
40.000% <= 1.903 milliseconds (cumulative count 4000)
50.190% <= 2.007 milliseconds (cumulative count 5019)
57.040% <= 2.103 milliseconds (cumulative count 5704)
98.910% <= 3.103 milliseconds (cumulative count 9891)
99.900% <= 4.103 milliseconds (cumulative count 9990)
100.000% <= 5.103 milliseconds (cumulative count 10000)

Summary:
  throughput summary: 20325.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.094     0.928     2.007     2.855     3.127     4.703

➜  ~
```

### 2

测试命令

```sh
docker run -itd --name redis-test -p 6379:6379 -v $PWD/data:/data redis:latest redis-server

docker exec -it redis-test redis-cli
keys *
info
dbsize
debug populate 10000

$ cat /tmp/script.lua
return redis.call('set',KEYS[1],ARGV[1])
$ redis-cli --eval /tmp/script.lua foo , bar
OK
```

```sh
# value 字节数 , key 个数
redis-cli --eval tmp/script.lua 10 , 10
```

tmp/script.lua
```lua
local value = ''

-- 生成指定长度 value
for i=1, KEYS[1] do
    value = value .. 'H'
end

-- 加入指定数量 key
for i=1, ARGV[1] do
    redis.call('set', i, value)
end

return
```

测试前

```sh
127.0.0.1:6379>
127.0.0.1:6379> info
# Server
redis_version:6.2.4
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:cdd247b73d61004a
redis_mode:standalone
os:Linux 5.10.25-linuxkit x86_64
arch_bits:64
multiplexing_api:epoll
atomicvar_api:c11-builtin
gcc_version:8.3.0
process_id:1
process_supervised:no
run_id:44f562e7eee35c87de10e0e8e50f948ac58769cf
tcp_port:6379
server_time_usec:1624150990645863
uptime_in_seconds:61
uptime_in_days:0
hz:10
configured_hz:10
lru_clock:13538254
executable:/data/redis-server
config_file:
io_threads_active:0

# Clients
connected_clients:1
cluster_connections:0
maxclients:10000
client_recent_max_input_buffer:16
client_recent_max_output_buffer:0
blocked_clients:0
tracking_clients:0
clients_in_timeout_table:0

# Memory
used_memory:873584
used_memory_human:853.11K
used_memory_rss:8032256
used_memory_rss_human:7.66M
used_memory_peak:931584
used_memory_peak_human:909.75K
used_memory_peak_perc:93.77%
used_memory_overhead:830352
used_memory_startup:809856
used_memory_dataset:43232
used_memory_dataset_perc:67.84%
allocator_allocated:1051272
allocator_active:1277952
allocator_resident:3985408
total_system_memory:2083807232
total_system_memory_human:1.94G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.22
allocator_frag_bytes:226680
allocator_rss_ratio:3.12
allocator_rss_bytes:2707456
rss_overhead_ratio:2.02
rss_overhead_bytes:4046848
mem_fragmentation_ratio:9.67
mem_fragmentation_bytes:7201440
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

# Persistence
loading:0
current_cow_size:0
current_cow_size_age:0
current_fork_perc:0.00
current_save_keys_processed:0
current_save_keys_total:0
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1624150929
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_last_cow_size:0
aof_enabled:0
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:-1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_last_write_status:ok
aof_last_cow_size:0
module_fork_in_progress:0
module_fork_last_cow_size:0

# Stats
total_connections_received:1
total_commands_processed:1
instantaneous_ops_per_sec:0
total_net_input_bytes:31
total_net_output_bytes:20324
instantaneous_input_kbps:0.00
instantaneous_output_kbps:0.00
rejected_connections:0
sync_full:0
sync_partial_ok:0
sync_partial_err:0
expired_keys:0
expired_stale_perc:0.00
expired_time_cap_reached_count:0
expire_cycle_cpu_milliseconds:12
evicted_keys:0
keyspace_hits:0
keyspace_misses:0
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:0
total_forks:0
migrate_cached_sockets:0
slave_expires_tracked_keys:0
active_defrag_hits:0
active_defrag_misses:0
active_defrag_key_hits:0
active_defrag_key_misses:0
tracking_total_keys:0
tracking_total_items:0
tracking_total_prefixes:0
unexpected_error_replies:0
total_error_replies:0
dump_payload_sanitizations:0
total_reads_processed:2
total_writes_processed:1
io_threaded_reads_processed:0
io_threaded_writes_processed:0

# Replication
role:master
connected_slaves:0
master_failover_state:no-failover
master_replid:18975fd3e7fc9fa18ed9c6960d0ebefa7694513e
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0

# CPU
used_cpu_sys:0.237540
used_cpu_user:0.106971
used_cpu_sys_children:0.000651
used_cpu_user_children:0.002749
used_cpu_sys_main_thread:0.227807
used_cpu_user_main_thread:0.105262

# Modules

# Errorstats

# Cluster
cluster_enabled:0

# Keyspace
127.0.0.1:6379>
```

测试后

```sh
redis-cli --eval tmp/script.lua 10 , 10000

# Memory
used_memory:1645704
used_memory_human:1.57M
used_memory_rss:8949760
used_memory_rss_human:8.54M
used_memory_peak:1746536
used_memory_peak_human:1.67M
used_memory_peak_perc:94.23%
used_memory_overhead:1361888
used_memory_startup:809856
used_memory_dataset:283816
used_memory_dataset_perc:33.96%
allocator_allocated:1739512
allocator_active:2031616
allocator_resident:4710400
total_system_memory:2083807232
total_system_memory_human:1.94G
used_memory_lua:43008
used_memory_lua_human:42.00K
used_memory_scripts:416
used_memory_scripts_human:416B
number_of_cached_scripts:2
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.17
allocator_frag_bytes:292104
allocator_rss_ratio:2.32
allocator_rss_bytes:2678784
rss_overhead_ratio:1.90
rss_overhead_bytes:4239360
mem_fragmentation_ratio:5.58
mem_fragmentation_bytes:7346816
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

```sh
redis-cli --eval tmp/script.lua 20 , 10000

# Memory
used_memory:1724392
used_memory_human:1.64M
used_memory_rss:9310208
used_memory_rss_human:8.88M
used_memory_peak:1746536
used_memory_peak_human:1.67M
used_memory_peak_perc:98.73%
used_memory_overhead:1361888
used_memory_startup:809856
used_memory_dataset:362504
used_memory_dataset_perc:39.64%
allocator_allocated:1797128
allocator_active:2105344
allocator_resident:5074944
total_system_memory:2083807232
total_system_memory_human:1.94G
used_memory_lua:46080
used_memory_lua_human:45.00K
used_memory_scripts:416
used_memory_scripts_human:416B
number_of_cached_scripts:2
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.17
allocator_frag_bytes:308216
allocator_rss_ratio:2.41
allocator_rss_bytes:2969600
rss_overhead_ratio:1.83
rss_overhead_bytes:4235264
mem_fragmentation_ratio:5.53
mem_fragmentation_bytes:7626832
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

```sh
redis-cli --eval tmp/script.lua 100 , 10000

# Memory
used_memory:2604824
used_memory_human:2.48M
used_memory_rss:9932800
used_memory_rss_human:9.47M
used_memory_peak:2626120
used_memory_peak_human:2.50M
used_memory_peak_perc:99.19%
used_memory_overhead:1361888
used_memory_startup:809856
used_memory_dataset:1242936
used_memory_dataset_perc:69.25%
allocator_allocated:2626008
allocator_active:2920448
allocator_resident:5705728
total_system_memory:2083807232
total_system_memory_human:1.94G
used_memory_lua:53248
used_memory_lua_human:52.00K
used_memory_scripts:416
used_memory_scripts_human:416B
number_of_cached_scripts:2
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.11
allocator_frag_bytes:294440
allocator_rss_ratio:1.95
allocator_rss_bytes:2785280
rss_overhead_ratio:1.74
rss_overhead_bytes:4227072
mem_fragmentation_ratio:3.87
mem_fragmentation_bytes:7368992
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

```sh
redis-cli --eval tmp/script.lua 200 , 10000

# Memory
used_memory:3725256
used_memory_human:3.55M
used_memory_rss:11149312
used_memory_rss_human:10.63M
used_memory_peak:3746552
used_memory_peak_human:3.57M
used_memory_peak_perc:99.43%
used_memory_overhead:1361888
used_memory_startup:809856
used_memory_dataset:2363368
used_memory_dataset_perc:81.06%
allocator_allocated:3754992
allocator_active:4091904
allocator_resident:6848512
total_system_memory:2083807232
total_system_memory_human:1.94G
used_memory_lua:74752
used_memory_lua_human:73.00K
used_memory_scripts:416
used_memory_scripts_human:416B
number_of_cached_scripts:2
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.09
allocator_frag_bytes:336912
allocator_rss_ratio:1.67
allocator_rss_bytes:2756608
rss_overhead_ratio:1.63
rss_overhead_bytes:4300800
mem_fragmentation_ratio:3.03
mem_fragmentation_bytes:7465072
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

```sh
redis-cli --eval tmp/script.lua 1000 , 10000

# Memory
used_memory:11725688
used_memory_human:11.18M
used_memory_rss:19660800
used_memory_rss_human:18.75M
used_memory_peak:11746984
used_memory_peak_human:11.20M
used_memory_peak_perc:99.82%
used_memory_overhead:1361888
used_memory_startup:809856
used_memory_dataset:10363800
used_memory_dataset_perc:94.94%
allocator_allocated:11773616
allocator_active:12099584
allocator_resident:15138816
total_system_memory:2083807232
total_system_memory_human:1.94G
used_memory_lua:81920
used_memory_lua_human:80.00K
used_memory_scripts:416
used_memory_scripts_human:416B
number_of_cached_scripts:2
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.03
allocator_frag_bytes:325968
allocator_rss_ratio:1.25
allocator_rss_bytes:3039232
rss_overhead_ratio:1.30
rss_overhead_bytes:4521984
mem_fragmentation_ratio:1.68
mem_fragmentation_bytes:7976128
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

```sh
redis-cli --eval tmp/script.lua 5000 , 10000

# Memory
used_memory:52686120
used_memory_human:50.25M
used_memory_rss:71307264
used_memory_rss_human:68.00M
used_memory_peak:52707416
used_memory_peak_human:50.27M
used_memory_peak_perc:99.96%
used_memory_overhead:1361888
used_memory_startup:809856
used_memory_dataset:51324232
used_memory_dataset_perc:98.94%
allocator_allocated:52715408
allocator_active:53043200
allocator_resident:66637824
total_system_memory:2083807232
total_system_memory_human:1.94G
used_memory_lua:52224
used_memory_lua_human:51.00K
used_memory_scripts:416
used_memory_scripts_human:416B
number_of_cached_scripts:2
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.01
allocator_frag_bytes:327792
allocator_rss_ratio:1.26
allocator_rss_bytes:13594624
rss_overhead_ratio:1.07
rss_overhead_bytes:4669440
mem_fragmentation_ratio:1.35
mem_fragmentation_bytes:18662160
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20504
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

|                   | 测试前    | 场景 1    | 场景 2    | 场景 3    | 场景 4     | 场景 5      | 场景 6      |
|-------------------|--------|---------|---------|---------|----------|-----------|-----------|
| value 字节          | 1      | 10      | 20      | 100     | 200      | 1000      | 5000      |
| key 个数            | 1      | 10000   | 10000   | 10000   | 10000    | 10000     | 10000     |
| 初始 used_memory    | 873584 | 873584  | 873584  | 873584  | 873584   | 873584    | 873584    |
| set 后 used_memory | 873584 | 1645704 | 1724392 | 2604824 | 3725256  | 11725688  | 52686120  |
| 每个 key 大小 (4-3/2) | 0      | 77.212  | 85.0808 | 173.124 | 285.1672 | 1085.2104 | 5181.2536 |
