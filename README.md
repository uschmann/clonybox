# Env

Copy .env.example to .env to configure the environment variables.

| Key | Description |
| --- | --- |
| SPOTIFY_ID | Spotify app-id |
| SPOTIFY_SECRET | Spotify app-sercret |
| RFID_TYPE | Type of RFID-Scanner to use. Possible values: stdio, evdev, mfrc522 |
| RFID_EVDEV_FILE | Path to the evdev file to use if RFID_TYPE is evdev. Default: /dev/input/event0 |
| MFRC522_RESET_PIN | Reset-Pin to use for the mfrc522. Default: P1_22 |
| MFRC522_IRQ_PIN | IRQ-Pin to use for the mfrc522. Default: P1_18 |

# Building the frontend
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