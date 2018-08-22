package go_vagrant

import (
	"fmt"
	"os"
	"testing"
)

func newMockVagrantClient() *VagrantClient {
	return &VagrantClient{}
}

func assertArguments(t *testing.T, args []string, expected ...string) {
	if len(args) != len(expected) {
		t.Fatalf("Expected %v args; got %v", len(expected), len(args))
	}
	for i, arg := range args {
		if arg != expected[i] {
			t.Errorf("Expected arg %v to be '%v'; got %v", i, expected[i], arg)
		}
	}
}

func Example() {
	client, err := NewVagrantClient("example")
	if err != nil {
		fmt.Println("Got error while creating client:", err)
		os.Exit(-1)
	}

	// Let's start bringing up the vm
	upcmd := client.Up()
	upcmd.Verbose = true
	fmt.Println("Bringing up the vm")
	if err := upcmd.Start(); err != nil {
		fmt.Println("Error bringing up vm:", err)
		os.Exit(-1)
	}

	// while we're waiting, let's get version info
	vercmd := client.Version()
	if err := vercmd.Run(); err != nil {
		fmt.Println("Error retrieving version info:", err)
	}

	// now wait for up to finish
	if err := upcmd.Wait(); err != nil {
		fmt.Println("Error waiting for up:", err)
		os.Exit(-1)
	}

	fmt.Println("\n\nInstalled Vagrant version:", vercmd.InstalledVersion)

	// Get vm status
	statuscmd := client.Status()
	if err := statuscmd.Run(); err != nil {
		fmt.Println("Error getting status:", err)
	} else {
		for vm, status := range statuscmd.Status {
			fmt.Printf("%v: %v\n", vm, status)
		}
	}

	// Destroy vm
	if err := client.Destroy().Run(); err != nil {
		fmt.Println("Error destroying vm:", err)
		os.Exit(-1)
	}
}
