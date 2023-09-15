# `go-plt`

Plotting in the style of `printf` debugging.

We add bindings to allow plotting on the Go while making "good default" style choices depending on the type and quantity of data.

We buffer data, and dynamically flush it to graph formats, which grow with your data.

For the experimental prototype the only supported backend is `gonum.org/v1/plot`.

## API

Like the `fmt` package, `plt` has several patterns of similar method names.

* `{Noun}`  Plot to the global `Plot` instance.
* `{Noun}f`         ... global `Plot` ... with formatting directives. 
* `P{noun}` Plot to a buffered `Plot` instance.
* `P{noun}f`      ... buffered `Plot` ... with formatting directives.

### Nouns

* `Bar`
* `Box`
* `Candle`
* `Hist`
* `Line`

## Examples

### Line

```
for i := 0; i < 10; i++ {
    plt.Line(i)
}
```