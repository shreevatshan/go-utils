package database

import (
	"reflect"
	"testing"
)

func Test_getDatabaseOperation(t *testing.T) {

	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want []Operation
	}{
		{
			name: "test1",
			args: args{
				query: "SELECT * FROM mytable",
			},
			want: []Operation{
				{
					Name:  "SELECT",
					Table: "mytable",
				},
			},
		},
		{
			name: "test2",
			args: args{
				query: "INSERT INTO mytable VALUES (1, 2, 3)",
			},
			want: []Operation{
				{
					Name:  "INSERT",
					Table: "mytable",
				},
			},
		},
		{
			name: "test3",
			args: args{
				query: "UPDATE mytable SET column1 = 1",
			},
			want: []Operation{
				{
					Name:  "UPDATE",
					Table: "mytable",
				},
			},
		},
		{
			name: "test4",
			args: args{
				query: "CREATE TABLE mytable (column1 INT)",
			},
			want: []Operation{
				{
					Name:  "CREATE",
					Table: "mytable",
				},
			},
		},
		{
			name: "test5",
			args: args{
				query: "ALTER TABLE mytable ADD column1 INT",
			},
			want: []Operation{
				{
					Name:  "ALTER",
					Table: "mytable",
				},
			},
		},
		{
			name: "test6",
			args: args{
				query: "CALL myprocedure()",
			},
			want: []Operation{
				{
					Name:  "CALL",
					Table: "myprocedure",
				},
			},
		},
		{
			name: "test7",
			args: args{
				query: "SELECT * FROM Customers WHERE Country='Germany'; SELECT * FROM Suppliers WHERE Country='USA';",
			},
			want: []Operation{
				{
					Name:  "SELECT",
					Table: "Customers",
				},
				{
					Name:  "SELECT",
					Table: "Suppliers",
				},
			},
		},
		{
			name: "test8",
			args: args{
				query: "SELECT * FROM Customers WHERE Country='Germany'; INSERT INTO Customers (CustomerName, ContactName, Address, City, PostalCode, Country) VALUES ('Cardinal; braun - stwet' 'stwert - linda; Tom B. Erichsen', 'Skagen;Tommy - 21', 'henry - lilone, Stavanger', '4006 - mikele', 'Norway - faraway');",
			},
			want: []Operation{
				{
					Name:  "SELECT",
					Table: "Customers",
				},
				{
					Name:  "INSERT",
					Table: "Customers",
				},
			},
		},
		{
			name: "test9",
			args: args{
				query: "Nosql: find key supplier.[name]",
			},
			want: []Operation{
				{
					Name:  "find",
					Table: "supplier.[name]",
				},
			},
		},
		{
			name: "test10",
			args: args{
				query: "NoSQL: delete from suppliers.name",
			},
			want: []Operation{
				{
					Name:  "delete",
					Table: "suppliers.name",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseStatement(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDatabaseOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
