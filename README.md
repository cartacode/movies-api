# Go API Server for Vuli

Vuli Movie Delivery API

## Overview
- API version: 1
- Build date: N/A
For more information, please visit [Drone Build Page](https://drone.vuli.io/VuliTv/go-movie-api)


### Running the server
To run the server, follow these simple steps:

```
go run main.go

http://127.0.0.1:3001
```


### Generating Docs
You first need to install swagger

```
download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
curl -o /usr/local/bin/swagger -L'#' "$download_url"
chmod +x /usr/local/bin/swagger
```

Generate the json file and serve it up on your browser

```
swagger generate spec -i swagger.yml -o ./swagger.json --scan-models
swagger serve -F=swagger swagger.json
```
