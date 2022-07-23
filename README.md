# SvelteKit embeded in Go

For more details check out the blog post at:
https://www.liip.ch/en/blog/embed-sveltekit-into-a-go-binary

## Requirements

You'll need a recent go and node version.

## Install & Build

```bash
git clone https://github.com/munxar/goapi
cd goapi

go generate ./...
go build

./goapi
```

Now open a browser _http://localhost:3001_ for the api or _http://localhost:3001/admin_ for the admin ui.
