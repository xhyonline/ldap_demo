package ldap

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

const (
	username = "cn=Manager,dc=sso,dc=zime,dc=edu,dc=cn"
	password = "secret"
	baseDN   = "ou=people,dc=sso,dc=zime,dc=edu,dc=cn"
)

func main1() {
	conn, err := ldap.DialURL("ldap://192.168.1.240:389")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	err = conn.Bind(username, password)
	if err != nil {
		panic(err)
	}

	pagingControl := ldap.NewControlPaging(10000) // 分页遍历参数
	req := &ldap.SearchRequest{
		BaseDN:       baseDN,
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    999999, // 这个值随心调 ,默认是 0, 如果报错就调大,真不行就写 0 ,当记录数一共 1000条,但是这个值设置999就报错,设置1000就不会
		TimeLimit:    0,
		TypesOnly:    false,
		//Filter:       "(&(objectClass=users))",
		Filter: "(&(uid=16015))", // 这个条件本地不能执行,但是 linux 上可以
		//Filter:       "(&(uid=wifistaf))",	// 这个条件本地能执行
		Attributes: nil,
		Controls:   []ldap.Control{pagingControl},
	}
	for {
		result, err := conn.Search(req)
		if err != nil {
			panic(err)
		}
		// 遍历每一条记录
		for _, item := range result.Entries {
			for _, v := range item.Attributes {
				fmt.Println("属性名:", v.Name, "属性值:", v.Values[0])
			}
		}
		// 该作用不知道...... 只知道它是一个分页的
		updatedControl := ldap.FindControl(result.Controls, ldap.ControlTypePaging)
		if ctrl, ok := updatedControl.(*ldap.ControlPaging); ctrl != nil && ok && len(ctrl.Cookie) != 0 {
			pagingControl.SetCookie(ctrl.Cookie)
			continue
		} else {
			// 如果不能分页了就 break
			break
		}
	}
}

func main() {
	conn, err := ldap.DialURL("ldap://192.168.0.122:389")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	err = conn.Bind(username, password)
	if err != nil {
		panic(err)
	}

	pagingControl := ldap.NewControlPaging(10000) // 分页遍历参数
	req := &ldap.SearchRequest{
		BaseDN:       baseDN,
		Scope:        ldap.ScopeWholeSubtree,
		DerefAliases: ldap.NeverDerefAliases,
		SizeLimit:    999999, // 这个值随心调 ,默认是 0, 如果报错就调大,真不行就写 0 ,当记录数一共 1000条,但是这个值设置999就报错,设置1000就不会
		TimeLimit:    0,
		TypesOnly:    false,
		//Filter:       "(&(objectClass=users))",
		Filter: "(&(uid=TEACHER_li))", // 这个条件本地不能执行,但是 linux 上可以
		//Filter:       "(&(uid=wifistaf))",	// 这个条件本地能执行
		Attributes: nil,
		Controls:   []ldap.Control{pagingControl},
	}
	for {
		result, err := conn.Search(req)
		if err != nil {
			panic(err)
		}
		// 遍历每一条记录
		for _, item := range result.Entries {
			for _, v := range item.Attributes {
				fmt.Println("属性名:", v.Name, "属性值:", v.Values[0])
			}
		}
		// 该作用不知道...... 只知道它是一个分页的
		updatedControl := ldap.FindControl(result.Controls, ldap.ControlTypePaging)
		if ctrl, ok := updatedControl.(*ldap.ControlPaging); ctrl != nil && ok && len(ctrl.Cookie) != 0 {
			pagingControl.SetCookie(ctrl.Cookie)
			continue
		} else {
			// 如果不能分页了就 break
			break
		}
	}
}
