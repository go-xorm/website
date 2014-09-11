rm -rf output_linux_64
mkdir output_linux_64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
chmod +x website
mv website ./output_linux_64/
cp -r ./static/ ./output_linux_64/static/
cp -r ./templates/ ./output_linux_64/templates/
cp -r ./conf/ ./output_linux_64/conf/
