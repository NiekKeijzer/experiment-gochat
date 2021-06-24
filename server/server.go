package server

func Init() {
	r := CreateRouter()
	r.Run()
}
