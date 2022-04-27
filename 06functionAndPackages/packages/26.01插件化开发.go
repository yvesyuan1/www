//使用生成的插件
package main

import (
	"github.com/dullgiulio/pingo"
	"log"
)

func main() {
	// Make a new plugin from the executable we created. Connect to it via TCP
	p := pingo.NewPlugin("tcp", "./06functionAndPackages/packages/hello")
	// Actually start the plugin
	p.Start()
	// Remember to stop the plugin when done using it
	defer p.Stop()

	var resp string

	// Call a function from the object we created previously
	if err := p.Call("MyPlugin.SayHello", "YvesYuan", &resp); err != nil {
		log.Print(err)
	} else {
		log.Print(resp)
	}
}
