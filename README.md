# SQL query builder
## Usage
```golang
import "github.com/vojtechrichter/sql_query_builder"

func main() {
	builder := InitQueryBuilder()
	builder.Select("user", "email", "is_admin").From("administration").StartWhere().Equals("user", "admin").EndWhere().OrderBy("createstamp").Limit(10)

	log.Println(builder.GetFinal())
}
```
