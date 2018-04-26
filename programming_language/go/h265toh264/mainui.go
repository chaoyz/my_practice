package main

import (
	"github.com/andlabs/ui"
	"github.com/chaoyz/h265toh264/futils"
)

var mainWindow *ui.Window

func main() {
	err := ui.Main(func() {

		mainWindow = ui.NewWindow("h265 to h264 tool", 400, 80, true)

		// outer container
		outBox := ui.NewVerticalBox()
		filePathLabel := ui.NewLabel("")
		outBox.Append(filePathLabel, false)
		selectFileButton := ui.NewButton("select h265 file")
		var filePath string
		selectFileButton.OnClicked(func(*ui.Button) {
			openFileWin := ui.NewWindow("select file", 400, 300, false)
			filePath = ui.OpenFile(openFileWin)
			if filePath != "" {
				// 检查文件格式是否是h265
				checkResult, err := futils.CheckH265Format(&filePath)
				if err != nil || !checkResult {
					// 检查文件不是h265弹窗说明
					warnningWindow("选择文件不是h265文件")
				} else {
					filePathLabel.SetText(filePath)
				}
			} else {
				// 文件没有正确打开或者选择问价为空
				warnningWindow("没有选择需要转码的文件")
			}

		})
		outBox.Append(selectFileButton, false)

		startConvertButton := ui.NewButton("开始转换")
		i := 0
		startConvertButton.OnClicked(func(*ui.Button) {
			// get filePath parent
			// parent, err := futils.GetOutputDir(&filePath)
			// if err != nil {
			// 	warnningWindow("内部出现异常")
			// }
			// _, err = futils.ConvertH265toH264(filePath, *parent)
			// if err != nil {
			// 	fmt.Println(err)
			// 	warnningWindow("转换出现异常")
			// } else {
			// 	warnningWindow("转换完成")
			// }
			if i%2 == 0 {
				startConvertButton.SetText("开始")
			} else {
				startConvertButton.SetText("取消")
			}
			i = i + 1

		})
		outBox.Append(startConvertButton, false)
		mainWindow.SetChild(outBox)
		mainWindow.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		mainWindow.Show()
	})
	if err != nil {
		panic(err)
	}
}

func warnningWindow(text string) {
	win := ui.NewWindow("警告", 200, 50, false)
	box := ui.NewVerticalBox()
	label := ui.NewLabel(text)
	box.Append(label, false)
	okButton := ui.NewButton("确认")
	box.Append(okButton, false)
	win.SetChild(box)
	okButton.OnClicked(func(*ui.Button) {
		win.Hide()
	})
	win.Show()
	win.OnClosing(func(*ui.Window) bool {
		// win.Destroy()
		win.Hide()
		return true
	})
}
