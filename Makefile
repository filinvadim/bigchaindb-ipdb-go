gen:
	rm -rf models
	rm -rf client
	./_swagger_linux_amd64 generate client -f swagger.yml -A bigchain-client
	goimports -w ./client || true
