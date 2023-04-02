# Go Cache
An implementation of in-memory cache in Go, the applicaiton can
1. hold multiple data
2. keep time limited data in memory
3. support data lock to maintain data consistency
4. evict cache data by LRU policy

### How to use?
Step1: start out the application
```
go run .
```

Step2: visit the endpoint below and view logs in terminal
```
curl http://localhost:8080
```

### APIs
- Set
- Get

### Reference
- [Inmemory cache in golang](https://medium.com/@uvidya16/inmemory-cache-in-golang-7eaf3f5c3d7a)
- [LRU](https://github.com/bmf-san/go-snippets/blob/master/architecture_design/cache/lru.go)
- [Redis Transaction](https://www.youtube.com/watch?v=93MiCYL9OgE)
- [bmf-san/cache](https://github.com/bmf-san/go-snippets/tree/master/architecture_design/cache)
