# Env

| Key | Description |
| --- | --- |
| SPOTIFY_ID | Spotify app-id |
| SPOTIFY_SECRET | Spotify app-sercret |

# Buildung the frontend
```bash
cd clonybpx-frontend
npm install
npm run build
```

Note: *You have to rename **dist/index.html** to **dist/index.html** after buildung. 

# Crosscompile to raspberry-pi 3
```bash
sudo apt-get update -y
sudo apt install gcc-arm-linux-gnueabi

CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-extldflags=-static" -tags sqlite_omit_load_extension . 
```

# Allow binding to port 80
```bash
sudo setcap CAP_NET_BIND_SERVICE=+eip clonybox
```