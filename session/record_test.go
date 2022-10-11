package session

import "testing"

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func testRecordInt(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed insert test records")
	}
	return s
}

func testSession_Limit(t *testing.T){
	s := testRecordInt(t)
	var users []User
	err := s.Limit().Find(&users)
	if err != nil || len(users) != 1{
        t.Fatal("failed tp query with limit condition")
	}
}

func TestSession_Update(t *testing.T){
	s := testRecordInt(t)
	affected, _ :=s.Where("Name=?", "Tom").Update("Age", 30)
	u := &User()
	_ = .OrderBy("Age DESC").First(u)

	if affected != 1 || u.Age != 30{
		t.Fatal("failed to update")
	}
}