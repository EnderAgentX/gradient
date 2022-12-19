package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math"
	"strconv"
)

type matrix [][]float64
type vector []float64

func f1(x, y float64) float64 {
	return x - math.Sin(y) - 1

}

func f2(x, y float64) float64 {
	return x - (y * y) + 1
}

func fp1_x(x float64) float64 {
	return 1
}

func fp1_y(y float64) float64 {
	return math.Cos(y) * -1
}

func fp2_x(x float64) float64 {
	return 1
}

func fp2_y(y float64) float64 {
	return -2 * y
}

func main() {
	newApp := app.New()
	w := newApp.NewWindow("Метод градиентного спуска")
	w.Resize(fyne.NewSize(300, 600))
	w.CenterOnScreen()

	answer := widget.NewLabel("\n")
	label1 := widget.NewLabel("Введите начальное приближение ")
	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	label2 := widget.NewLabel("Введите точность ")
	entry3 := widget.NewEntry()
	btn1 := widget.NewButton("Посчитать", func() {
		Gradient_method(answer, entry1, entry2, entry3)
	})
	w.SetContent(container.NewVBox(
		label1,
		entry1,
		entry2,
		label2,
		entry3,
		btn1,
		answer,
	))

	w.ShowAndRun()

}

func find_W(x, y float64) matrix {
	W := make(matrix, 2)
	for i := range W {
		W[i] = make([]float64, 2)
	}
	W[0][0] = fp1_x(x)
	W[0][1] = fp1_y(y)
	W[1][0] = fp2_x(x)
	W[1][1] = fp2_y(y)
	return W
}

func Gradient_method(answer *widget.Label, entry1, entry2, entry3 *widget.Entry) {
	np1str := entry1.Text
	np2str := entry2.Text
	epsStr := entry3.Text

	answer.Text = "\n"
	answer.SetText(answer.Text)
	xo := make(vector, 2)
	x := make(vector, 2)
	fk := make(vector, 2)
	fk1 := make(vector, 2)
	wk := make(matrix, 2)
	for i := range wk {
		wk[i] = make([]float64, 2)
	}
	wtk := make(matrix, 2)
	for i := range wtk {
		wtk[i] = make([]float64, 2)
	}
	a := make(vector, 2)
	b := make(vector, 2)
	var fkb float64 = 0
	var bb float64 = 0
	np1, err := strconv.ParseFloat(np1str, 64)
	np2, err := strconv.ParseFloat(np2str, 64)
	xo[0] = np1 //0.1
	xo[1] = np2 //0.1
	cnt := 100
	eps, err := strconv.ParseFloat(epsStr, 64) //0.0001
	if err != nil {
		panic(err)
	}
	k := 1
	for k < 100 {
		fk[0] = f1(xo[0], xo[1])
		fk[1] = f2(xo[0], xo[1])

		wk = find_W(xo[0], xo[1])

		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				wtk[i][j] = wk[j][i]
			}
		}

		for j := 0; j < 2; j++ {
			var total float64 = 0
			for h := 0; h < 2; h++ {
				total += wtk[j][h] * fk[h]
			}
			a[j] = total
		}

		for j := 0; j < 2; j++ {
			var total float64 = 0
			for h := 0; h < 2; h++ {
				total += wk[j][h] * a[h]
			}
			b[j] = total
		}

		for h := 0; h < 2; h++ {
			fkb += fk[h] * b[h]
		}

		for h := 0; h < 2; h++ {
			bb += b[h] * b[h]
		}

		lambed := fkb / bb

		for h := 0; h < 2; h++ {
			x[h] = xo[h] - lambed*a[h]
		}

		fk1[0] = f1(x[0], x[1])
		fk1[1] = f2(x[0], x[1])
		answer.Text = answer.Text + "R" + fmt.Sprintf("%v ", k) + " = " + fmt.Sprintf("%9f ", fk1[0]) + "   " + fmt.Sprintf("%9f ", fk1[1]) + "\n"
		answer.SetText(answer.Text)

		var normxo float64 = 0
		for h := 0; h < 2; h++ {
			normxo += xo[h] * xo[h]
		}
		normxo = math.Sqrt(normxo)

		var normxxo float64 = 0
		for h := 0; h < 2; h++ {
			normxxo += (x[h] - xo[h]) * (x[h] - xo[h])
		}
		normxxo = math.Sqrt(normxxo)

		n := normxxo / normxo

		if normxo != 0 && (n) < eps {
			answer.Text = "Iterations " + fmt.Sprintf("%9v ", k) + "\n" + "Solution " + fmt.Sprintf("%9f ", x[0]) + "   " + fmt.Sprintf("%9f ", x[1]) + answer.Text
			k = cnt
		}
		xo[0] = x[0]
		xo[1] = x[1]
		k++
		fkb = 0
		bb = 0
		answer.SetText(answer.Text)

	}
}
