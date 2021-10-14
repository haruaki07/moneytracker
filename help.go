package main

import (
	"fmt"
	"strings"
)

type Helps struct {
	helps map[string]Help
}

type Help struct {
	Description string
	Usage       []string
	Example     []string
	Commands    FieldMap
	Arguments   FieldMap
}

type FieldMap map[string]string

func (helps *Helps) Add(key string, h Help) {
	helps.helps[key] = h
}

func (helps *Helps) Print(key string) string {
	help := helps.helps[key]

	return fmt.Sprintf("%s", help.String())
}

func (h *Help) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("  %s\n\n", h.Description))

	if len(h.Usage) > 0 {
		sb.WriteString(
			fmt.Sprintf("  Usage:\n%s\n",
				formatSectionList(h.Usage),
			),
		)
	}

	if len(h.Example) > 0 {
		sb.WriteString(
			fmt.Sprintf("  Example:\n%s\n",
				formatSectionList(h.Example),
			),
		)
	}

	if len(h.Commands) > 0 {
		sb.WriteString(
			fmt.Sprintf("  Commands:\n%s\n",
				formatSectionListMap(h.Commands),
			),
		)
	}

	if len(h.Arguments) > 0 {
		sb.WriteString(
			fmt.Sprintf("  Arguments:\n%s\n",
				formatSectionListMap(h.Arguments),
			),
		)
	}

	return sb.String()
}

func formatSectionList(section []string) string {
	var sb strings.Builder

	for _, str := range section {
		sb.WriteString(fmt.Sprintf("    %s\n", str))
	}

	return sb.String()
}

func formatSectionListMap(section FieldMap) string {
	var cmds []string
	for cmd, desc := range section {
		cmds = append(cmds, fmt.Sprintf("    %-20s%s", cmd, desc))
	}

	return strings.Join(cmds, "\n")
}

var help = new(Helps)

func ShowHelp(params []string) {
	if len(params) <= 0 {
		fmt.Println(help.Print("help"))
		return
	}

  var available = false

  for k, _ := range help.helps {
    if (k == params[0]) {
      available = true
    }
  }

  if available {
    fmt.Println(help.Print(params[0]))
  } else {
    PrintError(fmt.Sprintf("Cannot find help utility for command '%s'", params[0]))
  }
}

func RegisterHelp() {
	help.helps = make(map[string]Help)

	commands := map[string]Help{}

	commands["help"] = Help{
		Description: "Help information",
		Usage:       []string{"moneytrack> [command] [arguments]"},
		Example: []string{
			`moneytrack> add -a 20000 -d "new income" --in -t "25/10/2021 14:20"`,
			`moneytrack> set <id> -d "out" --in`,
			`moneytrack> get [id]`,
			`moneytrack> del [id]`,
      `moneytrack> clear`,
			`moneytrack> help <command>`,
		},
		Commands: FieldMap{
			"add":   `- Add new record`,
			"get":   `- Get records`,
			"set":   `- Update a record`,
			"del":   `- Delete records by id`,
			"clear": `- Delete all records`,
			"exit":  `- Close`,
			"help":  `- Help information`,
		},
	}

	commands["add"] = Help{
		Description: "Add new record",
    Usage: []string{`add [arguments]`},
		Example: []string{
			`add -a 20000 -d "new income" --in -t "25/10/2021 14:20"`,
		},
		Arguments: FieldMap{
			`-a <amount>`:      `- Amount (required)`,
			`-d <description>`: `- Description`,
			`-t <datetime>`:    `- Date, format: "dd/mm/yyyy hh:MM" (with double quotes)`,
			`--in`:             `- Set type to income`,
			`--out`:            `- Set type to outcome`,
		},
	}

  commands["set"] = Help{
    Description: "Update selected record",
    Usage: []string{`set <id> [arguments]`},
    Example: []string{
      `set <id> --out # update record type to outcome`,
    },
    Arguments: FieldMap{
      `-a <amount>`:      `- Set amount`,
      `-d <description>`: `- Set description`,
      `-t <datetime>`:    `- Set date, format: "dd/mm/yyyy hh:MM" (with double quotes)`,
      `--in`:             `- Set type to income`,
      `--out`:            `- Set type to outcome`,
    },
  }
  commands["get"] = Help{
    Description: "Get all records, pass id if want to get specific record",
    Usage: []string{`get [arguments]`},
    Example: []string{
      `get # get all`,
      `get --in # get all income`,
      `get <id> <id> # get by id`,
    },
    Arguments: FieldMap{
      `--in`:             `- Get all income`,
      `--out`:            `- Get all outcome`,
    },
  }

  commands["del"] = Help{
    Description: "Delete record",
    Usage: []string{`del [id]`},
  }

  commands["clear"] = Help{
    Description: "Delete all record",
    Usage: []string{`clear`},
  }

	for k, v := range commands {
		help.Add(k, v)
	}
}
