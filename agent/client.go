package agent

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"github.com/pkg/errors"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

type gopsProc struct {
	pid        int
	executable string
	file       string
}

func (p *gopsProc) String() string {
	return fmt.Sprintf("%v (pid %v)", p.executable, p.pid)
}

func GetGopsProcesses() ([]*gopsProc, error) {
	var result []*gopsProc
	matches, err := filepath.Glob(path.Join(os.TempDir(), fmt.Sprintf("%v.*.sock", SockPrefix)))
	if err != nil {
		return nil, err
	}
	re, err := regexp.Compile(`.*\.(?P<PID>\d+)\.sock`)
	if err != nil {
		return nil, err
	}
	for _, match := range matches {
		res := re.FindSubmatch([]byte(match))
		if len(res) != 0 {
			pid, err := strconv.Atoi(string(res[1]))
			if err != nil {
				return nil, err
			}
			// process is alive and is reachable by us
			if process, err := os.FindProcess(pid); err == nil {
				if err = process.Signal(syscall.Signal(0)); err == nil {
					proc, err := ps.FindProcess(pid)
					if err != nil {
						return nil, err
					}
					result = append(result, &gopsProc{
						pid:        pid,
						executable: proc.Executable(),
						file:       match,
					})
				}
			}
		}
	}
	return result, nil
}

func GetUnixSockForPid(pid int) string {
	return path.Join(os.TempDir(), fmt.Sprintf("%v.%v.sock", SockPrefix, pid))
}

func tooManyProcsError(procs []*gopsProc) error {
	var list []string
	for _, v := range procs {
		list = append(list, v.String())
	}
	builder := &strings.Builder{}
	builder.WriteString("too many gops-agent process found, including [")
	builder.WriteString(strings.Join(list, ", "))
	builder.WriteString("]")
	return errors.New(builder.String())
}

// ParseGopsAddress tries to parse the target string, be it remote host:port
// or local process's PID or executable name
func ParseGopsAddress(args []string) (string, error) {
	if len(args) == 0 {
		return "", nil
	}
	target := args[0]
	if strings.Contains(target, ":") {
		return target, nil
	}

	// try to find port by pid then, connect to local
	num, err := strconv.Atoi(target)
	if err == nil {
		filePath := GetUnixSockForPid(num)
		if _, err := os.Stat(filePath); err == nil {
			return "unix:" + filePath, nil
		}
		return "tcp:127.0.0.1:" + target, nil
	}

	return target, nil
}

func MakeRequest(addr string, signal byte, params []byte, out io.Writer) error {
	network := "tcp"
	if addr == "" {
		network = "unix"
		procs, err := GetGopsProcesses()
		if err != nil {
			return err
		}
		if len(procs) == 0 {
			return errors.New("no gops-agent processes found")
		}
		if len(procs) > 1 {
			return tooManyProcsError(procs)
		}
		addr = procs[0].file
	} else {
		procs, err := GetGopsProcesses()
		found := false
		if err == nil {
			var filtered []*gopsProc
			for _, proc := range procs {
				if strings.Contains(proc.executable, addr) {
					filtered = append(filtered, proc)
				}
			}
			if len(filtered) > 1 {
				return tooManyProcsError(filtered)
			}
			if len(filtered) == 1 {
				found = true
				network = "unix"
				addr = filtered[0].file
			}
		}

		if !found {
			parts := strings.SplitN(addr, ":", 2)
			if len(parts) == 2 {
				network = parts[0]
				addr = parts[1]
			}
		}
	}

	conn, err := net.Dial(network, addr)
	if err != nil {
		return err
	}

	var buf []byte
	buf = append(buf, Magic...)
	buf = append(buf, signal)
	buf = append(buf, params...)

	if _, err := conn.Write(buf); err != nil {
		return err
	}

	_, err = io.Copy(out, conn)
	return err
}
