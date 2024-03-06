# temporal-workflow

## Usage

### Start the Worker
```bash
  $ go run main.go
```

Output:
```bash
2024/03/06 17:47:27 INFO  No logger configured for temporal client. Created default one.
2024/03/06 17:47:27 INFO  Started Worker Namespace default TaskQueue greeting-tasks WorkerID 22696@TM-ND0131.local@
```
Worker will keep running until it is terminated.

### To Execute the  Workflow using tctl
```bash
  $ tctl workflow start \                                
    --workflow_type GreetSomeone \
    --taskqueue greeting-tasks \
    --workflow_id my-first-workflow \
    --input '"Pulkit"'
```

Output:
```bash
Started Workflow Id: my-first-workflow, run Id: 07414494-9399-4c89-a922-e879cbe36caa
```
When you run the command, it submits your execution request to the cluster, which responds with the Workflow ID, which will be the same as the one you provided, or assigned UUID if you omitted it.

It also displays a Run ID, which uniquely identifies this specific execution of the Workflow.

### Fetch the result
```bash
  $ tctl workflow observe --workflow_id my-first-workflow
```

Output:
```bash
  $ Progress:
  1, 2024-03-06T12:25:46Z, WorkflowExecutionStarted
  2, 2024-03-06T12:25:46Z, WorkflowTaskScheduled
  3, 2024-03-06T12:25:46Z, WorkflowTaskStarted
  4, 2024-03-06T12:25:46Z, WorkflowTaskCompleted
  5, 2024-03-06T12:25:46Z, WorkflowExecutionCompleted

Result:
  Run Time: 1 seconds
  Status: COMPLETED
  Output: ["Hello Pulkit!"]

```

The tctl workflow observe command shows the progress of the Event History of a Workflow Execution.

