package passwd_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/softask-app/lib-go-password/v1/pkg/passwd"
)

func TestPassword_String(t *testing.T) {
	Convey("Password.String", t, func() {
		So(fmt.Sprint(passwd.Password("Hello")), ShouldEqual, "***")
	})
}

func TestPassword_MarshalJSON(t *testing.T) {
	Convey("Password.MarshalJSON", t, func() {
		a, b := json.Marshal(passwd.Password("Goodbye"))
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `"***"`)
	})
}

func TestPassword_MarshalText(t *testing.T) {
	Convey("Password.MarshalText", t, func() {
		a, b := passwd.Password("fizzbuzz").MarshalText()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, "***")
	})
}

func TestPassword_MarshalXML(t *testing.T) {
	Convey("Password.MarshalXML", t, func() {
		a, b := xml.Marshal(passwd.Password("foobar"))
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, "<Password>***</Password>")
	})
}

func TestPassword_MarshalXMLAttr(t *testing.T) {
	Convey("Password.MarshalXMLAttr", t, func() {
		type test struct {
			Foo passwd.Password `xml:"foo,attr"`
		}

		a, b := xml.Marshal(test{"Yo"})
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `<test foo="***"></test>`)
	})
}

func ExamplePassword_String() {
	pass := passwd.Password("Hello World")
	fmt.Println(pass)
	fmt.Println(pass.String())
	// Output:
	// ***
	// ***
}

func ExamplePassword_MarshalJSON() {
	pass := passwd.Password("What you doin")

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(pass)
	// Output:
	// "***"
}

func ExamplePassword_MarshalXML() {
	type tmp struct {
		Pass passwd.Password `xml:"pass"`
	}

	pass := tmp{"XML Time"}

	enc := xml.NewEncoder(os.Stdout)
	enc.Encode(pass)
	// Output:
	// <tmp><pass>***</pass></tmp>
}

func ExamplePassword_MarshalXMLAttr() {
	type tmp struct {
		Pass passwd.Password `xml:"pass,attr"`
	}

	pass := tmp{"What about an attribute?"}

	enc := xml.NewEncoder(os.Stdout)
	enc.Encode(pass)
	// Output:
	// <tmp><pass="***"></pass></tmp>
}
