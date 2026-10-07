package controller

import (
	"html/template"
	"path/filepath"
	"runtime"
	"sync"
	"os"

	"cciicc/service"
)

var (
	tmplError        *template.Template
	tmplContentHost  *template.Template
	tmplContentGuest *template.Template
	tmplMain         *template.Template
	tmplHost         *template.Template
	tmplGuest        *template.Template

	tmplMutex sync.RWMutex
	initialized bool
)

// getBaseDir returns the project root directory
func getBaseDir() string {
	// If we're running tests, the current directory might be different
	// Let's first try to find it using runtime.Caller
	_, b, _, ok := runtime.Caller(0)
	if ok {
		dir := filepath.Dir(filepath.Dir(b))
		// Verify if wwwfiles exists in this directory
		if _, err := os.Stat(filepath.Join(dir, "wwwfiles")); err == nil {
			return dir
		}
	}

	// Fallback to current working directory
	dir, _ := os.Getwd()
	return dir
}

func GetTmplError() *template.Template {
	ensureInitialized()
	tmplMutex.RLock()
	defer tmplMutex.RUnlock()
	return tmplError
}

func GetTmplContentHost() *template.Template {
	ensureInitialized()
	tmplMutex.RLock()
	defer tmplMutex.RUnlock()
	return tmplContentHost
}

func GetTmplContentGuest() *template.Template {
	ensureInitialized()
	tmplMutex.RLock()
	defer tmplMutex.RUnlock()
	return tmplContentGuest
}

func GetTmplMain() *template.Template {
	ensureInitialized()
	tmplMutex.RLock()
	defer tmplMutex.RUnlock()
	return tmplMain
}

func GetTmplHost() *template.Template {
	ensureInitialized()
	tmplMutex.RLock()
	defer tmplMutex.RUnlock()
	return tmplHost
}

func GetTmplGuest() *template.Template {
	ensureInitialized()
	tmplMutex.RLock()
	defer tmplMutex.RUnlock()
	return tmplGuest
}

func ensureInitialized() {
	tmplMutex.RLock()
	if initialized {
		tmplMutex.RUnlock()
		return
	}
	tmplMutex.RUnlock()

	tmplMutex.Lock()
	defer tmplMutex.Unlock()

	if initialized {
		return
	}

	basepath := getBaseDir()
	var err error

	tmplError, err = template.ParseFiles(filepath.Join(basepath, "wwwfiles/assets/error/index.html"))
	if err != nil {
		service.CriticalErr(err, "template html 로드 error")
	}

	tmplContentHost, err = template.ParseFiles(filepath.Join(basepath, "wwwfiles/content_host.html"))
	if err != nil {
		service.CriticalErr(err, "template html 로드 content_host")
	}

	tmplContentGuest, err = template.ParseFiles(filepath.Join(basepath, "wwwfiles/content_guest.html"))
	if err != nil {
		service.CriticalErr(err, "template html 로드 content_guest")
	}

	tmplMain, err = template.ParseFiles(filepath.Join(basepath, "wwwfiles/main.html"))
	if err != nil {
		service.CriticalErr(err, "template html 로드 main")
	}

	tmplHost, err = template.ParseFiles(filepath.Join(basepath, "wwwfiles/host.html"))
	if err != nil {
		service.CriticalErr(err, "template html 로드 host")
	}

	tmplGuest, err = template.ParseFiles(filepath.Join(basepath, "wwwfiles/guest.html"))
	if err != nil {
		service.CriticalErr(err, "template html 로드 guest")
	}

	initialized = true
}

// ReloadTemplates allows reloading templates, useful for tests
func ReloadTemplates() {
	tmplMutex.Lock()
	defer tmplMutex.Unlock()
	initialized = false
}
