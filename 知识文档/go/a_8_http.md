- 使用http客户端发送请求
- 使用http.Client控制请求头部等
- 使用httputil简化工作

### http服务器的性能分析
- import _ "net/http/pprof"
- 访问/debug/pprof
- 使用go tool pprof分析性能
```
//获得30s性能分析
go tool pprof http://127.0.0.1:8888/debug/pprof/profile
//内存使用请客等
go tool pprof http://127.0.0.1:8888/debug/pprof/heap
```