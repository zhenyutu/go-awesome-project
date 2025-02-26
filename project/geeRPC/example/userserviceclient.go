package example

type UserServiceClient struct {
	Hello    func() string
	SayHello func(name string) string
}

func (usc *UserServiceClient) Name() string {
	return "UserService"
}
