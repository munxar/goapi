# SvelteKit embeded in Go

## Requirements

You'll need a recent go and node version.

## Install & Build

```bash
git clone github.com/munxar/goapi
cd goapi

go generate ./...
go build

./goapi
```

Now open a browser _http://localhost:3001_ for the api or _http://localhost:3001/admin_ for the admin ui.
