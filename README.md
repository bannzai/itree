# itree
itree can be interactive file browse and exec any command.
<img width=100% src="https://user-images.githubusercontent.com/10897361/67949494-2705de00-fc2b-11e9-86f3-fd85f4019057.gif" />

## Install

### go get
Required [go](https://github.com/golang/go).

```shell
$ go get github.com/bannzai/itree
```

### Releases
You can get executable file from [latest release](https://github.com/bannzai/itree/releases) for each environment (e.g macOS, Linux, Windows).

## Usage
```shell
$ ./itree --help                       
itree displayed file system tree and command interactively about file system.

Usage:
	itree [options]

The options are:
	--path=$PATH specified start path. default is ./
	--help       displayed help message for itree
```

### Command
Press `?` on itree, to see all commands.

```
(c) Copy selected node file path
(C) Copy selected node absolute file path
(r) Rename file node
(o) '$ open $FILE_PATH'
(n) New file
(N) New directory
(e) Open current node with $EDITOR. Default is vim
(i) appear information for current node
(/) change mode for file search
(?) help message for usage itree
```

## LICENSE
**itree** is available under the MIT license. See the LICENSE file for more info.

