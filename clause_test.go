package sql_test

import (
	"github.com/go-ginseng/sql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Clause Builder", func() {

	It("should build - eq", func() {
		cls := sql.Eq("id", 1)
		stm, values := cls.Build()
		Expect(stm).To(Equal("id = ?"))
		Expect(values).To(Equal([]interface{}{1}))
	})

	It("should build - neq", func() {
		cls := sql.Neq("id", 1)
		stm, values := cls.Build()
		Expect(stm).To(Equal("id <> ?"))
		Expect(values).To(Equal([]interface{}{1}))
	})

	It("should build - gt", func() {
		cls := sql.Gt("id", 1)
		stm, values := cls.Build()
		Expect(stm).To(Equal("id > ?"))
		Expect(values).To(Equal([]interface{}{1}))
	})

	It("should build - gte", func() {
		cls := sql.Gte("id", 1)
		stm, values := cls.Build()
		Expect(stm).To(Equal("id >= ?"))
		Expect(values).To(Equal([]interface{}{1}))
	})

	It("should build - lt", func() {
		cls := sql.Lt("id", 1)
		stm, values := cls.Build()
		Expect(stm).To(Equal("id < ?"))
		Expect(values).To(Equal([]interface{}{1}))
	})

	It("should build - lte", func() {
		cls := sql.Lte("id", 1)
		stm, values := cls.Build()
		Expect(stm).To(Equal("id <= ?"))
		Expect(values).To(Equal([]interface{}{1}))
	})

	It("should build - like", func() {
		cls := sql.Lk("name", "test")
		stm, values := cls.Build()
		Expect(stm).To(Equal("name LIKE ?"))
		Expect(values).To(Equal([]interface{}{"test"}))
	})

	It("should build - not like", func() {
		cls := sql.Nlk("name", "test")
		stm, values := cls.Build()
		Expect(stm).To(Equal("name NOT LIKE ?"))
		Expect(values).To(Equal([]interface{}{"test"}))
	})

	It("should build - in", func() {
		cls := sql.In("id", []int{1, 2, 3})
		stm, values := cls.Build()
		Expect(stm).To(Equal("id IN ?"))
		Expect(len(values)).To(Equal(1))
		Expect(values[0]).To(Equal([]int{1, 2, 3}))
	})

	It("should build - not in", func() {
		cls := sql.Nin("id", []int{1, 2, 3})
		stm, values := cls.Build()
		Expect(stm).To(Equal("id NOT IN ?"))
		Expect(len(values)).To(Equal(1))
		Expect(values[0]).To(Equal([]int{1, 2, 3}))
	})

	It("should build - is null", func() {
		cls := sql.Null("id")
		stm, values := cls.Build()
		Expect(stm).To(Equal("id IS NULL"))
		Expect(len(values)).To(Equal(0))
	})

	It("should build - is not null", func() {
		cls := sql.NotNull("id")
		stm, values := cls.Build()
		Expect(stm).To(Equal("id IS NOT NULL"))
		Expect(len(values)).To(Equal(0))
	})

	It("should build - and", func() {
		cls := sql.And(sql.Eq("id", 1), sql.Eq("name", "test"))
		stm, values := cls.Build()
		Expect(stm).To(Equal("(id = ? AND name = ?)"))
		Expect(values).To(Equal([]interface{}{1, "test"}))
	})

	It("should build - or", func() {
		cls := sql.Or(sql.Eq("id", 1), sql.Eq("name", "test"))
		stm, values := cls.Build()
		Expect(stm).To(Equal("(id = ? OR name = ?)"))
		Expect(values).To(Equal([]interface{}{1, "test"}))
	})

	It("should build - between", func() {
		cls := sql.Between("id", 1, 10)
		stm, values := cls.Build()
		Expect(stm).To(Equal("(id >= ? AND id <= ?)"))
		Expect(values).To(Equal([]interface{}{1, 10}))
	})
})
