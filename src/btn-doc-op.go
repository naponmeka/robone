package src

import (
	"fmt"

	"github.com/naponmeka/bsonparser"
	"github.com/naponmeka/robone/connectdb"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"go.mongodb.org/mongo-driver/bson"
)

func registerDocOperationBtn(
	mainWidget *widgets.QWidget,
	resultTreeview *widgets.QTreeView,
	documents *[]bson.M,
	mongoURI *string,
	currentDB *string,
	currentCollection *string,
	globalState *GlobalState,
) {
	editDocBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("editDocBtn", core.Qt__FindChildrenRecursively).Pointer())
	editDocBtn.ConnectClicked(func(bool) {
		selected := findRow(resultTreeview, resultTreeview.CurrentIndex())
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle(fmt.Sprintf("Edit: %d", selected))
		subwin.SetLayout(widgets.NewQHBoxLayout())
		docStr := ""
		if selected < len(*documents) && selected >= 0 {
			docByte, _ := bson.MarshalExtJSON((*documents)[selected], false, true)
			docStr, _ = bsonparser.JsonToBson(string(docByte[:]))
		}
		editLayout := NewEditLayout(docStr)
		subwin.Layout().AddWidget(editLayout)
		dbCollection := connectdb.GetCollection(*mongoURI, *currentDB, *currentCollection)
		RegisterEditLayoutBtn(editLayout, subwin, dbCollection, globalState)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		subwin.Exec()
	})
	viewDocBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("viewDocBtn", core.Qt__FindChildrenRecursively).Pointer())
	viewDocBtn.ConnectClicked(func(bool) {
		// x := globalState.resultTextEdit.TextCursor().Position()
		// debug := fmt.Sprintln(x)
		// widgets.QMessageBox_Information(nil, "OK", debug, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		selected := findRow(resultTreeview, resultTreeview.CurrentIndex())
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle(fmt.Sprintf("View: %d", selected))
		subwin.SetLayout(widgets.NewQHBoxLayout())
		docStr := ""
		if selected < len(*documents) && selected >= 0 {
			docByte, _ := bson.MarshalExtJSON((*documents)[selected], false, true)
			docStr, _ = bsonparser.JsonToBson(string(docByte[:]))
		}
		viewLayout := NewViewLayout(docStr)
		subwin.Layout().AddWidget(viewLayout)
		RegisterViewLayoutBtn(viewLayout, subwin)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		subwin.Exec()
	})
	insertDocBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("insertDocBtn", core.Qt__FindChildrenRecursively).Pointer())
	insertDocBtn.ConnectClicked(func(bool) {
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle("Insert documents")
		subwin.SetLayout(widgets.NewQHBoxLayout())
		insertLayout := NewInsertLayout()
		subwin.Layout().AddWidget(insertLayout)
		dbCollection := connectdb.GetCollection(*mongoURI, *currentDB, *currentCollection)
		RegisterInsertLayoutBtn(insertLayout, subwin, dbCollection, globalState)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(640, 480)
		subwin.Exec()
	})

	deleteDocBtn := widgets.NewQPushButtonFromPointer(mainWidget.FindChild("deleteDocBtn", core.Qt__FindChildrenRecursively).Pointer())
	deleteDocBtn.ConnectClicked(func(bool) {
		selected := findRow(resultTreeview, resultTreeview.CurrentIndex())
		subwin := widgets.NewQDialog(nil, 0)
		subwin.SetWindowTitle(fmt.Sprintf("Delete: %d", selected))
		subwin.SetLayout(widgets.NewQHBoxLayout())
		deleteConfirmLayout := NewConfirmLayout("Confirm delete?")
		subwin.Layout().AddWidget(deleteConfirmLayout)
		dbCollection := connectdb.GetCollection(*mongoURI, *currentDB, *currentCollection)
		RegisterConfirmDeleteLayoutBtn(
			deleteConfirmLayout,
			subwin,
			dbCollection,
			(*documents)[selected],
			globalState,
		)
		subwin.SetModal(true)
		subwin.SetMinimumSize2(100, 100)
		subwin.Exec()
	})
}
