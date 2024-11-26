package repo

import (
	"fmt"
	pomodoro "pomodoro/pomo"
	"sync"
)

type inMemoryRepo struct {
	sync.RWMutex
	intervals []pomodoro.Interval
}

func NewInMemoryRepo() *inMemoryRepo {
	return &inMemoryRepo{
		intervals: []pomodoro.Interval{},
	}
}

func (r *inMemoryRepo) Create(i pomodoro.Interval) (int64, error) {
	r.Lock()
	defer r.Unlock()
	i.Id = int64(len(r.intervals) + 1)
	r.intervals = append(r.intervals, i)

	return i.Id, nil
}

func (r *inMemoryRepo) Update(i pomodoro.Interval) error {
	r.Lock()
	defer r.Unlock()

	if i.Id == 0 {
		return fmt.Errorf("%w: %d", pomodoro.ErrInvalidId, i.Id)
	}

	r.intervals[i.Id-1] = i
	return nil
}

func (r *inMemoryRepo) ById(id int64) (pomodoro.Interval, error) {
	r.RLock()
	defer r.RUnlock()
	i := pomodoro.Interval{}
	if id == 0 {
		return i, fmt.Errorf("%w: %d", pomodoro.ErrInvalidId, id)
	}

	i = r.intervals[id-1]
	return i, nil
}

func (r *inMemoryRepo) Last() (pomodoro.Interval, error) {
	r.RLock()
	defer r.Unlock()

	i := pomodoro.Interval{}

	if len(r.intervals) == 0 {
		return i, pomodoro.ErrNoIntervals
	}

	return r.intervals[len(r.intervals)-1], nil
}

func (r *inMemoryRepo) Breaks(n int) ([]pomodoro.Interval, error) {
	r.RLock()
	defer r.RUnlock()

	data := []pomodoro.Interval{}

	for k := len(r.intervals) - 1; k >= 0; k-- {
		if r.intervals[k].Category == pomodoro.Pomodoro {
			continue
		}

		data = append(data, r.intervals[k])

		if len(data) == n {
			return data, nil
		}
	}
	return data, nil
}
