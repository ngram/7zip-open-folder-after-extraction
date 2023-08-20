# 7zip-open-folder-after-extraction

Utility that automatically opens the extraction folder when double-clicking a file when using 7-Zip

## Requirements

* Windows10, 11
* 7-Zip
* Go 1.2.x

## Installation

### 1. Build and install package

```console
$ git clone https://github.com/ngram/7zip-open-folder-after-extraction.git
$ cd 7zip-open-folder-after-extraction
$ go build -ldflags -H=windowsgui -o .\bin\7zip-ofae.exe .\main.go
$ cp .\bin\7zip-ofae.exe YOUR_BINARY_LOCATION/
```

### 2. Change default application for .zip

Refere to follow manual to change .zip default application

* https://support.microsoft.com/en-us/windows/change-default-programs-in-windows-e5d82cad-17d1-c53b-3505-f10a32e1894d