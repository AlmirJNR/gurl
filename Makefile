default_arch := amd64
go_build_flags := go build -ldflags="-s -w"
out_dir := ./out
go_main_file := ./cmd/gurl/main.go

all:
	GOOS="linux" GOARCH=$(default_arch) $(go_build_flags) -o $(out_dir)/linux_amd64 $(go_main_file)
	GOOS="windows" GOARCH=$(default_arch) $(go_build_flags) -o $(out_dir)/windows_amd64.exe $(go_main_file)

clean:
	find $(out_dir) -type f ! -name ".gitkeep" -exec rm {} +
