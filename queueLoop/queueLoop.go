package queueLoop

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

type Player struct {
	UserId     string
	Elo        int
	TimeJoined time.Time
}

var mu sync.Mutex
var queue []Player

func Add(p Player) {
	mu.Lock()
	queue = append(queue, p)
	mu.Unlock()
}

func Remove(id string) {
	mu.Lock()
	fmt.Println("locked queue to remove player: " + id)
	for i, p := range queue {
		if p.UserId == id {
			queue = append(queue[:i], queue[i+1:]...)
			fmt.Println("queue after removing player: " + id)
		}
	}
	mu.Unlock()
	fmt.Println("unlocked queue after removing: " + id)
}

func remove(id string) {
	for i, p := range queue {
		if p.UserId == id {
			queue = append(queue[:i], queue[i+1:]...)
			break
		}
	}
}

var onMatchFound func(p1, p2 Player)

func SetOnMatchFoundCallback(cb func(p1, p2 Player)) {
	onMatchFound = cb
}

func StartLoop() {
	go func() {
		for {
			mu.Lock()
			p1, p2 := tryToFindAMatch()
			if p1 != nil && p2 != nil {
				player1, player2 := *p1, *p2
				if onMatchFound != nil {
					onMatchFound(*p1, *p2)
				}
				remove(player1.UserId)
				remove(player2.UserId)
			}

			mu.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()
}

func tryToFindAMatch() (*Player, *Player) {
	if len(queue) < 2 {
		return nil, nil
	}

	indices := make([]int, len(queue))
	for i := range queue {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return queue[indices[i]].TimeJoined.Before(queue[indices[j]].TimeJoined)
	})

	now := time.Now()

	for i, idx1 := range indices {
		p1 := &queue[idx1]
		waitTime := now.Sub(p1.TimeJoined).Seconds()
		tolerance := int(100 + waitTime/30*5)
		for _, idx2 := range indices[i+1:] {
			p2 := &queue[idx2]
			fmt.Printf("Checking %s vs %s, Elo diff: %d, tolerance: %d\n",
				p1.UserId, p2.UserId, int(math.Abs(float64(p1.Elo-p2.Elo))), tolerance)
			if p1.UserId != p2.UserId && int(math.Abs(float64(p1.Elo-p2.Elo))) <= tolerance {
				return p1, p2
			}
		}
	}

	return nil, nil
}
