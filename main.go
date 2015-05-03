package main

import (
	"./nntp/"
	"./peers/"
	"./tools/"
	"./upnp/"
	"log"
	"os"
)

// No root. End of story

func init() {

	if (os.Getuid() == 0) || (os.Getgid() == 0) {
		log.Printf("[OMG] %s", "AAAARGH! ROOT! ROOT! ROOOOOT! ")
		os.Exit(1)
	}

	go tools.RotateLogFolder()
	go upnp.AllUpnpOpen()
	go peers.RotateKeysAndCert()

}

// main will only manage local data

func main() {

	tools.SetLogFolder()

	log.Println("[TLS] Initializing engine")
	peers.CreateKeysAndCert(tools.RandSeq(6), tools.RandSeq(8), tools.RandSeq(7))
	log.Println("[TLS] Certs and key created")

	log.Println("[OMG] AVERNO starts now!")

	nntp.NNTP_Frontend()

	//	select {}

	os.Exit(0)

}
