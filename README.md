# Convertwords

Words Conversion Designed for Asian Languages.

### Install

`go get github.com/ArsFy/convertwords@v0.0.2`

### Example

```go
func main() {
	root := New()
	root.Insert("蘋果", "水果")
	root.Insert("香蕉", "水果")
	root.Insert("蘿蔔", "蔬菜")

	input := "我喜歡蘋果和香蕉，但是不喜歡蘿蔔。"
	output := ConvertWords(input, root)
	println(output)
    // 我喜歡水果和水果，但是不喜歡蔬菜
}
```