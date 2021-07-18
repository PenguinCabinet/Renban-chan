# Renban-chan
The cli tool downloading files by continuous URL.

# Get started

## How to download files by continuous URL

In example,It is downloading files,1.jpg~6.jpg, by "http://example.com/".

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

It is making a directory and downloading into it.

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
