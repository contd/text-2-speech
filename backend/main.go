package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/Sirupsen/logrus"
	pb "github.com/contd/text-2-speech/api"
	"github.com/contd/text-2-speech/flite"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create temp file: %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close %s: %v", f.Name(), err)
	}

	if err := flite.TextToSpeech(f.Name(), text.Text); err != nil {
		return nil, fmt.Errorf("flite failed: %v", err)
	}
	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read temp file: %v", err)
	}
	return &pb.Speech{Audio: data}, nil
}

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}
	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	if err := s.Serve(lis); err != nil {
		logrus.Fatal(err)
	}
}

/*
	cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
*/
