package main

import (
	"strings"

	"github.com/gookit/color"
	"github.com/i582/cfmt/cmd/cfmt"
)

func highlightSQL(sql string) {
	blueBold := color.New(color.FgBlue, color.OpBold)

	replacer := strings.NewReplacer(
		"ANY", blueBold.Sprintf("ANY"),
		"AS", blueBold.Sprintf("AS"),
		"ALL", blueBold.Sprintf("ALL"),
		"ASC", blueBold.Sprintf("ASC"),
		"AND", blueBold.Sprintf("AND"),
		"BETWEEN", blueBold.Sprintf("BETWEEN"),
		"BY", blueBold.Sprintf("BY"),
		"CREATE", blueBold.Sprintf("CREATE"),
		"CASE", blueBold.Sprintf("CASE"),
		"CAST", blueBold.Sprintf("CAST"),
		"COALESCE", blueBold.Sprintf("COALESCE"),
		"COLLATE", blueBold.Sprintf("COLLATE"),
		"COLUMN", blueBold.Sprintf("COLUMN"),
		"COUNT", blueBold.Sprintf("COUNT"),
		"CURRENT_TIME", blueBold.Sprintf("CURRENT_TIME"),
		"CURRENT_TIMESTAMP", blueBold.Sprintf("CURRENT_TIMESTAMP"),
		"DATE", blueBold.Sprintf("DATE"),
		"DATETIME", blueBold.Sprintf("DATETIME"),
		"DO", blueBold.Sprintf("DO"),
		"DELETE", blueBold.Sprintf("UPDATE"),
		"DISTINCT", blueBold.Sprintf("DISTINCT"),
		"EXTRACT", blueBold.Sprintf("EXTRACT"),
		"EXPLAIN", blueBold.Sprintf("EXPLAIN"),
		"EXTEND", blueBold.Sprintf("EXTEND"),
		"ELSE", blueBold.Sprintf("ELSE"),
		"EXISTS", blueBold.Sprintf("EXISTS"),
		"FROM", blueBold.Sprintf("FROM"),
		"FALSE", blueBold.Sprintf("FALSE"),
		"FALSE", blueBold.Sprintf("FALSE"),
		"FULL", blueBold.Sprintf("FULL"),
		"FLOAT", blueBold.Sprintf("FLOAT"),
		"GROUP", blueBold.Sprintf("GROUP"),
		"GLOBAL", blueBold.Sprintf("GLOBAL"),
		"HAVING", blueBold.Sprintf("HAVING"),
		"INSERT", blueBold.Sprintf("INSERT"),
		"IN", blueBold.Sprintf("IN"),
		"INTERVAL", blueBold.Sprintf("INTERVAL"),
		"IS", blueBold.Sprintf("IS"),
		"IS", blueBold.Sprintf("IS"),
		"INTO", blueBold.Sprintf("INTO"),
		"INNER", blueBold.Sprintf("INNER"),
		"JOIN", blueBold.Sprintf("JOIN"),
		"LISTEN", blueBold.Sprintf("LISTEN"),
		"LIKE", blueBold.Sprintf("LIKE"),
		"LIMIT", blueBold.Sprintf("LIMIT"),
		"LEFT", blueBold.Sprintf("LEFT"),
		"NOT", blueBold.Sprintf("NOT"),
		"NOTIFY", blueBold.Sprintf("NOTIFY"),
		"NULL", blueBold.Sprintf("NULL"),
		"OUTER", blueBold.Sprintf("OUTER"),
		"ORDER", blueBold.Sprintf("ORDER"),
		"ON", blueBold.Sprintf("ON"),
		"OR", blueBold.Sprintf("OR"),
		"OFFSET", blueBold.Sprintf("OFFSET"),
		"PROCEDURE", blueBold.Sprintf("PROCEDURE"),
		"PRIMARY", blueBold.Sprintf("PRIMARY"),
		"REFERENCES", blueBold.Sprintf("REFERENCES"),
		"RESET", blueBold.Sprintf("RESET"),
		"RIGHT", blueBold.Sprintf("RIGHT"),
		"SELECT", blueBold.Sprintf("SELECT"),
		"SUBSTRING", blueBold.Sprintf("SUBSTRING"),
		"SUM", blueBold.Sprintf("SUM"),
		"SCHEMA", blueBold.Sprintf("SCHEMA"),
		"TRUE", blueBold.Sprintf("TRUE"),
		"VARCHAR", blueBold.Sprintf("VARCHAR"),
		"VALUES", blueBold.Sprintf("VALUES"),
		"VIEW", blueBold.Sprintf("VIEW"),
		"WITH", blueBold.Sprintf("WITH"),
		"WHERE", blueBold.Sprintf("WHERE"),
		"WHEN", blueBold.Sprintf("WHEN"),
		"UPDATE", blueBold.Sprintf("UPDATE"),
		"UNIQUE", blueBold.Sprintf("UNIQUE"),
		"USING", blueBold.Sprintf("USING"),
		"UNTIL", blueBold.Sprintf("UNTIL"),
		"UNION", blueBold.Sprintf("UNION"),
	)

	sql = replacer.Replace(sql)

	_, _ = cfmt.Println(sql)
}
