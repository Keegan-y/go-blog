package article

type Article interface {
	Get(url string) (string,string)
}
