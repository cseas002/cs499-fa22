package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/ucy-coast/hotel-app/internal/rate"
	"github.com/ucy-coast/hotel-app/pkg/tracing"
)

var (
	port        = flag.Int("port", 8084, "The service port")
	addr        = flag.String("addr", "0.0.0.0", "Address of the service")
	jaegeraddr  = flag.String("jaeger", "jaeger:6831", "Jaeger address")
	jsonDataDir = flag.String("jsondata", "data/medium", "Directory containing json data files")
)

func main() {
	flag.Parse()

	tracer, err := tracing.NewTracer("profile", *jaegeraddr)
	if err != nil {
		log.Panicf("Got error while initializing jaeger agent: %v", err)
	}

	// Initialize Database
	pdb := initializeRateDatabase()

	srv := rate.NewRate(*addr, *port, pdb, tracer)

	if err := srv.Run(); err != nil {
		log.Fatalf("run main error: %v", err)
	}
}
