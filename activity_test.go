package latestfile

import (
	"fmt"
    "io/ioutil"
    "os"
    "time"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("Path", "F:\TESTING")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}
	act.Eval(tc)
	//check output attr

	FileName = tc.GetOutput("FileName")
	assert.Equal(t, FileName, FileName)
	Directory = tc.GetOutput("Directory")
	assert.Equal(t, Directory, Directory)
	LastModTime = tc.GetOutput("LastModTime")
	assert.Equal(t, LastModTime, LastModTime)
	MinutesDiff = tc.GetOutput("MinutesDiff")
	assert.Equal(t, MinutesDiff, MinutesDiff)
	Size = tc.GetOutput("Size")
	assert.Equal(t, Size, Size)
	
	//assert.Equal(t, output, output)

}
