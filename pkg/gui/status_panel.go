package gui

import (
	"strings"

	"github.com/fatih/color"
	"github.com/jesseduffield/gocui"
)

// never call this on its own, it should only be called from within refreshCommits()
func (gui *Gui) refreshStatus() {
	status := "current package"

	gui.g.Update(func(*gocui.Gui) error {
		gui.setViewContent(gui.g, gui.getStatusView(), status)
		return nil
	})
}

func runeCount(str string) int {
	return len([]rune(str))
}

func cursorInSubstring(cx int, prefix string, substring string) bool {
	return cx >= runeCount(prefix) && cx < runeCount(prefix+substring)
}

func (gui *Gui) handleCheckForUpdate(g *gocui.Gui, v *gocui.View) error {
	gui.Updater.CheckForNewUpdate(gui.onUserUpdateCheckFinish, true)
	return gui.createLoaderPanel(gui.g, v, gui.Tr.SLocalize("CheckingForUpdates"))
}

func (gui *Gui) handleStatusClick(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func (gui *Gui) handleStatusSelect(g *gocui.Gui, v *gocui.View) error {
	if gui.popupPanelFocused() {
		return nil
	}

	gui.State.SplitMainPanel = false

	if _, err := gui.g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	gui.getMainView().Title = ""

	magenta := color.New(color.FgMagenta)

	dashboardString := strings.Join(
		[]string{
			lazynpmTitle(),
			"Copyright (c) 2020 Jesse Duffield",
			"Keybindings: https://github.com/jesseduffield/lazynpm/blob/master/docs/keybindings",
			"Config Options: https://github.com/jesseduffield/lazynpm/blob/master/docs/Config.md",
			"Raise an Issue: https://github.com/jesseduffield/lazynpm/issues",
			magenta.Sprint("Become a sponsor (github is matching all donations for 12 months): https://github.com/sponsors/jesseduffield"), // caffeine ain't free
		}, "\n\n")

	return gui.newStringTask("main", dashboardString)
}

func (gui *Gui) handleOpenConfig(g *gocui.Gui, v *gocui.View) error {
	return gui.openFile(gui.Config.GetUserConfig().ConfigFileUsed())
}

func (gui *Gui) handleEditConfig(g *gocui.Gui, v *gocui.View) error {
	filename := gui.Config.GetUserConfig().ConfigFileUsed()
	return gui.editFile(filename)
}

func lazynpmTitle() string {
	return `
	__
	[  |
	 | |  ,--.   ____    _   __  _ .--.  _ .--.   _ .--..--.
	 | | '_\ : [_   ]  [ \ [  ][ \.-. |[ '/'\\ \[ \.-. .-. |
	 | | // | |, .' /_   \ '/ /  | | | | | \__/ | | | | | | |
	[___]'-;__/[_____][\_:  /  [___||__]| ;.__/ [___||__||__]
											\__.'           [__|                  `
}