package quicktask

import (
	"errors"
	"strings"
	"sync"

	"github.com/lithammer/shortuuid"
)

var (
	quicktaskMutex             = sync.RWMutex{}
	quickTask                  = make(map[string]*Quicktask)
	ErrTaskDoesNotExist        = errors.New("task does not exist")
	taskTypes                  = make(map[string]*TaskType)
	ErrTaskTypeDoesNotExist    = errors.New("task type does not exist")
	ErrTaskHandlerDoesNotExist = errors.New("task handler does not exist")
)

func CreateQuicktask(Site, Tasks_Quantity, Profiles, Accounts, Email, Proxylist, Payment_Method, Credit_Card, Other string) string {
	quicktaskMutex.Lock()
	defer quicktaskMutex.Unlock()

	id := shortuuid.New()

	quickTask[id] = &Quicktask{
		Site:           strings.ToLower(Site),
		Tasks_Quantity: Tasks_Quantity,
		Profiles:       Profiles,
		Accounts:       strings.ToLower(Accounts),
		Email:          strings.ToLower(Email),
		Proxylist:      strings.ToLower(Proxylist),
		Payment_Method: strings.ToLower(Payment_Method),
		Credit_Card:    Credit_Card,
		Other:          Other,
	}

	return id
}

// DoesTaskExist checks if a Quicktask exists
func DoesTaskExist(id string) bool {
	quicktaskMutex.RLock()
	defer quicktaskMutex.RUnlock()
	_, ok := quickTask[id]
	return ok
}

// GetQuicktask gets a quicktask
func GetQuicktask(id string) (*Quicktask, error) {
	if !DoesTaskExist(id) {
		return &Quicktask{}, ErrTaskDoesNotExist
	}

	quicktaskMutex.RLock()
	defer quicktaskMutex.RUnlock()

	return quickTask[id], nil
}
