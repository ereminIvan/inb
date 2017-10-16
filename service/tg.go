package service


type tgService struct {
	token string
}

func NewTG(token string) *tgService {
	return &tgService{token:token}
}
