package memento

func ExampleMemento() {
	editor := NewEditor(NewDoc("《孔令飞沉思录》"))
	editor.Append("标题：自我晋升篇")
	editor.Append("标题：自我觉醒篇")
	editor.Append("标题：自我反思篇")

	editor.Delete()
	editor.Show()

	editor.Undo()
	editor.Undo()
	// Output:
	// 打开文档《孔令飞沉思录》
	// ===> 插入操作，文档内容如下：
	// 标题：自我晋升篇
	//
	// ===> 插入操作，文档内容如下：
	// 标题：自我晋升篇
	// 标题：自我觉醒篇
	//
	// ===> 插入操作，文档内容如下：
	// 标题：自我晋升篇
	// 标题：自我觉醒篇
	// 标题：自我反思篇
	//
	// ===> 删除操作，文档内容如下：
	//
	// ===> 撤销操作，文档内容如下：
	// 标题：自我晋升篇
	// 标题：自我觉醒篇
	// 标题：自我反思篇
	//
	// ===> 撤销操作，文档内容如下：
	// 标题：自我晋升篇
	// 标题：自我觉醒篇
}
