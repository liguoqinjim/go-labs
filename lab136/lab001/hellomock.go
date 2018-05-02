package hellomock

type Talker interface {
	SayHello(word string) (response string)
}
