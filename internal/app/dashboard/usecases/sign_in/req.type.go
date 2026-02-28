package sign_in

type Request struct {
	Email    string
	Password string
}

type Response struct {
	Token string
}
