# Go API Server for Vuli

Vuli Movie Delivery API

## Overview
- API version: 3
- Build date: 2018-09-25T01:11:57.183Z
For more information, please visit [http://spry.is](http://spry.is)


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
swagger generate spec -i swagger.yml -o ./swagger.json --scan-models && swagger serve -F=swagger swagger.json
```
