package retchan

type putData struct {
	value 		interface{}
	errStack 	bool
	ret 		chan interface{}
}
