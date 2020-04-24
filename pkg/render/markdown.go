package render

import (
	"github.com/guumaster/tablewriter"

	"github.com/guumaster/hostctl/pkg/types"
)

func NewMarkdownRenderer(opts *TableRendererOptions) TableRenderer {
	table := createTableWriter(opts)

	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetTablePadding("\t") // pad with tabs

	return TableRenderer{
		Type:    Markdown,
		Columns: opts.Columns,
		table:   table,
		opts:    opts,
		meta: &types.Meta{
			Rows: 0,
		},
	}
}
