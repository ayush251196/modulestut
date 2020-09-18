package modulestut

// func Hello() string {
// 	return quote.Hello()
// }
import "rsc.io/quote/v3"

func V3Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
