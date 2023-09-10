package me

import (
	"io"
	"text/tabwriter"
)

type Me struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

func (m *Me) Print(w io.Writer) {
	tw := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.TabIndent)
	tw.Write([]byte("ID:\t" + m.ID + "\n"))
	tw.Write([]byte("Login:\t" + m.Login + "\n"))
	tw.Write([]byte("Name:\t" + m.Name + "\n"))
	tw.Flush()
}
