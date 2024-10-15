package group

type ForeachGroup struct {
	fs []func() (keep bool)
}

func NewForeachGroup() *ForeachGroup {
	return &ForeachGroup{}
}

// Append 尾追加
func (fg *ForeachGroup) Append(f func() bool) {
	fg.fs = append(fg.fs, f)
}

// Foreach 顺序执行
func (fg *ForeachGroup) Foreach() {
	for _, f := range fg.fs {
		tf := f
		// 如果不继续就退出执行
		if keep := tf(); !keep {
			return
		}
	}
	return
}
