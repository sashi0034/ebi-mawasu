package core

import (
	"cmp"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type ActorList struct {
	actors []Actor
}

func (l *ActorList) Add(actor Actor) {
	l.actors = append(l.actors, actor)
}

func (l *ActorList) Update() error {
	if len(l.actors) > 1 {
		slices.SortStableFunc(l.actors, func(a, b Actor) int {
			return cmp.Compare(updatePriority(a), updatePriority(b))
		})
	}

	for _, actor := range l.actors {
		if err := actor.Update(); err != nil {
			return err
		}
	}

	l.removeDeadActors()
	return nil
}

func (l *ActorList) Draw(screen *ebiten.Image) {
	if len(l.actors) > 1 {
		slices.SortStableFunc(l.actors, func(a, b Actor) int {
			return cmp.Compare(drawPriority(a), drawPriority(b))
		})
	}

	for _, actor := range l.actors {
		actor.Draw(screen)
	}
}

func (l *ActorList) removeDeadActors() {
	aliveActors := l.actors[:0]
	for _, actor := range l.actors {
		if actor.Alive() {
			aliveActors = append(aliveActors, actor)
		}
	}
	l.actors = aliveActors
}

func updatePriority(actor Actor) int {
	if actor, ok := actor.(UpdatePrioritized); ok {
		return actor.UpdatePriority()
	}
	return 0
}

func drawPriority(actor Actor) int {
	if actor, ok := actor.(DrawPrioritized); ok {
		return actor.DrawPriority()
	}
	return 0
}
