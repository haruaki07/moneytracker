# moneytracker

# Run
```bash 
go run .
```

# Usage
```bash
moneytracker> add -a 20000 -d "new income" --in -t "25/10/2021 14:20"
moneytracker> set <id> -d "out" --in
moneytracker> get [id]
moneytracker> del [id]
moneytracker> clear
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
    set [id] --out

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
    get <id> # get one
    get <id> <id>
```

### del
Delete record.
```bash
  Example: 
    del <id>
    del <id> <id>
```

### clear
Clear all records.
```bash
  Example: 
    clear
```

### exit
Close app.
