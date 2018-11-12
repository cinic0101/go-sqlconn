# go-sqlconn
A simple SQL data adapter

Query Sample
```go
    type User struct {
        ID sql.NullString `db:"id"`
        Name sql.NullString `db:"name"`
    }

    dbs := new(sqlconn.Databases)

    db1 := dbs.NewInstance("DB1")
    rows, err := db1.Query(`SELECT id, name FROM User WHERE id = ?`, "ID1")
    defer rows.Close()

    if err != nil {
        return nil, err
    }

    var got []User
    for rows.Next() {
        var u User
        err = rows.Scan(&u.ID, &u.Name)
        if err != nil {
            panic(err)
        }
        got = append(got, u)
    }

    return got, nil
```

Switch between prod and dev environment
```bash
    CONFIG=configs/dbs.dev.yml
```