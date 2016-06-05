package main

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// EdgeScroller is a System that allows for scrolling when the cursor is near the edges of
// the window
type EntityEdgeScroller struct {
	ScrollSpeed    float32
	EdgeMargin     float64
	SpaceComponent *common.SpaceComponent
}

func (*EntityEdgeScroller) Priority() int          { return common.EdgeScrollerPriority }
func (*EntityEdgeScroller) Remove(ecs.BasicEntity) {}

// TODO: Warning doesn't get the cursor position
func (c *EntityEdgeScroller) Update(dt float32) {
	pos := c.SpaceComponent.Position
	//maxX, maxY := engo.CanvasWidth(), engo.CanvasHeight()
	fmt.Printf("maxX: %0.2f, maxY: %0.2f, pos: %v\n", maxX, maxY, pos)

	incremental := true
	if float64(pos.X) < c.EdgeMargin {
		engo.Mailbox.Dispatch(common.CameraMessage{Axis: common.XAxis, Value: -c.ScrollSpeed * dt, Incremental: incremental})
	} else if float64(pos.X) > float64(maxX)-c.EdgeMargin {
		engo.Mailbox.Dispatch(common.CameraMessage{Axis: common.XAxis, Value: c.ScrollSpeed * dt, Incremental: incremental})
	}

	if float64(pos.Y) < c.EdgeMargin {
		engo.Mailbox.Dispatch(common.CameraMessage{Axis: common.YAxis, Value: -c.ScrollSpeed * dt, Incremental: incremental})
	} else if float64(pos.Y) > float64(maxY)-c.EdgeMargin {
		engo.Mailbox.Dispatch(common.CameraMessage{Axis: common.YAxis, Value: c.ScrollSpeed * dt, Incremental: incremental})
	}
}
