package config


type Configurations struct {
	Server Server
	Profile Profile
	Test_variable string
}

type Server struct {
	Port int
}

type Profile struct {
	Active string
}