package main
import "os"
import "strconv"
import "strings"
import "os/exec"
import "log"
import "math/rand"
import "time"

func main() {
    f, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
	    panic(err)
	}

	defer f.Close()
	uuid := getUuid()

	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	levels := []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NO MATCH"}

	for i := 1; i <= getNumLinesToGen(); i++ {
		randNum := r1.Intn(10)
		var randLevel string

		if randNum < 6 {
			randLevel = levels[randNum]
		} else {
			randLevel = levels[6]
		}

		lineNum := strconv.Itoa(i)

		logLine :=  lineNum + " - " + randLevel + " - " + uuid + "\n"
		if _, err = f.WriteString(logLine); err != nil {
		    panic(err)
		}
	}
}

func getUuid() string {
	out, err := exec.Command("uuidgen").Output()
    if err != nil {
        log.Fatal(err)
    }
    var outStr string = string(out[:])
    return strings.TrimSpace(outStr)
}

func getNumLinesToGen() int {
	args := os.Args[1:]

	if len(args) > 0 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
		    panic(err)
		}
		return i
	} else {
		return 100
	}
}