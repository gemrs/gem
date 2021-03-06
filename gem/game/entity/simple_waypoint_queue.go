package entity

import (
	"github.com/gemrs/gem/gem/game/position"
)

/* mob directions */
const (
	DirectionNorth     int = 1
	DirectionNorthEast int = 2
	DirectionEast      int = 4
	DirectionSouthEast int = 7
	DirectionSouth     int = 6
	DirectionSouthWest int = 5
	DirectionWest      int = 3
	DirectionNorthWest int = 0
	DirectionNone      int = -1
)

var directionMap = [3][3]int{
	{DirectionNorthWest, DirectionWest, DirectionSouthWest},
	{DirectionNorth, DirectionNone, DirectionSouth},
	{DirectionNorthEast, DirectionEast, DirectionSouthEast},
}

// A SimpleWaypointQueue trusts the points generated by the client and does
// simple interpolation to determine the next position.
// In future, we might want to create another implementation which performs
// server-side path finding
type SimpleWaypointQueue struct {
	points        []*position.Absolute
	lastDirection int
	direction     int
	running       bool
}

func NewSimpleWaypointQueue() *SimpleWaypointQueue {
	return &SimpleWaypointQueue{
		points:        make([]*position.Absolute, 0),
		lastDirection: DirectionNone,
		direction:     DirectionNone,
	}
}

// Empty determines if there are any points queued
func (q *SimpleWaypointQueue) Empty() bool {
	return len(q.points) == 0
}

// Clear clears the waypoint queue
func (q *SimpleWaypointQueue) Clear() {
	q.points = []*position.Absolute{}
	q.direction = DirectionNone
	q.lastDirection = DirectionNone
}

// Push appends a point to the waypoint queue
func (q *SimpleWaypointQueue) Push(point *position.Absolute) {
	q.points = append(q.points, point)
}

func (q *SimpleWaypointQueue) SetRunning(running bool) {
	q.running = running
}

// Tick advances the waypoint queue, and returns the next position of the mob
func (q *SimpleWaypointQueue) Tick(entity Entity) bool {
	mob := entity.(Movable)

	if len(q.points) == 0 {
		// Nothing to do
		return true
	}

	nextWaypoint := q.points[0]
	current := mob.Position()
	if current.Compare(nextWaypoint) {
		// We've reached a waypoint, dequeue it and continue
		q.points = q.points[1:]
		return q.Tick(mob)
	}

	next := current.NextInterpolatedPoint(nextWaypoint)

	dx, dy, _ := current.DeltaTo(next)

	q.lastDirection = q.direction
	q.direction = directionMap[dx+1][dy+1]

	mob.SetNextStep(next)
	if q.running {
		mob.SetNextStep(next)
	}

	current = mob.Position()
	return current.Compare(nextWaypoint) && len(q.points) == 0
}

// WalkDirection returns the mob's current and (in the case of running) last walking direction
func (q *SimpleWaypointQueue) WalkDirection() (current, last int) {
	return q.direction, q.lastDirection
}

func (q *SimpleWaypointQueue) Interruptible() bool {
	return true
}
