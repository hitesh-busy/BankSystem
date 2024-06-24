package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/BankSystem/service"
	"github.com/go-co-op/gocron"
)

// Scheduler holds the gocron scheduler instance
var Scheduler *gocron.Scheduler

func init() {
	Scheduler = gocron.NewScheduler(time.Now().Local().Location())

	//start the cron scheduler 
	Scheduler.StartAsync()
}

// ScheduleEmail schedules an email to be sent after a specified delay or on a certain day
func ScheduleEmail(email, subject, body string) {
	fmt.Println("sending email ")
	now := time.Now()
	scheduleTime := now.Add(time.Minute * 1).Format("15:04") 

	fmt.Printf("Current time: %v\n", now)
	fmt.Printf("Scheduling job at: %v\n", scheduleTime)

	//only for testing purposes the time is the next minute, can use Exact time as well
	Scheduler.Every(1).Day().At(scheduleTime).Do(func() {
		fmt.Printf("Executing scheduled task at: %v\n", time.Now().Format("15:04"))

		err := service.SendEmail(
			"sharmahitesh472001@gmail.com", // From email
			email,                          // To email
			subject,                        // Subject
			body)                           // Body
		if err != nil {
			log.Printf("Failed to send email to %s: %v", email, err)
		} else {
			log.Printf("Email sent to %s", email)
		}

	})
}

func ScheduleTransactionCalculation() {
	fmt.Println("cron starts ")
	var emailBody string
	now := time.Now()

	//for testing purposes we are calling the function the next minute
	scheduleTime := now.Add(time.Minute * 1).Format("15:04")

	Scheduler.Every(1).Day().At(scheduleTime).Do(func() {
		emailBody = service.CalculateTransactionsForTheDay()

		//transaction cron itself calls email cron.
		//HTML and complex data can be sent as well
		if emailBody != "" {
			fmt.Println("Calling email utility ")
			ScheduleEmail("sharmahitesh472001@gmail.com", "Transaction of the Day", emailBody)
		} else {
			fmt.Println("Cannot send email as email is empty")
		}

	})

}
