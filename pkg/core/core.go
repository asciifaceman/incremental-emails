package core

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/asciifaceman/incremental-emails/pkg/data"
)

// New returns a new Game construct
func New(version string) *Game {
	return &Game{
		Version: version,
	}
}

// Game is our core game wrapper from whence all things come
type Game struct {
	Version     string
	Application fyne.App
	Backend     *data.Backend
	Player      *data.Player
	WindowEmail *fyne.Container
}

func (g *Game) Init() {
	g.Application = app.New()
	g.Backend = data.NewBackend("incremental_emails.save", data.DEFAULT_MODE)

	// If no save is present then start new game by defalt without asking to load
	if g.Backend.NoSave() {
		g.NewGame()
	} else {
		fmt.Println("Not implemented")
	}
	g.Run()
}

func (g *Game) Run() {
	g.Application.Run()
}

func (g *Game) NewGame() {
	w := g.Application.NewWindow("New Game")
	g.BuildWindows()

	motd := widget.NewLabel(`Welcome to EmailOS!

	We just need to collect a little information
	to help set your new computer up!`)
	nameEntry := widget.NewEntry()
	nameTouched := false
	nameEntry.SetPlaceHolder("Character Name")

	emailEntry := widget.NewEntry()
	emailTouched := false
	emailEntry.SetPlaceHolder("Character Email")

	continueButton := widget.NewButton("Continue", func() {
		g.Player = data.NewPlayer(nameEntry.Text, emailEntry.Text)
		g.MainGameWindow()
		w.Close()
	})
	continueButton.Disable()

	cont := container.NewVBox(
		motd,
		nameEntry,
		g.quitButton("Quit"),
		continueButton,
	)

	nameEntry.OnChanged = func(content string) {
		if !nameTouched {
			cont.Objects = cont.Objects[:len(cont.Objects)-2]
			cont.Add(emailEntry)
			cont.Add(g.quitButton("Quit"))
			cont.Add(continueButton)
			cont.Refresh()
			nameTouched = true
		}
	}

	emailEntry.OnChanged = func(content string) {
		if !emailTouched {
			continueButton.Enable()
			cont.Refresh()
			emailTouched = true
		}
	}

	w.SetContent(cont)

	w.Canvas().Focus(nameEntry)
	w.Show()
}

func (g *Game) BuildWindows() {
	g.WindowEmail = container.NewVBox(
		g.AnEmail("Hey check this out"),
		g.AnEmail("Big dongers today!"),
		g.AnEmail("Please return and update me the same"),
		g.AnEmail("Free VIAGRA best from the snake himself"),
		g.AnEmail("IRS is looking for you check out why"),
	)
}

func (g *Game) AnEmail(text string) *fyne.Container {
	txt := widget.NewLabel(text)
	received := time.Now().Format("15:04:05 01/02/2006")

	cont := container.NewAdaptiveGrid(5,
		txt,
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewLabel(received),
	)
	return cont
}

func (g *Game) MainGameWindow() {
	w := g.Application.NewWindow(fmt.Sprintf("EmailOS v%s: Incremental Emails", g.Version))
	w.Resize(fyne.NewSize(512+256, 512+256))
	w.CenterOnScreen()

	toolbar := widget.NewToolbar(
		NewToolbarLabel(fmt.Sprintf("Hello, %s!", g.Player.Name)),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.LogoutIcon(), func() { os.Exit(0) }),
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Inbox", g.MainEmailWindow()),
		container.NewTabItem("Browser", widget.NewLabel("Web Browser screen")),
		container.NewTabItem("Profile", widget.NewLabel("Profile screen")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	tabs.SetTabLocation(container.TabLocationTop)

	content := container.NewBorder(toolbar, nil, nil, nil, tabs)
	w.SetContent(content)

	w.Show()
}

func (g *Game) MainEmailWindow() *fyne.Container {

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DeleteIcon(), func() {}),
		widget.NewToolbarAction(theme.MailSendIcon(), func() {}),
		widget.NewToolbarAction(theme.MailAttachmentIcon(), func() {}),
		widget.NewToolbarAction(theme.MailReplyIcon(), func() {}),
		widget.NewToolbarAction(theme.MailForwardIcon(), func() {}),
	)

	content := container.NewBorder(toolbar, nil, nil, nil, container.NewVScroll(g.WindowEmail))
	return content
}

func (g *Game) quitButton(label string) *widget.Button {
	return widget.NewButton(label, func() {
		os.Exit(0)
	})
}
