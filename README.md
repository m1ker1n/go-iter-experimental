# GoLand

To remove red highlightings of `"iter"` package usage, you should go `ctrl+alt+S` -> `Go` -> `Build tags` and add `goexperiment.rangefunc` to `Custom tags` input: ![img.png](img.png)

To build your application, you should add env variable: `GOEXPERIMENT=rangefunc` to launch configuration: ![img_1.png](img_1.png)

Now you will see something like this and you can run it: ![img_2.png](img_2.png)