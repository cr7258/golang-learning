package main

/**
 * @description
 * @author chengzw
 * @since 2022/5/26
 */

type Assets struct {
	assets []Asset
}

func (a *Assets) DoStartWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Unlock()
			d.Open()
		}
	}
}
func (a *Assets) DoStopWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok {
			d.Close()
			d.Lock()
		}
	}
}
