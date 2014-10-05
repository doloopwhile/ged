package main

import (
	"fmt"
	// "github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	// "os"
	// "os/exec"
	// "path"
	// "regexp"
	// "sort"
	// "strings"
)

type View struct {
	window *gtk.Window
}

func NewView(model *Model) *View {
	view := &View{}

	var menuitem *gtk.MenuItem
	gtk.Init(nil)

	view.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	view.window.SetPosition(gtk.WIN_POS_CENTER)
	view.window.SetTitle("GTK Go!")
	view.window.SetIconName("gtk-dialog-info")
	view.window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")

	vbox := gtk.NewVBox(false, 1)

	// Menu Bar
	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	// Menu Items
	cascademenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(cascademenu)
	submenu := gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	menuitem = gtk.NewMenuItemWithMnemonic("_Open")
	menuitem.Connect("activate", func() {
		path, _ := view.OpenGifWithDialog()
		println(path)
	})
	submenu.Append(menuitem)

	menuitem = gtk.NewMenuItemWithMnemonic("E_xit")
	menuitem.Connect("activate", func() {
		gtk.MainQuit()
	})
	submenu.Append(menuitem)

	view.window.Add(vbox)
	view.window.SetSizeRequest(600, 600)

	return view
}

func (view *View) OpenGifWithDialog() (string, error) {
	f := gtk.NewFileFilter()
	f.AddPattern("*.go")

	var path string
	d := gtk.NewFileChooserDialog(
		"Choose File...",
		view.window,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		gtk.STOCK_OK,
		gtk.RESPONSE_ACCEPT)
	d.AddFilter(f)
	d.Response(func() {
		fmt.Println(d.GetFilename())
		d.Destroy()
	})
	d.Run()

	return path, nil
}

func (view *View) Main() {
	view.window.ShowAll()
	gtk.Main()
}
