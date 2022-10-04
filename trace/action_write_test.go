package trace_test

import (
	"github.com/go-ginseng/sql"
	"github.com/go-ginseng/sql/conn"
	"github.com/go-ginseng/sql/trace"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type TestModel struct {
	trace.Model
	Name string
}

type TestTraceModel struct {
	trace.Trace
	Name string
}

var _ = Describe("Trace Action", func() {

	var db *gorm.DB

	BeforeEach(func() {
		db, _ = conn.SqliteMemory(false, true)
		db.AutoMigrate(&TestModel{}, &TestTraceModel{})
	})

	It("should create a new record and a new trace record", func() {
		m := TestModel{Name: "test"}
		err := trace.Create(db, &m, &TestTraceModel{})
		Expect(err).To(BeNil())

		m2, err := sql.FindOne[TestModel](db, sql.Eq("id", m.ID))
		Expect(err).To(BeNil())

		Expect(m2).ToNot(BeNil())
		Expect(m2.Name).To(Equal("test"))
		Expect(m2.Version).To(Equal(uint(1)))
		Expect(m2.CreatedAt).ToNot(BeNil())
		Expect(m2.UpdatedAt).ToNot(BeNil())
		Expect(m2.DeletedAt.Valid).To(BeFalse())

		t, err := sql.FindOne[TestTraceModel](db, sql.Eq("record_id", m.ID))
		Expect(err).To(BeNil())

		Expect(t).ToNot(BeNil())
		Expect(t.Name).To(Equal("test"))
		Expect(t.Version).To(Equal(uint(1)))
		Expect(t.CreatedAt).ToNot(BeNil())
		Expect(t.UpdatedAt).ToNot(BeNil())
		Expect(t.DeletedAt.Valid).To(BeFalse())
		Expect(t.TraceTime).ToNot(BeNil())
		Expect(t.TraceAction).To(Equal("create"))
	})

	It("should update a record and create a new trace record", func() {
		m := TestModel{Name: "test"}
		err := trace.Create(db, &m, &TestTraceModel{})
		Expect(err).To(BeNil())

		m.Name = "test2"
		err = trace.Update(db, &m, &TestTraceModel{})
		Expect(err).To(BeNil())

		m2, err := sql.FindOne[TestModel](db, sql.Eq("id", m.ID))
		Expect(err).To(BeNil())

		Expect(m2).ToNot(BeNil())
		Expect(m2.Name).To(Equal("test2"))
		Expect(m2.Version).To(Equal(uint(2)))
		Expect(m2.CreatedAt).ToNot(BeNil())
		Expect(m2.UpdatedAt).ToNot(BeNil())
		Expect(m2.DeletedAt.Valid).To(BeFalse())

		t, err := sql.FindOne[TestTraceModel](db, sql.And(sql.Eq("record_id", m.ID), sql.Eq("trace_action", "update")))
		Expect(err).To(BeNil())

		Expect(t).ToNot(BeNil())
		Expect(t.Name).To(Equal("test2"))
		Expect(t.Version).To(Equal(uint(2)))
		Expect(t.CreatedAt).ToNot(BeNil())
		Expect(t.UpdatedAt).ToNot(BeNil())
		Expect(t.DeletedAt.Valid).To(BeFalse())
		Expect(t.TraceTime).ToNot(BeNil())
		Expect(t.TraceAction).To(Equal("update"))
	})

	It("should delete a record and create a new trace record", func() {
		m := TestModel{Name: "test"}
		err := trace.Create(db, &m, &TestTraceModel{})
		Expect(err).To(BeNil())

		err = trace.Delete(db, &m, &TestTraceModel{})
		Expect(err).To(BeNil())

		m2, err := sql.FindOne[TestModel](db, sql.Eq("id", m.ID))
		Expect(err).ToNot(BeNil())
		Expect(m2).To(BeNil())

		t, err := sql.UnscopedFindOne[TestTraceModel](db, sql.And(sql.Eq("record_id", m.ID), sql.Eq("trace_action", "delete")))
		Expect(err).To(BeNil())

		Expect(t).ToNot(BeNil())
		Expect(t.Name).To(Equal("test"))
		Expect(t.Version).To(Equal(uint(1)))
		Expect(t.CreatedAt).ToNot(BeNil())
		Expect(t.UpdatedAt).ToNot(BeNil())
		Expect(t.DeletedAt.Valid).To(BeTrue())
		Expect(t.TraceTime).ToNot(BeNil())
		Expect(t.TraceAction).To(Equal("delete"))
	})
})
