package http


type HttpRequest interface {
	DoGet(url string) (body []byte, err error)
}