# goaster

golang の抽象構文木を生成してエディターで閲覧できるコマンドです

# 概要

editor で指定した go ファイルの抽象構文木を閲覧する

```
goaster main.go
```

![goaster](https://user-images.githubusercontent.com/44115752/92212447-ef139b00-eecc-11ea-8823-b88e82ab350a.gif)


指定したファイルの抽象構文木がかかれたファイルを生成する

```
goaster main.go -w
```

# 注意

ast/ ディレクトリ以下は go/ast の関数 Fprintf の出力のフォーマットのみ変更したファイルです

元のソースコードは以下の github にあります。
https://github.com/golang/go/tree/master/src/go/ast
