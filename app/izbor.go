package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/gioapp/gel"
	"github.com/w-ingsolutions/c/pkg/gelook"
	"github.com/w-ingsolutions/c/pkg/latcyr"
)

var (
	IzborVrsteRadovaPanelElement = gel.NewPanel()
	Podvrstaradova               string
	Elementi                     string
)

func (w *WingCal) IzbornikRadovaStrana() func(gtx C) D {
	levo := w.IzborVrsteRadova()
	if w.Element {
		levo = w.PrikazaniElementIzgled()
	}
	return levo
}

func (w *WingCal) IzborVrsteRadova() func(gtx C) D {
	return func(gtx C) D {
		return w.UI.BezMargine.Layout(gtx, func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(w.Nazad()),
				layout.Flexed(1, func(gtx C) D {
					IzborVrsteRadovaPanelElement.PanelObject = w.IzbornikRadova
					IzborVrsteRadovaPanelElement.PanelObjectsNumber = len(w.IzbornikRadova)
					//izborVrsteRadovaPanel := w.Tema.DuoUIpanel()
					//izborVrsteRadovaPanel.ScrollBar = w.Tema.ScrollBar(0)
					//izborVrsteRadovaPanel.Layout(w.Context, IzborVrsteRadovaPanelElement, func(i int, in interface{}) {
					//izborVrsteRadovaPanel.Layout(w.Context, IzborVrsteRadovaPanelElement, func(i int, in interface{}) {
					return izborVrsteRadovaPanel.Layout(gtx, len(w.IzbornikRadova), func(gtx C, i int) D {
						vrstarada := w.IzbornikRadova[i]
						//return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
						//	layout.Rigid(func(gtx C) D {
						//		return w.UI.BezMargine.Layout(gtx, func(gtx C) D {

						//layout.UniformInset(unit.Dp(0)).Layout(w.Context, func() {
						//	layout.Flex{Axis: layout.Vertical}.Layout(w.Context,
						//		layout.Rigid(func() {

						btn := material.Button(w.UI.Tema.T, w.LinkoviIzboraVrsteRadova[i], latcyr.C(vrstarada.Title, w.Cyr))
						btn.CornerRadius = unit.Dp(0)
						btn.Background = gelook.HexARGB(w.UI.Tema.Colors["Gray"])
						if vrstarada.Materijal {
							btn.Background = gelook.HexARGB(w.UI.Tema.Colors["DarkGray"])
						}

						for w.LinkoviIzboraVrsteRadova[i].Clicked() {
							//fmt.Println(i)

							komanda := fmt.Sprint(i + 1)

							if len(w.Putanja) == 1 {
								komanda = fmt.Sprint(i + 1)
								Podvrstaradova = fmt.Sprint(i + 1)
								w.Podvrsta = i + 1
							}
							if len(w.Putanja) == 2 {
								komanda = Podvrstaradova + "/" + fmt.Sprint(i+1)
								Elementi = fmt.Sprint(i + 1)
								w.Roditelj = i + 1
							}
							if len(w.Putanja) == 3 {
								komanda = Podvrstaradova + "/" + Elementi + "/" + fmt.Sprint(i+1)
							}
							if len(w.Putanja) == 1 {
								w.APIpozivIzbornik("radovi/" + komanda)
							}
							if len(w.Putanja) == 2 {
								w.APIpozivElementi("radovi/" + komanda)
							}
							if len(w.Putanja) == 3 {
								w.APIpozivElement("radovi/" + komanda)
								w.Element = true
							}
							if len(w.Putanja) < 3 {
								w.Putanja = append(w.Putanja, vrstarada.Title)
							}
							w.GenerisanjeLinkova(w.IzbornikRadova)
							kolicina.Value = 0
						}
						return btn.Layout(gtx)
					})
				}))
		})
	}
}

//btn.Layout(w.Context, w.LinkoviIzboraVrsteRadova[i])
//layout.Rigid(w.Tema.DuoUIline(w.Context, 0, 0, 0, w.Tema.Colors["Gray"])),
