package event

import (
	"runtime"
	"unsafe"

	"github.com/anton2920/gofa/syscall"
	"github.com/anton2920/gofa/time"
)

type Queue struct {
	platformEventQueue

	Pinner    runtime.Pinner
	LastPause int64
}

type Request int

const (
	RequestRead Request = (1 << iota)
	RequestWrite
)

type Trigger int

const (
	TriggerLevel Trigger = iota
	TriggerEdge
)

func NewQueue() (*Queue, error) {
	q := new(Queue)
	if err := platformNewEventQueue(q); err != nil {
		return nil, err
	}
	return q, nil
}

func (q *Queue) AddSocket(sock int32, request Request, trigger Trigger, userData unsafe.Pointer) error {
	return platformQueueAddSocket(q, sock, request, trigger, userData)
}

func (q *Queue) AddSignals(sigs ...syscall.Signal) error {
	for i := 0; i < len(sigs); i++ {
		if err := platformQueueAddSignal(q, sigs[i]); err != nil {
			return err
		}
	}
	return nil
}

func (q *Queue) AddTimer(identifier uintptr, duration int, units DurationUnits, userData unsafe.Pointer) error {
	return platformQueueAddTimer(q, identifier, duration, units, userData)
}

func (q *Queue) AppendEvent(event Event) {
	platformQueueAppendEvent(q, event)
}

func (q *Queue) Close() error {
	q.Pinner.Unpin()
	return platformQueueClose(q)
}

func (q *Queue) GetEvents(events []Event) (int, error) {
	return platformQueueGetEvents(q, events)
}

func (q *Queue) HasEvents() bool {
	return platformQueueHasEvents(q)
}

func (q *Queue) Pause(FPS int) {
	now := time.UnixNs()
	durationBetweenPauses := now - q.LastPause
	targetRate := int64(1000.0/float32(FPS)) * 1_000_000

	duration := targetRate - durationBetweenPauses
	if duration > 0 {
		platformQueuePause(q, duration)
		now = time.UnixNs()
	}
	q.LastPause = now
}
