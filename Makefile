testuser:
	go test ./features/user... -coverprofile=cover.out && go tool cover -html=cover.out

testauth:
	go test ./features/auth... -coverprofile=cover.out && go tool cover -html=cover.out

run:
	go run main.go