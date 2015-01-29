package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/minio-io/minio/pkg/server"
)

func main() {
	app := cli.NewApp()
	app.Name = "minio"
	app.Usage = ""
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "http-address,a",
			Value: ":8080",
			Usage: "http address to listen on",
		},
		cli.StringFlag{
			Name:  "cert,c",
			Value: "",
			Usage: "cert.pem",
		},
		cli.StringFlag{
			Name:  "key,k",
			Value: "",
			Usage: "key.pem",
		},
		cli.BoolFlag{
			Name:  "tls,t",
			Usage: "http address to listen on",
		},
		cli.StringFlag{
			Name:  "storage-type,s",
			Value: "file",
			Usage: "valid entries: file,inmemory",
		},
	}
	app.Action = func(c *cli.Context) {
		tls := c.Bool("tls")
		storageTypeStr := c.String("storage-type")
		address := c.String("http-address")
		log.Println(address)
		certFile := c.String("cert")
		keyFile := c.String("key")
		storageType := getStorageType(storageTypeStr)
		serverConfig := server.ServerConfig{
			Address:     address,
			Tls:         tls,
			CertFile:    certFile,
			KeyFile:     keyFile,
			StorageType: storageType,
		}
		server.Start(serverConfig)
	}
	app.Run(os.Args)
}

func getStorageType(input string) server.StorageType {
	switch {
	case input == "file":
		return server.FileStorage
	case input == "inmemory":
		return server.InMemoryStorage
	default:
		{
			log.Println("Unknown storage type:", input)
			log.Println("Choosing default storage type as 'file'..")
			return server.FileStorage
		}
	}
}