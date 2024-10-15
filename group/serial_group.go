package group

type SerialGroup struct {
	fs []func() error
}

func NewSerialGroup() *SerialGroup {
	return &SerialGroup{}
}

// Append 尾追加
func (fg *SerialGroup) Append(f func() error) {
	fg.fs = append(fg.fs, f)
}

// Exec 顺序执行
func (fg *SerialGroup) Exec() error {
	for _, f := range fg.fs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
