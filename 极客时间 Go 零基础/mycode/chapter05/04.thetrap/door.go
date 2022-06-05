package main

/**
 * @description
 * @author chengzw
 * @since 2022/5/26
 */

type Door interface {
	Unlock()
	Open()
	Close()
	Lock()
}
