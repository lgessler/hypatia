rm ../dist/linux-x64/hypatia 2>/dev/null
rm ../dist/linux/hypatia 2>/dev/null
rm ../dist/windows-x64/hypatia.exe 2>/dev/null
rm ../dist/windows/hypatia.exe 2>/dev/null
rm ../dist/osx-x64/hypatia 2>/dev/null
rm ../dist/osx/hypatia 2>/dev/null
cd src
GOOS=linux GOARCH=amd64 go build -o ../dist/linux-x64/hypatia
echo "Built Linux x64"
GOOS=linux GOARCH=386 go build -o ../dist/linux/hypatia
echo "Built Linux x86"
GOOS=windows GOARCH=amd64 go build -o ../dist/windows-x64/hypatia.exe
echo "Built Windows x64"
GOOS=windows GOARCH=386 go build -o ../dist/windows/hypatia.exe
echo "Built Windows x86"
GOOS=darwin GOARCH=amd64 go build -o ../dist/osx-x64/hypatia
echo "Built OS X x64"
GOOS=darwin GOARCH=386 go build -o ../dist/osx/hypatia
echo "Built OS X x86"
cd ..
