package sql_test

import (
	"github.com/nelsonlai-go/sql"
	"github.com/nelsonlai-go/sql/conn"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type TestTable struct {
	gorm.Model
	Name string
	Age  int
}

var _ = Describe("Sql Query", func() {

	var db *gorm.DB

	BeforeEach(func() {
		var err error
		db, err = conn.SqliteMemory(false, false)
		Expect(err).To(BeNil())

		err = db.AutoMigrate(&TestTable{})
		Expect(err).To(BeNil())

		names := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		ages := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i := 0; i < 10; i++ {
			db.Save(&TestTable{Name: names[i], Age: ages[i]})
		}
	})

	It("should find one - equal", func() {
		t, err := sql.FindOne[TestTable](db, sql.Eq("name", "a"))
		Expect(err).To(BeNil())
		Expect(t.Name).To(Equal("a"))
		Expect(t.Age).To(Equal(1))
	})

	It("should find all - not equal", func() {
		t, err := sql.FindAll[TestTable](db, sql.Neq("name", "a"), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(t)).To(Equal(9))
		for _, t := range t {
			Expect(t.Name).ToNot(Equal("a"))
		}
	})

	It("should find all - greater than", func() {
		ts, err := sql.FindAll[TestTable](db, sql.Gt("age", 5), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(5))
		for _, t := range ts {
			Expect(t.Age).To(BeNumerically(">", 5))
		}
	})

	It("should find all - greater than or equal", func() {
		ts, err := sql.FindAll[TestTable](db, sql.Gte("age", 5), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(6))
		for _, t := range ts {
			Expect(t.Age).To(BeNumerically(">=", 5))
		}
	})

	It("should find all - less than", func() {
		ts, err := sql.FindAll[TestTable](db, sql.Lt("age", 5), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(4))
		for _, t := range ts {
			Expect(t.Age).To(BeNumerically("<", 5))
		}
	})

	It("should find all - less than or equal", func() {
		ts, err := sql.FindAll[TestTable](db, sql.Lte("age", 5), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(5))
		for _, t := range ts {
			Expect(t.Age).To(BeNumerically("<=", 5))
		}
	})

	It("should find all - in", func() {
		ts, err := sql.FindAll[TestTable](db, sql.In("age", []uint{1, 2, 3}), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(3))
		for _, t := range ts {
			Expect(t.Age).To(BeElementOf([]int{1, 2, 3}))
		}
	})

	It("should find all - not in", func() {
		ts, err := sql.FindAll[TestTable](db, sql.Nin("age", []uint{1, 2, 3}), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(7))
		for _, t := range ts {
			Expect(t.Age).ToNot(BeElementOf([]int{1, 2, 3}))
		}
	})

	It("should find all - like", func() {
		ts, err := sql.FindAll[TestTable](db, sql.Lk("name", "%a%"), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(1))
		for _, t := range ts {
			Expect(t.Name).To(ContainSubstring("a"))
		}
	})

	It("should find all - not like", func() {
		ts, err := sql.FindAll[TestTable](db, sql.Nlk("name", "%a%"), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(9))
		for _, t := range ts {
			Expect(t.Name).ToNot(ContainSubstring("a"))
		}
	})

	It("should find all - between", func() {
		ts, err := sql.FindAll[TestTable](db, sql.Between("age", 5, 7), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(3))
		for _, t := range ts {
			Expect(t.Age).To(BeNumerically(">=", 5))
			Expect(t.Age).To(BeNumerically("<=", 7))
		}
	})

	It("should find all - order by", func() {
		ts, err := sql.FindAll[TestTable](db, nil, nil, &sql.Sort{By: "age", Asc: false})
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(10))
		Expect(ts[0].Age).To(Equal(10))
		Expect(ts[9].Age).To(Equal(1))
	})

	It("should find all - pagination", func() {
		ts, err := sql.FindAll[TestTable](db, nil, &sql.Pagination{Page: 1, Size: 3}, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(3))
		Expect(ts[0].Age).To(Equal(1))
		Expect(ts[1].Age).To(Equal(2))
		Expect(ts[2].Age).To(Equal(3))

		ts2, err := sql.FindAll[TestTable](db, nil, &sql.Pagination{Page: 2, Size: 3}, nil)
		Expect(err).To(BeNil())
		Expect(len(ts2)).To(Equal(3))
		Expect(ts2[0].Age).To(Equal(4))
		Expect(ts2[1].Age).To(Equal(5))
		Expect(ts2[2].Age).To(Equal(6))
	})

	It("should count", func() {
		c, err := sql.Count[TestTable](db, nil)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(int64(10)))

		c, err = sql.Count[TestTable](db, sql.Eq("age", 1))
		Expect(err).To(BeNil())
		Expect(c).To(Equal(int64(1)))
	})

	It("should query with pagination and sort", func() {
		ts, err := sql.FindAll[TestTable](db, nil, &sql.Pagination{Page: 2, Size: 3}, &sql.Sort{By: "name", Asc: true})
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(3))
		Expect(ts[0].Age).To(Equal(4))
		Expect(ts[1].Age).To(Equal(5))
		Expect(ts[2].Age).To(Equal(6))
	})

	It("should return nil if not found", func() {
		t, err := sql.FindOne[TestTable](db, sql.Eq("name", "z"))
		Expect(err).NotTo(BeNil())
		Expect(t).To(BeNil())
	})

	It("should return empty list if not found", func() {
		ts, err := sql.FindAll[TestTable](db, sql.Eq("name", "z"), nil, nil)
		Expect(err).To(BeNil())
		Expect(len(ts)).To(Equal(0))
	})
})
