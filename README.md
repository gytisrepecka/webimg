# webimg

Image manipulation library for web written in Go.

## Features

Currently webimg library supports following functionality:

* ***Watermarking*** – add PNG logo on top of JPEG file at desired bottom-right offset in pixels, set desired transparency of PNG logo.

## Use in your Go project

Download sources to your GOPATH:
```
go get -d -u code.gyt.is/webimg
```
Mind the arguments: `-d` flag instructs get to stop after downloading the packages; `-u` flag instructs get to use the network to update the named packages and their dependencies.

Include in your application:
```
import (
	"code.gyt.is/webimg"
)
```
And then call function:
```go
// Input image, watermark image, result image, bottom-right offset X, bottom-right offset Y, watermark alpha
doWatermark := webimg.Watermark("smplayer_preferences.jpg", "watermark_inretio-logo.png", "result_img.jpg", 30, 30, 70)
if doWatermark != nil {
  fmt.Println("There was an error watermarking image...")
}
```

Currently the library is able to watermark JPEG image with PNG watermark (which should have transparent background). You can set offset in pixels from bottom-right corner (in example `30, 30`) and transparency of watermark image (in example `70`, scale is 0-255 where 0 is solid and 255 is transparent).

## Source

Sources are published on privately hosted instance of Gitea: [source.gyt.is/gytisrepecka/webimg/](https://source.gyt.is/gytisrepecka/webimg/). To get sources either use `go get` as described above or clone repository:
```
git clone https://source.gyt.is/gytisrepecka/webimg.git
```

## Contact

Follow ***webimg*** blog for news and changelog from fediverse: [@webimg@fedi.dev](https://fedi.dev/webimg/).

If you have any feedback or ideas, drop me an email at [gytis@repecka.com](mailto:gytis@repecka.com) or on Mastodon at [@gytis@mastodon.lt](https://mastodon.lt/@gytis).
