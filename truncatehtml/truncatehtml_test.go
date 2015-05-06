package truncatehtml

import "testing"

// TestTruncateHtml performs some basic sanity checks of TruncateHtml.
func TestTruncateHtml(t *testing.T) {
  cases := []struct {
      in string
      limit int
      ellipsis string
      want string
  }{
    {
      "",
      0,
      "",
      "",
    },
    {
      "",
      5,
      "",
      "",
    },
    {
      "123",
      0,
      "",
      "",
    },
    {
      "123",
      2,
      "",
      "12",
    },
    {
      "123",
      3,
      "",
      "123",
    },
    {
      "1234",
      3,
      "",
      "123",
    },
    {
      "<b>123</b>",
      5,
      "",
      "<b>123</b>",
    },
    {
      "<b>12345</b>",
      5,
      "",
      "<b>12345</b>",
    },
    {
      "<b>1234567</b>",
      5,
      "",
      "<b>12345</b>",
    },
    {
      "<b>Monty Python</b>",
      5,
      "",
      "<b>Monty</b>",
    },
    {
      "<img />",
      5,
      "",
      "<img />",
    },
    {
      "<img>",
      5,
      "",
      "<img>",
    },
    {
      "<h1><u>test<img blah blah>ing 1 2 3</u></h1>",
      5,
      "",
      "<h1><u>test<img blah blah>i</u></h1>",
    },
    {
      "123<h1><u> 456 <img blah blah> 789 012</u></h1>",
      7,
      "",
      "123<h1><u> 456 <img blah blah> 7</u></h1>",
    },
    {
      "<h1><u>😄u n i 😄 c😄o😄d😄e</u></h1>",
      5,
      "",
      "<h1><u>😄u n i 😄</u></h1>",
    },
    {
      "<h1><u>1234567</u></h1>",
      5,
      "...",
      "<h1><u>12345...</u></h1>",
    },
  }

  for _, c := range cases {
    out, err := TruncateHtml([]byte(c.in), c.limit, c.ellipsis)
    got := string(out)
    if err != nil {
      t.Errorf("Got error calling TruncateHtml(%q, 5, \"\"). Wanted: %q. Error:", c.in, c.want, err.Error())
    }
    if got != c.want {
      t.Errorf("TruncateHtml(%q, %d, %q) == %q, want %q", c.in, c.limit, c.ellipsis, got, c.want)
    }
  }
}