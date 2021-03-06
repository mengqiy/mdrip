package util

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// getProcessGroupId purports to get a process group Id common to all
// subprocesses of its pid argument.
//
// There should be a better way to do this.
//
// Goal is to be able to support killing any subprocesses created by
// RunInSubShell.  At the moment, its up to authors to clean up after
// themselves.
func GetProcesssGroupId(pid int) (int, error) {
	//  /bin/ps -o pid,pgid,rgid,ppid,cmd
	//  /bin/ps -o pgid=12492 --no-headers
	cmdOut, execErr := exec.Command(
		"/bin/ps", "--pid", strconv.Itoa(pid), "-o", "pgid", "--no-headers").Output()
	groupId := strings.TrimSpace(string(cmdOut))
	if execErr != nil || len(groupId) < 1 {
		return 0, errors.New(
			"Unable to yank groupId from ps command: " + groupId + " " + execErr.Error())
	}
	pgid, convErr := strconv.Atoi(groupId)
	if convErr != nil {
		return 0, convErr
	}
	return pgid, nil
}

// check reports the error fatally if it's non-nil.
func Check(msg string, err error) {
	if err != nil {
		glog.Fatal(errors.Wrap(err, msg))
	}
}

// An attempt to kill any and all child processes.
func killProcesssGroup(pgid int) {
	killer := exec.Command("/bin/kill", "-TERM", "--", fmt.Sprintf("-%v", pgid))
	killer.Start()
}

// Convert tabs, newlines, etc. to normal blanks.
func convertBadWhiteSpaceToBlanks(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
			return ' '
		default:
			return r
		}
	}, s)
}

var leading = regexp.MustCompile("^[0-9]+_")

// Drop leading numbers and underscores.
func DropLeadingNumbers(s string) string {
	r := leading.FindStringIndex(s)
	if r == nil {
		return s
	}
	return s[r[1]:]
}

// Convert long multi-line string to a short one-line sample.
func SampleString(incoming string, max int) string {
	s := len(incoming)
	if s > max {
		s = max
	}
	return convertBadWhiteSpaceToBlanks(strings.TrimSpace(incoming[:s]))
}

// Returns a string of length n with only spaces.
func Spaces(n int) string {
	if n < 1 {
		return ""
	}
	return fmt.Sprintf("%"+strconv.Itoa(n)+"s", " ")
}
