package main

type SnapshotManager struct {
	snapshotList []*Snapshot
}

func (sm *SnapshotManager) addSnapshot(s *Snapshot) {
	sm.snapshotList = append(sm.snapshotList, s)
}

func (sm *SnapshotManager) getSnapshot(index int) *Snapshot {
	return sm.snapshotList[index]
}
