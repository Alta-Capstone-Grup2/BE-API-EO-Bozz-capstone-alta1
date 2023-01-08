testcity:
	go test ./features/city... -coverprofile=cover.out && go tool cover -html=cover.out

testreview:
	go test ./features/review... -coverprofile=cover.out && go tool cover -html=cover.out

testdiscussion:
	go test ./features/discussion... -coverprofile=cover.out && go tool cover -html=cover.out

run:
	go run main.go