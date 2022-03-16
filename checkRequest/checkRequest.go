package main

import (
	"fmt"
)


func main() {
	reqPermission := []string{"PT", "CNX", "NEXT"}
	allowedPermission := []string{"CNX"}
	
	if err := CheckPermissions(reqPermission, allowedPermission); err != nil {
		fmt.Println(err)
	}
}

// CheckPermissions ...
func CheckPermissions(reqPermission []string, allowedPermission []string) error {
	var result []string
	
	if len(reqPermission) == 1 {
		for _, v := range allowedPermission {
			if v == reqPermission[0] {
				result = append(result, "match")
				return nil
			} 
		}
		result = append(result, "unmatch")
		return nil
	}

	if len(reqPermission) > 1 {
		if CheckPermissions(reqPermission[0:1], allowedPermission) == nil {
			return nil
		}
		CheckPermissions(reqPermission[1:], allowedPermission)
	}

	return nil
}
