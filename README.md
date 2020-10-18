# GopherStack
Simple routine to monitor goroutines on the stack.


GopherStack simplifies the output from the pprof.Lookup("goroutine"):
```
goroutine profile: total 7
1 @ 0x43f50b 0x408a6a 0x40880b 0xb753e9 0x472171
#	0xb753e8	main.wsDebugWriter+0x58	/home/user/Development/vflow/cmd/vflowapp/debugHandler.go:75

1 @ 0x529a8c 0x896ff9 0x95fd36 0xb7a700 0x472171
#	0x529a8b	fmt.Sprint+0xeb					/usr/local/go/src/fmt/print.go:247
#	0x896ff8	github.com/gobuffalo/packr/v2/plog.Debug+0x1a8	/home/user/go/pkg/mod/github.com/gobuffalo/packr/v2@v2.8.0/plog/plog.go:19
#	0x95fd35	github.com/gobuffalo/packr/v2.New+0x1e5		/home/user/go/pkg/mod/github.com/gobuffalo/packr/v2@v2.8.0/box.go:51
#	0xb7a6ff	main.debugServer+0x5f				/home/user/Development/vflow/cmd/vflowapp/debugServer.go:83

1 @ 0x9ed365 0x9ecff9 0x9e8220 0xb7b107 0x43f148 0x472171
#	0x9ed364	runtime/pprof.writeRuntimeProfile+0x104	/usr/local/go/src/runtime/pprof/pprof.go:694
#	0x9ecff8	runtime/pprof.writeGoroutine+0xb8	/usr/local/go/src/runtime/pprof/pprof.go:656
#	0x9e821f	runtime/pprof.(*Profile).WriteTo+0x9f	/usr/local/go/src/runtime/pprof/pprof.go:329
#	0xb7b106	main.main+0x476				/home/user/Development/vflow/cmd/vflowapp/main.go:50
#	0x43f147	runtime.main+0x1c7			/usr/local/go/src/runtime/proc.go:203

1 @ 0xb84ea1 0x472171
#	0xb84ea0	main.statsRunner+0x0	/home/user/Development/vflow/cmd/vflowapp/svstepstats.go:54

1 @ 0xb93241 0x472171
#	0xb93240	main.vflowServer+0x0	/home/user/Development/vflow/cmd/vflowapp/vflowserver.go:21

1 @ 0xb98af1 0x472171
#	0xb98af0	main.main.func1+0x0	/home/user/Development/vflow/cmd/vflowapp/main.go:46
```
And simplifies it to single lines, the number of goroutines and the highest level function name:
```
GoRoutine Dump
	 2 ==> net/http.(*conn).serve
	 1 ==> main.wsDebugWriter
	 1 ==> main.readHandler
	 1 ==> main.main.func1
	 1 ==> main.statsRunner
	 1 ==> runtime.main
	 1 ==> main.debugServer
	 1 ==> main.vflowServer
	 1 ==> github.com/klaxxon/GopherStack.Run.func1
 ```
If you pass in showChanges as true, it will compare the previous list with the new one and show the changes:
```
2020/10/18 18:17:23 Routine net/http.(*persistConn).readLoop added 1
2020/10/18 18:17:23 Routine main.(*call).handleCall added 1
2020/10/18 18:17:23 Routine main.(*call).vflow added 1
2020/10/18 18:17:23 Routine database/sql.(*DB).connectionOpener added 2
2020/10/18 18:17:23 Routine database/sql.(*DB).connectionResetter added 2
2020/10/18 18:17:23 Routine net/http.(*persistConn).writeLoop added 1
2020/10/18 18:17:28 Routine net/http.(*persistConn).readLoop ended
2020/10/18 18:17:28 Routine net/http.(*persistConn).writeLoop ended

```

To use:
```
import (
  log
  gopherstack "github.com/klaxxon/GopherStack"
 )
 
func main() {
  gopherstack.Run("/home/user/Development/vflow", 5, true)
  .
  .
  .
  .
}
```
Passing in the path of the projects base directory will allow the function to return the highest function in the package, if it exists.

