package client

type packageTCP struct {
	From head
	Body []byte
}

type head struct {
	From     string
	Filename string
	SHA256   [32]byte
}

type Answer struct {
	res bool
}
