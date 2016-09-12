package main

import "github.com/ldelossa/vimeoserver/server"

func main() {
	service := server.NewVimeoService()
	service.HTTPServer.ListenAndServe()
}
