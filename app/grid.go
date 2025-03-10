package app

import (
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/terminalapi"
)

func newGrid(b *buttonSet, w *widgets, t terminalapi.Terminal) (*container.Container, error) {
	builder := grid.New()

	// Add the first row
	builder.Add(
		grid.RowHeightPerc(30,
			// Add left column
			grid.ColWidthPercWithOpts(30, []container.Option{
				container.Border(linestyle.Light),
				container.BorderTitle("Press Q to Quit"),
			},
				grid.RowHeightPerc(80,
					grid.Widget(w.donTimer),
				),
				grid.RowHeightPercWithOpts(20,
					[]container.Option{
						container.AlignHorizontal(align.HorizontalCenter),
					},
					grid.Widget(w.txtTimer,
						container.AlignHorizontal(align.HorizontalCenter),
						container.AlignVertical(align.VerticalMiddle),
						container.PaddingLeftPercent(49),
					),
				),
			),
			// Add right column
			grid.ColWidthPerc(70,
				grid.RowHeightPerc(80,
					grid.Widget(w.disType, container.Border(linestyle.Light)),
				),
				grid.RowHeightPerc(20,
					grid.Widget(w.txtInfo, container.Border(linestyle.Light)),
				),
			),
		),
	)

	// Add the second row
	builder.Add(
		grid.RowHeightPerc(10,
			grid.ColWidthPerc(50,
				grid.Widget(b.btStart),
			),
			grid.ColWidthPerc(50,
				grid.Widget(b.btPause),
			),
		),
	)

	// Add the third row
	builder.Add(
		grid.RowHeightPerc(60),
	)

	// Build the grid layout and handle errors
	gridOpts, err := builder.Build()
	if err != nil {
		return nil, err
	}

	// Create and return the container
	container, err := container.New(t, gridOpts...)
	if err != nil {
		return nil, err
	}

	return container, nil
}
