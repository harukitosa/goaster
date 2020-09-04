# goaster

golang の抽象構文木を生成してエディターで閲覧できるコマンドです

# 概要

editor で指定した go ファイルの抽象構文木を閲覧する

```
goaster main.go
```

指定したファイルの抽象構文木がかかれたファイルを生成する

```
goaster main.go -w
```

# 注意

ast/ ディレクトリ以下は go/ast の関数 Fprintf の出力のフォーマットのみ変更したファイルです

元のソースコードは以下の github にあります。
https://github.com/golang/go/tree/master/src/go/ast
