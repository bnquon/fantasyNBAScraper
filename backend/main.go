package helloworld

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/go-rod/rod"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct {
	Name         string `bson:"name"`
	FantasyPoints string `bson:"fantasyPoints"`
}


func init() {
	functions.HTTP("HelloHTTP", helloHTTP)
}

// RunJob contains your scraping and database logic
func RunJob() error {
	var MONGO_URI = os.Getenv("MONGO_URL")
	var topPlayers []Player

	browser := rod.New().MustConnect()
    defer browser.MustClose() // Ensure the browser closes when done

    // Navigate to the web page
    page := browser.MustPage("https://fantasy.espn.com/basketball/leaders").MustWaitStable()

    // Find the table with the player data
    table := page.MustElements("tbody.Table__TBODY")

    // Find the rows in the table
	playerRows := table[0].MustElements("tr")
    fantasyRows := table[2].MustElements("tr")

	
	for i := 0; i < 10; i++ {
        
		temp := Player{}
		text := playerRows[i].MustText()
    	name := ""

    // Split the text by new line characters to extract the name
    for _, char := range text {
        if char == '\n' {
            break // Stop accumulating name at the first newline character
        }
        name += string(char) // Accumulate the character to the name string
    }
		temp.FantasyPoints = fantasyRows[i].MustText()
		temp.Name = name
		topPlayers = append(topPlayers, temp)

    }

	now := time.Now()

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGO_URI).SetServerAPIOptions(serverApi)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	database := client.Database("NBAFantasyProject")
	playerCollection := database.Collection("DailyFantasyPoints")

	doc := bson.D{{Key: "date", Value: now.Format("01-02-2006")}, {Key: "players", Value: topPlayers}}

	_, err = playerCollection.InsertOne(context.TODO(), bson.D(doc))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted the top players of the day: ", doc)

	var Emails []struct {
		Email string `bson:"email"`
	}

	emailCollection := database.Collection("Emails")
	cursor, err := emailCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.Background(), &Emails); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Emails: ", Emails)

	for _, email := range Emails {
		sendEmail(email.Email, topPlayers, now.Format("01-02-2006"))
	}

	return nil
}

// helloHTTP is your HTTP handler function
func helloHTTP(w http.ResponseWriter, r *http.Request) {

	// Trigger RunJob when the HTTP request comes in
	err := RunJob()
	if err != nil {
		http.Error(w, fmt.Sprintf("Job failed: %v", err), http.StatusInternalServerError)
		return
	}

	// If RunJob is successful
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Job completed successfully!")
}

func sendEmail(email string, dailyResults []Player, date string) {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	formatResults := "Top 10 fantasy players of the day: " + date + "\n\n"

	for index, player := range dailyResults {
		formatResults += fmt.Sprintf("%d. %s - %s\n", index+1, player.Name, player.FantasyPoints)
	}

	message := []byte("Subject: NBA Daily Fantasy Points\n\n" + formatResults)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, message)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	fmt.Println("Email sent!")
}
