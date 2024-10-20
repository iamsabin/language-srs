package main

import "language-srs/service"

func main() {
	srv := service.NewService()

	srv.CreateEnglishToJapaneseDeck()
}
