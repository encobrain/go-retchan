package retchan

func New () *Chan {
	return &Chan{
		put_ch: make(chan *putData),
	}
}

// Should be called on new go routine because waits putData
func (c *Chan) Handle (handler Handler) {
	for data := range c.put_ch {
		func() {
			var ret interface{}

			defer func() {
				if ret == nil { ret = recover() }

				data.ret<- ret
			}()

			ret = handler(data.value)
		}()
	}
}

func (c *Chan) Put (value interface{}) (ret interface{}, err error) {
	data := &putData{
		value,
		make(chan interface{}),
	}

	c.put_ch<-data
	ret = <-data.ret

	err,ok := ret.(error)

	if ok { ret = nil }

	return
}

func (c *Chan) Close () {
	close(c.put_ch)
}
