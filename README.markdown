colorcount is a tool to display the colours used in a PNG image.

Why? Because indexed images take up a lot less space, but it's not so easy to
find out how many colours an image uses *exactly*.

This should be the same as `Colors ➡ Info ➡ Colorcube analysis` in GIMP, just
easier. Create an indexed image with `Image ➡ Mode ➡ Indexed`

Install with `go get zgo.at/colorcount`, which will put the binary in
`~/go/bin`.

![screenshot.png](https://raw.githubusercontent.com/arp242/colorcount/master/screenshot.png)
