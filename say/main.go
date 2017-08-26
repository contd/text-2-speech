package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/contd/text-2-speech/api"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of say backend")
	output := flag.String("o", "output.wav", "wav file that was output")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n\tsay \"text to speak\"\n")
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to %s: %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)

	text := &pb.Text{Text: os.Args[1]}

	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("cound not %s: %v", text.Text, err)
	}
	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("could not write file %s: %v", *output, err)
	}
}
