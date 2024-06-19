default_arch := amd64
go_build_flags := go build -ldflags="-s -w"
out_dir := ./out
go_main_file := ./cmd/gurl/main.go
project_name := gurl

all:
	GOOS="linux" GOARCH=$(default_arch) $(go_build_flags) -o $(out_dir)/$(project_name) $(go_main_file)
	GOOS="windows" GOARCH=$(default_arch) $(go_build_flags) -o $(out_dir)/$(project_name).exe $(go_main_file)

	tar -czvf $(out_dir)/linux_amd64.tar.gz $(out_dir)/$(project_name)
	zip -r $(out_dir)/windows_amd64.zip $(out_dir)/$(project_name).exe

	rm -f $(out_dir)/$(project_name)
	rm -f $(out_dir)/$(project_name).exe

clean:
	find $(out_dir) -type f ! -name ".gitkeep" -exec rm {} +
