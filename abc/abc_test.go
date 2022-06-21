package abc_test

import (
	"context"
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestThing(t *testing.T) {
	// Use cluster 0.
	nsList0, err := client0.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("There are %d namespaces in the cluster 0\n", len(nsList0.Items))

	// Use cluster 1.
	nsList1, err := client1.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("There are %d namespaces in the cluster 1\n", len(nsList1.Items))
}

func TestOtherThing(t *testing.T) {
}
