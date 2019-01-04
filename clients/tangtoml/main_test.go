package tangtoml

import "log"

// ExampleGetTOML gets the tang.toml file for coins.asia
func ExampleClient_GetTangToml() {
	_, err := DefaultClient.GetTangToml("coins.asia")
	if err != nil {
		log.Fatal(err)
	}
}
