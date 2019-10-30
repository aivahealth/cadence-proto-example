# How to Break This Thing
1. Build it: `$ make`
2. Run it, and notice that it already breaks: `reflect: call of reflect.Value.Interface on zero Value`. This is because the ExampleWorkflow func takes a pointer argument. The DataConverter has nothing to serialize into!
3. Change the ExampleWorkflow arg into a non-pointer and run it: `panic: cannot assign function argument: 2 from type: *main.ExampleMsg to type: main.ExampleMsg`
4. Change the passed-in arg from a pointer into a non-pointer (main.go:47) and run it: `[CustomDataConverter.ToData()] Encoding a normal type: main.ExampleMsg`
5. So, the DataConverter failed to recognize this arg as a `proto.Message`, because it was passed in as a non-pointer.

And this is where the story ends for now...
