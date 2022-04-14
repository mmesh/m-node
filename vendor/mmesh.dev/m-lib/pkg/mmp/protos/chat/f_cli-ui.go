package chat

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TermView struct {
	Header       *tview.TextView
	IssueInfo    *tview.TextView
	chat         *tview.Flex
	conversation *tview.TextView
	userInput    *tview.InputField

	flex *tview.Flex

	app *tview.Application
}

var View *TermView = nil

func NewView() *TermView {
	return &TermView{
		Header: tview.NewTextView().
			SetDynamicColors(true),
		IssueInfo: tview.NewTextView().
			SetDynamicColors(true),
		chat: tview.NewFlex(),
		conversation: tview.NewTextView().
			SetDynamicColors(true).SetWrap(true),
		userInput: tview.NewInputField().
			SetFieldWidth(0).
			SetLabel("> ").
			SetLabelColor(tcell.ColorLime).
			SetFieldBackgroundColor(tcell.ColorBlack).
			SetFieldTextColor(tcell.ColorSilver),
	}
}

func (v *TermView) Start() {
	v.app = tview.NewApplication()

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack

	v.flex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(v.Header, 5, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(v.IssueInfo, 26, 1, false).
			AddItem(v.chat.SetDirection(tview.FlexRow).
				AddItem(v.conversation, 0, 3, false).
				AddItem(v.userInput, 3, 1, true), 0, 3, false), 0, 1, false)

	issueInfoT1 := "[gray::b]≡ [white::]Issue Details [gray::b]─────"
	issueInfoT2 := "[gray::b]≡ [white::b]Issue Details [gray::b]─────"
	chatT1 := "[gray::b]───── [white::]mmesh service management [gray::b]≡"
	chatT2 := "[gray::b]───── [white::b]mmesh service management [gray::b]≡"
	inputT := fmt.Sprintf("[gray::b]%s", strings.Repeat("─", 50))

	v.IssueInfo.
		SetTitle(issueInfoT1).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetBorderAttributes(tcell.AttrNone).
		SetBorderColor(tcell.ColorBlack)

	v.chat.
		SetTitle(chatT1).
		SetTitleAlign(tview.AlignRight).
		SetBorder(true).
		SetBorderAttributes(tcell.AttrNone).
		SetBorderColor(tcell.ColorBlack)

	v.userInput.
		SetTitle(inputT).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetBorderAttributes(tcell.AttrNone).
		SetBorderColor(tcell.ColorBlack)

	v.IssueInfo.
		SetChangedFunc(func() {
			v.redraw()
		}).
		SetDoneFunc(func(key tcell.Key) {
			switch key {
			case tcell.KeyEsc:
				closeUserInputCh <- struct{}{}
			}
		})
	v.conversation.
		SetChangedFunc(func() {
			v.redraw()
		}).
		SetDoneFunc(func(key tcell.Key) {
			switch key {
			case tcell.KeyEsc:
				closeUserInputCh <- struct{}{}
			}
		})
	v.userInput.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			UserInputQueue <- strings.TrimSpace(v.userInput.GetText())
			v.userInput.SetText("")
		case tcell.KeyEsc:
			closeUserInputCh <- struct{}{}
		}
	})

	// Capture user input
	v.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		case tcell.KeyTab:
			if v.conversation.HasFocus() {
				v.chat.SetTitle(chatT1)
				v.conversation.ScrollToEnd()
				v.userInputFocus()
				v.userInput.SetLabelColor(tcell.ColorLime)
				return nil
			}
			if v.userInput.HasFocus() {
				v.userInput.SetLabelColor(tcell.ColorGray)
				v.issueInfoFocus()
				v.IssueInfo.SetTitle(issueInfoT2)
				return nil
			}
			if v.IssueInfo.HasFocus() {
				v.IssueInfo.SetTitle(issueInfoT1)
				v.IssueInfo.ScrollToBeginning()
				v.conversationFocus()
				v.chat.SetTitle(chatT2)
				return nil
			}
			// case tcell.KeyEsc:
			// 	// Exit the application
			// 	return nil
		}

		return event
	})

	if err := v.app.SetRoot(v.flex, true).SetFocus(v.userInput).Run(); err != nil {
		panic(err)
	}

}

func (v *TermView) Stop() {
	v.app.Stop()
}

func (v *TermView) redraw() {
	v.app.Draw()
}

func (v *TermView) conversationFocus() {
	v.app.SetRoot(v.flex, true).SetFocus(v.conversation)
}

func (v *TermView) issueInfoFocus() {
	v.app.SetRoot(v.flex, true).SetFocus(v.IssueInfo)
}

func (v *TermView) userInputFocus() {
	v.app.SetRoot(v.flex, true).SetFocus(v.userInput)
}
