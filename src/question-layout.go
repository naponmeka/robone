package src

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func NewSingleQuestionLayout(question string) *widgets.QWidget {
	var widget = widgets.NewQWidget(nil, 0)
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/single-question.ui")
	file.Open(core.QIODevice__ReadOnly)
	dialogWidget := loader.Load(file, widget)
	file.Close()
	label := widgets.NewQLabelFromPointer(widget.FindChild("questionLabel", core.Qt__FindChildrenRecursively).Pointer())
	label.SetText(question)
	return dialogWidget
}
