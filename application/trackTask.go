package application

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tracker/domain"
	"tracker/infrastructure/sqlite"

	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

//Track track time in the current task
func Track(minutes int, taskName string, taskID int) {

	t, err := getTaskForTrack(taskName, taskID)

	if err != nil {
		return
	}

	start := time.Now()

	ch := make(chan interface{})

	setupSignal(start, *t)

	ticker := time.Tick(time.Second)

	go func() {
		//sleep until minutes has completed
		time.Sleep(time.Duration(minutes) * time.Minute)
		ch <- true
	}()

	fmt.Printf("\n")
	for {
		select {
		case <-ch:
			terminateAll(start, *t, "timeout")
		case <-ticker:
			diff := time.Now().Sub(start)
			out := time.Time{}.Add(diff)
			fmt.Printf("\r[Ctrl+c for stop] Tracking(%s), Goal time %d minutes -> %s", t.Name, minutes, out.Format("04:05"))
		default:
		}
	}
}

func getTaskForTrack(taskName string, taskID int) (*domain.Task, error) {

	var t *domain.Task
	var err error

	if taskName == "" && taskID == 0 {

		hrepo := sqlite.NewHeadsRepo()

		if t, _, err = hrepo.GetCurrentTask(); err != nil {
			log.Printf("Error :%v\n", err)
			return nil, err
		}

		//return current task(task marked as current in heads table)
		return t, nil
	}

	trepo := sqlite.NewTaskRepo()

	if taskName != "" {

		//search task by name
		if t, err = trepo.GetTaskByName(taskName); err != nil {
			fmt.Printf("Error : %v", err)
			return nil, err
		}

		return t, nil
	}

	//search task by id
	if t, err = trepo.GetTaskByID(uint(taskID)); err != nil {
		fmt.Printf("Error : %v", err)
		return nil, err
	}

	return t, nil
}

func terminateAll(start time.Time, t domain.Task, msg string) {
	endAndRegisterTrack(start, time.Now(), t)
	sendNotification(msg)
	os.Exit(0)
}

func endAndRegisterTrack(start, end time.Time, t domain.Task) {

	min := int(end.Sub(start).Minutes())

	t.Times = append(t.Times, domain.TaskTime{
		Start:           start,
		End:             end,
		DurationMinutes: min,
	})

	fmt.Printf("\n%v minutes complete\n", min)

	trepo := sqlite.NewTaskRepo()

	trepo.RegisterTrack(t)
}

func sendNotification(msg string) {
	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "tracker",
	})

	notify.Push("tracker says", msg, "/home/user/icon.png", notificator.UR_CRITICAL)
}

func setupSignal(start time.Time, t domain.Task) {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func(st time.Time, ta domain.Task) {
		<-c
		fmt.Printf("\nProcess terminated")
		terminateAll(st, ta, "process terminated")
	}(start, t)
}
