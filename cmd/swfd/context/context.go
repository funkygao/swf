package context

import (
	"net/http"
	"sync"

	"golang.org/x/net/context"
)

var (
	mutex sync.RWMutex
	data  = make(map[*http.Request]context.Context)
)

func Set(r *http.Request, c context.Context) {
	mutex.Lock()
	data[r] = c
	mutex.Unlock()
}

func Get(r *http.Request) context.Context {
	mutex.RLock()
	if ctx := data[r]; ctx != nil {
		mutex.RUnlock()
		return ctx
	}
	mutex.RUnlock()
	return nil
}

// GetOk returns stored value and presence state like multi-value return of map access.
func GetOk(r *http.Request) (context.Context, bool) {
	mutex.RLock()
	if c, ok := data[r]; ok {
		mutex.RUnlock()
		return c, ok
	}
	mutex.RUnlock()
	return nil, false
}

func GetOrCreate(r *http.Request) (context.Context, bool) {
	c, ok := GetOk(r)
	if ok {
		return c, ok
	}

	mutex.Lock()
	if c, ok := data[r]; ok {
		return c, ok
	}

	c = context.Background()
	data[r] = c
	mutex.Unlock()

	return c, false
}

// Delete removes a value stored for a given key in a given request.
func Delete(r *http.Request) {
	mutex.Lock()
	delete(data, r)
	mutex.Unlock()
}
