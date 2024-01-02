# install tailwind
```
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64

chmod +x tailwindcss-macos-arm64

mv tailwindcss-macos-arm64 tailwindcss
```

# start tailwind cli watch sessions
```./tailwindcss -i public/css/main.css -o public/css/output.css --watch```