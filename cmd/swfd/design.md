# Design of SWF


        +------------------+                +------------------+
        | ActivityTaskList |                | WorkflowTaskList |
        +------------------+                +------------------+
              |                                       |
              | Task                                  | HistoryEvent(latest state of each task)
              |                                       | Decision
              |                                       |
            worker                                  decider



### Decider

- manages task dependency
- provides input data to activity workers
- schedule tasks
- concurrency
- close the workflow when objective has been completed

### SWF

The role of the Amazon SWF service is to function as a reliable central hub through which data is exchanged between the decider, the activity workers, and other relevant entities such as the person administering the workflow. 

Amazon SWF also maintains the state of each workflow execution, which saves your application from having to store the state in a durable way.
