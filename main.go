package main

import (
	"facade/facade"
	"log"
)

func main() {
	process := facade.New("alexander", "I'm Sexy and I know it", "token-valido", "user-blog")
	//aqui se hace la llamada a la facade, como se ve, TODO lo que se hace detrás no se ve de lado de cliente :)
	if err := process.Comment(); err != nil {
		log.Fatal(err)
	}

}
