testcity:
	go test ./features/city... -coverprofile=cover.out && go tool cover -html=cover.out

run:
	go run main.go