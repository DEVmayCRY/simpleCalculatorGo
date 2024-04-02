package main

import (
	"strconv"

	"github.com/DEVmayCRY/simpleCalculatorGo/calculator"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var OptionsMath = []string{"Somar", "Subtrair", "Multiplicar", "Dividir", "Potencia"}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculadora")
	myWindow.Resize(fyne.NewSize(300, 400))

	// Definir os elementos da interface gráfica
	var selectedOption string
	var num1, num2 float64

	firstInput := widget.NewEntry()
	firstInput.Validator = func(text string) error {
		n, err := strconv.ParseFloat(text, 64)
		num1 = n
		return err
	}

	output := widget.NewLabel("0")
	output.Alignment = fyne.TextAlignCenter
	equals := widget.NewButton("=", func() {
		output.SetText(strconv.FormatFloat(calculator.CalculatorMath(num1, num2, selectedOption), 'f', -1, 64))
	})

	optionSelect := widget.NewSelect(OptionsMath, func(value string) {
		selectedOption = value
		equals.Enable()
	})

	optionSelect.PlaceHolder = "Selecione uma Operação"

	secondInput := widget.NewEntry()
	secondInput.Validator = func(text string) error {
		n, err := strconv.ParseFloat(text, 64)
		num2 = n
		return err
	}

	equals.Disable()

	// Organizar os elementos na janela
	content := container.NewVBox(
		firstInput,
		optionSelect,
		secondInput,
		equals,
		output,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
