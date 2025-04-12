```
generate id seed
python3 data/edit_id.py

```

run
go run cmd/auth-service/main.go

ดู CPU usage
go tool pprof -http=:3000 http://localhost:8082/debug/pprof/profile
ดู Memory usage
go tool pprof -http=:3001 http://localhost:8082/debug/pprof/heap
