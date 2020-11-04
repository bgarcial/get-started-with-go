module hello

go 1.15

replace example.com/greetings => ../greetings

// For production use, you’d publish your modules
// on a server, either inside your company or on the
//internet, and the Go command will download them from there.
//For now, you need to adapt the caller's module so it can
// find the greetings code on your local file system

require example.com/greetings v0.0.0-00010101000000-000000000000

// at the hello/ directory
// > go build
// go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000

// NOTES
// To build the module, Go found the local code in the ../greetings directory,
// then added a require directive to specify that hello is dependent on (requires)
// example.com/greetings. You created this dependency when you imported the greetings package
// (contained in the greetings module) in hello.go.

// The replace directive tells Go where to find the greetings module, because it isn't published yet.

// o reference a published module, a go.mod file would omit the replace directive and use a
// require directive with a tagged version number at the end.
// LIKE THIS require example.com/greetings v1.1.0
