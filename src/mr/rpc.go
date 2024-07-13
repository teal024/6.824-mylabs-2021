package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"fmt"
	"os"
	"time"
)
import "strconv"

type Task struct {
	Id           int       // Task id
	Type         string    // Task Type?
	MapInputFile string    // the file to be mapped
	WorkerId     int       // the assigned worker
	Deadline     time.Time // the assigned deadline
}

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.
type Args struct {
	WorkerID int
}

type Reply struct {
	TaskId       int
	TaskType     string
	MapInputFile string // ?
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}

func tmpMapOutFile(workerId int, mapId int, reduceId int) string {
	return fmt.Sprintf("tmp-worker-%d-%d-%d", workerId, mapId, reduceId)
}

func finalMapOutFile(mapId int, reduceId int) string {
	return fmt.Sprintf("mr-%d-%d", mapId, reduceId)
}

func tmpReduceOutFile(workerId int, reduceId int) string {
	return fmt.Sprintf("tmp-worker-%d-out-%d", workerId, reduceId)
}

func finalReduceOutFile(reduceId int) string {
	return fmt.Sprintf("mr-out-%d", reduceId)
}
