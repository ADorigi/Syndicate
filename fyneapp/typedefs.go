package fyneapp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TapCard struct {
	*widget.Card
	OnTapped func()
}

func (tc *TapCard) Tapped(*fyne.PointEvent) {
	if tc.OnTapped != nil {
		tc.OnTapped()
	}
}

// func (tc *TapCard) TappedSecondary(*fyne.PointEvent) {}

func NewTapCard(title string, subTitle string, content fyne.CanvasObject, tapped func()) *TapCard {
	tapCard := &TapCard{widget.NewCard(title, subTitle, content), tapped}
	tapCard.ExtendBaseWidget(tapCard)

	return tapCard
}
