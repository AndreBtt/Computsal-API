package team

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("oi")
	code := m.Run()

	os.Exit(code)
}

func TestCreateTeam(t *testing.T) {

	fmt.Println("ola")

	// a := app.App{}
	// a.Initialize("root", "andre1995", "Computsal")
	// defer a.DB.Close()
	// a.Run(":8080")

	// payload := []byte(`{"name":"test team", "photo": "www.url.com.br", "group":1}`)

	// req, _ := http.NewRequest("POST", "/team", bytes.NewBuffer(payload))
	// response := executeRequest(req)

	// checkResponseCode(t, http.StatusCreated, response.Code)

	// var m map[string]interface{}
	// json.Unmarshal(response.Body.Bytes(), &m)

	// if m["name"] != "test team" {
	// 	t.Errorf("Expected team name to be 'test team'. Got '%v'", m["name"])
	// }

	// if m["photo"] != "www.url.com.br" {
	// 	t.Errorf("Expected photoUrl to be 'www.url.com.br'. Got '%v'", m["photo"])
	// }

	// // the group is compared to 1.0 because JSON unmarshaling converts numbers to
	// // floats, when the target is a map[string]interface{}
	// if m["group"] != 1.0 {
	// 	t.Errorf("Expected user group to be '1'. Got '%v'", m["id"])
	// }

}
