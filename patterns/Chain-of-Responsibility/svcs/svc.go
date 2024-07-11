package svcs

type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	GotData    bool
	UpdateData bool
	Id         string
}
