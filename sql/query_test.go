package sql

import (
	"testing"

	"github.com/kr/pretty"
)

func TestQuery(t *testing.T) {
	user := Table("user")
	account := Table("account")
	accountStatus := Table("account_status")

	query := Select(user.AllColumns()).
		From(user).
		InnerJoin(account.Column("user_id"), user.Column("id")). // InnerJoin() joins on the table in the first arg
		Where(account.Column("id").In(
			Select(accountStatus.Column("account_id")).
				From(accountStatus).
				Where(accountStatus.Column("status").InStringArray([]string{"PENDING", "ACTIVE"})),
		)).
		Order(user.Column("name").SortAscending()).
		Limit(100)

	expectedSQL := `
SELECT
  user.*
FROM user 
INNER JOIN account ON account.user_id = user.id
WHERE account.id IN (
SELECT
  account_status.account_id
FROM account_status 
WHERE account_status.status IN ('PENDING', 'ACTIVE')
)
ORDER BY user.name ASC
LIMIT 100
;`
	if expectedSQL != query.ToSQL() {
		pretty.Println("--Expected--")
		pretty.Println(expectedSQL)
		pretty.Println("--Received--")
		pretty.Println(query.ToSQL())
		pretty.Println("---")
		t.Fail()
	}
}
