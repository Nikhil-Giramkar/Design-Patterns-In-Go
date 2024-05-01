package main

import "fmt"

func main() {
	snapshotManager := &SnapshotManager{
		snapshotList: make([]*Snapshot, 0),
	}

	virtualMachine := &VirtualMachine{
		state: "A",
	}

	fmt.Printf("VM's current state = %s\n", virtualMachine.getState())
	fmt.Println("Lets create a checkpoint on current state")
	snapshotManager.addSnapshot(virtualMachine.createCheckpoint())

	fmt.Println()

	fmt.Println("Setting current state of VM to B state")
	virtualMachine.setState("B")
	fmt.Printf("VM's current state = %s\n", virtualMachine.getState())
	fmt.Println("Lets create a checkpoint on current state")
	snapshotManager.addSnapshot(virtualMachine.createCheckpoint())

	fmt.Println()

	fmt.Println("Setting current state of VM to C state")
	virtualMachine.setState("C")
	fmt.Printf("VM's current state = %s\n", virtualMachine.getState())
	fmt.Println("Lets create a checkpoint on current state")
	snapshotManager.addSnapshot(virtualMachine.createCheckpoint())

	fmt.Println()

	fmt.Println("Now I wish to restore VM's state to Checkpoint - B")
	virtualMachine.restoreCheckpoint(snapshotManager.getSnapshot(1))
	fmt.Printf("VM's current state = %s\n", virtualMachine.getState())

	fmt.Println()

	fmt.Println("Now I wish to restore VM's state to Checkpoint - A")
	virtualMachine.restoreCheckpoint(snapshotManager.getSnapshot(0))
	fmt.Printf("VM's current state = %s\n", virtualMachine.getState())

}
