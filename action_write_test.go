package sql_test

import (
	"github.com/nelsonlai-go/sql"
	"github.com/nelsonlai-go/sql/conn"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("Sql Write", func() {

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

	It("should create", func() {
		t, err := sql.Create(db, &TestTable{Name: "k", Age: 11})
		Expect(err).To(BeNil())
		Expect(t.Name).To(Equal("k"))
		Expect(t.Age).To(Equal(11))

		t, err = sql.FindOne[TestTable](db, sql.Eq("name", "k"))
		Expect(err).To(BeNil())
		Expect(t.Name).To(Equal("k"))
		Expect(t.Age).To(Equal(11))
	})

	It("should update", func() {
		t, err := sql.FindOne[TestTable](db, sql.Eq("name", "a"))
		Expect(err).To(BeNil())
		Expect(t.Name).To(Equal("a"))
		Expect(t.Age).To(Equal(1))

		t.Name = "aa"
		t.Age = 11
		_, err = sql.Update(db, t)
		Expect(err).To(BeNil())

		t, err = sql.FindOne[TestTable](db, sql.Eq("name", "aa"))
		Expect(err).To(BeNil())
		Expect(t.Name).To(Equal("aa"))
		Expect(t.Age).To(Equal(11))
	})

	It("should delete", func() {
		t, err := sql.FindOne[TestTable](db, sql.Eq("name", "a"))
		Expect(err).To(BeNil())
		Expect(t.Name).To(Equal("a"))
		Expect(t.Age).To(Equal(1))

		err = sql.Delete(db, t)
		Expect(err).To(BeNil())

		t, err = sql.FindOne[TestTable](db, sql.Eq("name", "a"))
		Expect(err).NotTo(BeNil())
		Expect(t).To(BeNil())
	})

	It("should delete by", func() {
		t, err := sql.FindOne[TestTable](db, sql.Eq("name", "a"))
		Expect(err).To(BeNil())
		Expect(t.Name).To(Equal("a"))
		Expect(t.Age).To(Equal(1))

		err = sql.DeleteBy[TestTable](db, sql.Eq("name", "a"))
		Expect(err).To(BeNil())

		t, err = sql.FindOne[TestTable](db, sql.Eq("name", "a"))
		Expect(err).NotTo(BeNil())
		Expect(t).To(BeNil())
	})
})
