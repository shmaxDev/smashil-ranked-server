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

func StartLoop() {
	go func() {
		for {
			mu.Lock()
			p1, p2 := tryToFindAMatch()
			if p1 != nil && p2 != nil {
				player1, player2 := *p1, *p2
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

	copied := make([]Player, len(queue))
	copy(copied, queue)

	sort.Slice(copied, func(i, j int) bool {
		return copied[i].TimeJoined.Before(copied[j].TimeJoined)
	})

	now := time.Now()

	for i, p1 := range copied {
		waitTime := now.Sub(p1.TimeJoined).Seconds()
		tolerance := int(10 + waitTime/30*5)
		fmt.Println(tolerance)
		for j := i + 1; j < len(copied); j++ {
			p2 := copied[j]
			fmt.Println(p1.Elo, p2.Elo)
			if (p1.UserId != p2.UserId) && int(math.Abs(float64(p1.Elo-p2.Elo))) <= tolerance {
				return &p1, &p2
			}
		}
	}

	return nil, nil
}
