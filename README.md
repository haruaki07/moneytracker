# moneytrack

# Run
```bash 
go run .
```

# Usage
```bash
moneytrack> add -a 20000 -d "new income" --in -t "25/10/2021 14:20"
moneytrack> set <id> -d "out" --in
moneytrack> get [id]
moneytrack> del [id]
moneytrack> clear
```

## Command List
### add
Add new record.
```bash
  Example: 
    add -a 20000 -d "new income" --in -t "25/10/2021 14:20"

  Arguments:
    -a <amount>         - Amount (required)
    -d <description>    - Description
    -t <datetime>       - Date, format: "dd/mm/yyyy hh:MM" (with double quotes)
    --in                - Set type to income
    --out               - Set type to outcome
```

### set
Update selected record.
```bash
  Example: 
    set <id> --out

  Arguments:
    -a <amount>         - Set amount
    -d <description>    - Set description
    -t <datetime>       - Set date, format: "dd/mm/yyyy hh:MM" (with double quotes)
    --in                - Set type to income
    --out               - Set type to outcome
```

### get
Get all records, pass id if want to get specific record.
```bash
  Example: 
    get # get all
    get --in # get all income
    get <id> <id> # get by id

  Arguments:
    --in                - Get all income
    --out               - Get all outcome
```

### del
Delete record.
```bash
  Example: 
    del <id>
    del <id> <id>
```

### clear
Delete all records.
```bash
  Example: 
    clear
```

### help
Display help.
```bash
  Example: 
    help
    help <command>
```

### exit
Close app.
