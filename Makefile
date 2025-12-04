# AnhCao 2024 

.PHONY: test 

test: 
	go test ./...

bench:
	go test -bench=. ./... -benchmem
# Benchmark Analysis Guide
# (1) ns/op — latency of each operation
# This tells you how long one call takes.
# If your API handles 1,000 FPS and the function takes 20µs (20,000ns),
# that’s already 20% of your request time budget → might need improvement.
# (2) allocs/op — number of allocations
# This is extremely important in Go.
# 1–2 allocs per small operation → good
# 10+ allocs → probably unnecessary overhead
# 100+ → needs investigation
# Allocations hurt performance because they cause:
# more garbage collection
# more memory pressure
# (3) B/op — bytes allocated per operation
# < 2KB/op → generally fine
# 10KB/op → worth checking
# 100KB/op → likely a problem
# 1MB/op → serious problem