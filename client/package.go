package client

type packageTCP struct {
    From head
    Body []byte
}

type head struct {
    From string
    Mod bool
    Filename string
}