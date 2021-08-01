# Renban-chan
* このツールは連続したURLから画像をダウンロードすることができます。
* このツールはページに乗っている画像をすべてダウンロードできます。
* このツールはCLIツールです。

[ダウンロード](https://github.com/PenguinCabinet/Renban-chan/releases/latest)

# Get started

## 連続したURLからファイルをダウンロード

"http://example.com/"から1.jpg~6.jpgをダウンロードします。

```shell
/home/yours> Renban-chan i
URL template(http://example.com/imgs/**.jpg):http://example.com/imgs/**.jpg
Start index:1
End index:6
downloading http://example.com/imgs/1.jpg
downloading http://example.com/imgs/2.jpg
downloading http://example.com/imgs/3.jpg
downloading http://example.com/imgs/4.jpg
downloading http://example.com/imgs/5.jpg
downloading http://example.com/imgs/6.jpg
```

ディレクトリを作り、指定したうえでダウンロードすることもできます。

```shell
/home/yours> Renban-chan i pictures
URL template(http://example.com/imgs/**.jpg):http://example.com/imgs/**.jpg
Start index:1
End index:6
downloading http://example.com/imgs/1.jpg
downloading http://example.com/imgs/2.jpg
downloading http://example.com/imgs/3.jpg
downloading http://example.com/imgs/4.jpg
downloading http://example.com/imgs/5.jpg
downloading http://example.com/imgs/6.jpg
/home/yours> ls pictures
1.jpg
2.jpg
3.jpg
4.jpg
5.jpg
6.jpg
```

## How to download files in the homepage

"http://example.com/home.html"のすべての画像ファイルをダウンロードします。

```shell
/home/yours> Renban-chan ia pictures
URL(http://example.com/home.html):http://example.com/home.html
downloading http://example.com/imgs/Apple.jpg
downloading http://example.com/imgs/BANANA.jpg
downloading http://example.com/imgs/HAPPY.jpg
/home/yours> ls pictures
1.jpg
2.jpg
3.jpg
```

