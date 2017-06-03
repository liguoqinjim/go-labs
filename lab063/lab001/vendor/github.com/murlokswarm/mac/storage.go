package mac

/*
#include "storage.h"
#include "sandbox.h"
*/
import "C"
import (
	"os"
	"path/filepath"
	"strings"

	"github.com/murlokswarm/errors"
	"github.com/murlokswarm/log"
)

func resources() string {
	if isAppPackaged() {
		cresources := C.Storage_Resources()
		resourcesName := C.GoString(cresources)
		createDirIfNotExists(resourcesName)
		return resourcesName
	}

	resourcesName := "resources"
	createDirIfNotExists(resourcesName)
	return resourcesName
}

func isAppPackaged() (packaged bool) {
	execName := os.Args[0]
	path, err := filepath.Abs(filepath.Dir(execName))
	if err != nil {
		log.Error(errors.Newf("can't determine if app is packaged: %v", err))
		return
	}

	for _, dir := range strings.Split(path, "/") {
		if strings.HasSuffix(dir, ".app") {
			return true
		}
	}
	return
}

func storage() string {
	if C.Sandbox_IsSandboxed() != 0 {
		defaultName := getHomeDirname()
		createDirIfNotExists(defaultName)
		return defaultName
	}

	defaultName := getSupportDirname()
	createDirIfNotExists(defaultName)
	return defaultName
}

func getHomeDirname() string {
	chomeName := C.Storage_Home()
	return C.GoString(chomeName)
}

func getSupportDirname() string {
	csupportName := C.Storage_Support()
	supportName := C.GoString(csupportName)

	cbundleID := C.Storage_BundleID()
	bundleID := C.GoString(cbundleID)
	if len(bundleID) == 0 {
		wd, err := os.Getwd()
		if err != nil {
			log.Panic(err)
		}
		appname := filepath.Base(wd)
		return filepath.Join(supportName, "dev.murlok", appname)
	}
	return filepath.Join(supportName, bundleID)
}

func createDirIfNotExists(name string) {
	_, err := os.Stat(name)
	if err != nil {
		os.MkdirAll(name, os.ModeDir|0755)
		return
	}
}
