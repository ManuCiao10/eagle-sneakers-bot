package task

import (
	"errors"
	"reflect"
)

var (
	ErrTaskHandlerDoesNotExist = errors.New("task handler does not exist")
)

// GetHandler gets a handler by handler name
func (t *TaskType) GetHandler(handlerName TaskState) (reflect.Value, error) {
	handler, ok := t.handlers[string(handlerName)]

	if !ok {
		return reflect.Value{}, ErrTaskHandlerDoesNotExist
	}

	return handler, nil
}

// GetInternalType gets the internal type
func (t *TaskType) GetInternalType() reflect.Type {
	return t.internalType
}

// SetFirstHandlerState sets the first handler state
func (t *TaskType) SetFirstHandlerState(firstHandlerState TaskState) {
	t.firstHandlerState = firstHandlerState
}

// RegisterTaskType register task type
func RegisterTaskType(registerSiteName string) *TaskType {
	taskTypes[registerSiteName] = &TaskType{
		handlers: make(TaskReflectMap),
	}

	return taskTypes[registerSiteName]
}

// GetFirstHandlerState gets the first handler state
func (t *TaskType) GetFirstHandlerState() TaskState {
	return t.firstHandlerState
}

// HasHandlers check if there are handlers
func (t *TaskType) HasHandlers() bool {
	return len(t.handlers) > 0
}

// GetTaskType gets a task type
func GetTaskType(taskType string) (*TaskType, error) {
	if !DoesTaskTypeExist(taskType) {
		return &TaskType{}, ErrTaskTypeDoesNotExist
	}
	return taskTypes[taskType], nil
}

// DoesTaskTypeExist check if task type exists
func DoesTaskTypeExist(taskType string) bool {
	_, ok := taskTypes[taskType]
	return ok
}

// DoesTaskExist checks if a task exists
func DoesTaskExist(id string) bool {
	taskMutex.RLock()
	defer taskMutex.RUnlock()
	_, ok := tasks[id]
	return ok
}

// GetTask gets a task
func GetTask(id string) (*Task, error) {
	if !DoesTaskExist(id) {
		return &Task{}, ErrTaskDoesNotExist
	}

	taskMutex.RLock()
	defer taskMutex.RUnlock()

	return tasks[id], nil
}

func (t *TaskType) addHandler(handlerName TaskState, handler interface{}) {
	t.handlers[string(handlerName)] = reflect.ValueOf(handler)
}

// AddHandlers adds multiple handles to a task type
func (t *TaskType) AddHandlers(handlers TaskHandlerMap) {
	for handlerName, handler := range handlers {
		if t.internalType == nil {
			handleTypes := reflect.TypeOf(handler)
			// func (t *task.Task, internal *SiteInternal) task.TaskState

			// we want the second one because the first one (0 index) will be task.Task type
			handleType := handleTypes.In(0)

			t.internalType = handleType
		}

		t.addHandler(handlerName, handler)
	}
}
