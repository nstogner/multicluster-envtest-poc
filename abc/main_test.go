package abc_test

import (
	"log"
	"os"
	"os/signal"
	"testing"

	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

var (
	client0 *kubernetes.Clientset
	client1 *kubernetes.Clientset
)

func TestMain(m *testing.M) {
	// Call cleanup on Ctrl-C.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			cleanup()
			return
		}
	}()

	// Spin up 2 clusters.
	client0, client1 = setupTestClient(), setupTestClient()

	// Run all tests.
	code := m.Run()
	cleanup()
	os.Exit(code)
}

var cleanupFuncs []func() error

func cleanup() {
	for i, f := range cleanupFuncs {
		log.Printf("calling cleanup func [%v]", i)
		if err := f(); err != nil {
			log.Printf("failed calling cleanup function [%v]: %v", i, err)
		}
	}
}

func setupTestClient() *kubernetes.Clientset {
	env := envtest.Environment{}

	cleanupFuncs = append(cleanupFuncs, env.Stop)

	log.Print("starting test environment")
	restcfg, err := env.Start()
	if err != nil {
		cleanup()
		log.Fatalf("failed to start testenv: %v", err)
	}

	client, err := kubernetes.NewForConfig(restcfg)
	if err != nil {
		cleanup()
		log.Fatalf("failed to create new k8s client: %v", err)
	}

	return client
}
