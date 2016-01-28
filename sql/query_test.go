package sql

import (
	"testing"

	"github.com/kr/pretty"
)

func TestQuery(t *testing.T) {
	people := Table("people")
	account := Table("account")
	avatar := Table("avatar")

	query := Select(
		people.AllColumns(),
		account.Column("email"),
		avatar.Column("url"),
	).
		From(people).
		InnerJoin(account.Column("person_id"), people.Column("id")).
		InnerJoin(avatar.Column("person_id"), people.Column("id")).
		Where(
		people.Column("first_name").Like("rob%").
			Or(people.Column("last_name").Like("rob%"))).
		Order(people.Column("first_name"), Ascending).
		Limit(20)

	expectedSQL := `
SELECT
  people.*,
  account.email,
  avatar.url
FROM people
INNER JOIN account ON account.person_id = people.id
INNER JOIN avatar ON avatar.person_id = people.id
WHERE (people.first_name LIKE 'rob%') OR (people.last_name LIKE 'rob%')
ORDER BY people.first_name ASC
LIMIT 20;
`

	pretty.Println("---")
	pretty.Println(expectedSQL)
	pretty.Println("---")
	pretty.Println(query.ToSQL())
	pretty.Println("---")

	if expectedSQL != query.ToSQL() {
		t.Fail()
	}
}
