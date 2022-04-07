/*
	生成一个新插件
	go build -o "hello" 26.插件化开发.go	//生成hello可执行文件作为插件
*/
package main

import "github.com/dullgiulio/pingo"

// Create an object to be exported
type MyPlugin struct{}

// Exported method, with a RPC signature
func (p *MyPlugin) SayHello(name string, msg *string) error {
	*msg = "Hello, " + name
	return nil
}

func main() {
	plugin := &MyPlugin{}

	// Register the objects to be exported
	pingo.Register(plugin)
	// Run the main events handler
	pingo.Run()
}
