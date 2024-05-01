package main

type VirtualMachine struct {
	state string
}

func (vm *VirtualMachine) createCheckpoint() *Snapshot {
	return &Snapshot{state: vm.state}
}

func (vm *VirtualMachine) restoreCheckpoint(snap *Snapshot) {
	vm.state = snap.getSavedState()
}

func (vm *VirtualMachine) setState(state string) {
	vm.state = state
}

func (vm *VirtualMachine) getState() string {
	return vm.state
}
