package books_test

import (
	// 使用点号导入，把这两个包导入到当前命名空间
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBooks(t *testing.T) {
	// 将Ginkgo的Fail函数传递给Gomega，Fail函数用于标记测试失败，这是Ginkgo和Gomega唯一的交互点
	// 如果Gomega断言失败，就会调用Fail进行处理
	RegisterFailHandler(Fail)

	// 启动测试套件
	// 如果任意 specs（说明）失败了，Ginkgo 会自动使 testing.T 失败
	RunSpecs(t, "Books Suite")
}
