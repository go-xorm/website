rm -rf output
mkdir output
go build
chmod +x website
mv website ./output/
cp -r ./static/ ./output/static/
cp -r ./templates/ ./output/templates/
cp -r ./conf/ ./output/conf/
