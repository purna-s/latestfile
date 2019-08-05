package latestfile

import (
	"fmt"
    "io/ioutil"
    "os"
    "time"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-latestfile")

// MyActivity is a stub for your Activity implementation
type latestfile struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &latestfile{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *latestfile) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *latestfile) Eval(ctx activity.Context) (done bool, err error) {
	
	dir := ctx.GetInput("Path").(string)
	dt := time.Now()
	
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    var modTime time.Time
    var names string
	var size int64
    for _, fi := range files {
        if fi.Mode().IsRegular() {
            if !fi.ModTime().Before(modTime) {
                if fi.ModTime().After(modTime) {
                    modTime = fi.ModTime()
					size = fi.Size()
                    names = fi.Name()
                }
            }
        }
    }
	diff := dt.Sub(modTime)
	mins := int(diff.Minutes())
    if len(names) > 0 {
        
		ctx.SetOutput("FileName", names)
		ctx.SetOutput("Directory", dir)
		ctx.SetOutput("LastModTime", modTime)
		ctx.SetOutput("MinutesDiff", mins)
		ctx.SetOutput("Size", size)
		fmt.Println("FileName", names)
		fmt.Println("Directory", dir)
		fmt.Println("LastModTime", modTime)
		fmt.Println("MinutesDiff", mins)
		fmt.Println("Size", size)
    }

	activityLog.Debugf("Activity has listed out the latest file Successfully")
	fmt.Println("Activity has listed out the latest file Successfully")
	
	return true, nil
}

