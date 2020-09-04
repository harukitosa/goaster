# goaster

golang の抽象構文木を生成してエディターで閲覧できるコマンドです

コマンドライン上で対象ファイルの抽象構文木を簡単に閲覧することができます。

# How to download and build

install

```
go get github.com/harukitosa/goaster
```

or

```
git clone https://github.com/harukitosa/goaster
```

build

```
cd cmd/goaster/
go build
mv goaster $HOME/bin/go
```

※\$HOME/bin/go に PATH を通しておいてください

# 実行例

```
goaster main.go
```

![goaster](https://user-images.githubusercontent.com/44115752/92212447-ef139b00-eecc-11ea-8823-b88e82ab350a.gif)

指定したファイルの抽象構文木がかかれたファイルを生成する

```
goaster main.go -w
```

# メモ

ast/ ディレクトリ以下は go/ast の関数 Fprintf の出力のフォーマットのみ変更したファイルです

元のソースコードは以下の github にあります。
https://github.com/golang/go/tree/master/src/go/ast
