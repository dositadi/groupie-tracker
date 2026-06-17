package main

import (
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/cmd/app"
)

func main() {
	a := app.App{}

	a.Run()

	/* openCage := opencage.New("a45e2bfd61d04e13b6504d106de3db70", *jsonlog.New(os.Stdout, jsonlog.INFO))
	geo, err := openCage.FetchGeolocation("Frauenplan Weimar, Germany")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(geo) */
}
