package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":9000")
	fmt.Println("Serveur démarré sur :9000")
	for {
		conn, _ := ln.Accept()
		go func(c net.Conn) {
			defer c.Close()
			buf, err := io.ReadAll(c)
			if err != nil {
				return
			}
			img, _, err := image.Decode(bytes.NewReader(buf))
			if err != nil {
				return
			}
			b := img.Bounds()
			w, h := b.Max.X, b.Max.Y
			bw := blackWhiteParallel(extractPixelsParallel(img, w, h), w, h)
			jpeg.Encode(c, pixelsToImage(bw), nil)
			fmt.Println("Requête complétée")
		}(conn)
	}
}
