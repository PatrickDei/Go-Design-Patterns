package mediator

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManager() *StationManager {
	return &StationManager{isPlatformFree: true}
}

func (sm *StationManager) canArrive(t Train) bool {
	if sm.isPlatformFree {
		sm.isPlatformFree = false
		return true
	}
	sm.trainQueue = append(sm.trainQueue, t)
	return false
}

func (sm *StationManager) notifyAboutDeparture() {
	if !sm.isPlatformFree {
		sm.isPlatformFree = true
	}
	if len(sm.trainQueue) > 0 {
		firstTrainInQueue := sm.trainQueue[0]
		sm.trainQueue = sm.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}
