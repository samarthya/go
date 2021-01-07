package lesson1

import "testing"

func TestLesson(t *testing.T) {
	// Invoke the HelloVersion
	version := HelloVersion() 
	
	// Check the version
	if version != MyVersion {
		t.Errorf("HelloVersion() = %s, must be %s.", version, MyVersion)
	}
}