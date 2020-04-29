# MIDAS

[MIDAS](https://github.com/bhatiasiddharth/MIDAS) - edge stream anomaly detection - for Go


## Usage and installation

[Read the docs](https://godoc.org/github.com/steve0hh/midas)

Checkout the `examples` folder for detailed instructions

```go
import (
	"github.com/steve0hh/midas"
	"fmt"
)

func main () {
	src := []int{2,2,3,3,5,5,7,11,1,2}
	dst := []int{3,3,4,4,9,9,73,74,75,76}
	times := []int{1,1,2,2,2,2,2,2,2,2}

	midasAnormScore := midas.Midas(src, dst, times, 2, 769)
	midasRAnormScore := midas.MidasR(src, dst, times, 2, 769, 0.6)

	fmt.Println(midasAnormScore)
	fmt.Println(midasRAnormScore)
}
```

## Resources

- [MIDAS: Microcluster-Based Detector of Anomalies in Edge Streams](https://www.comp.nus.edu.sg/~sbhatia/assets/pdf/midas.pdf)

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- Report bugs
- Fix bugs and submit pull requests
- Write, clarify, or fix documentation
- Suggest or add new features

## TODOs

- [ ] Godocs documentation
- [ ] Add sklearn/keras fit & predict API
- [ ] More examples and tests
- [ ] Make code more efficient
