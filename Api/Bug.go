package Api

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/google/go-querystring/query"
	"net/http"
)

type BugParam struct {
	Page   int    `url:"page"`
	Status string `url:"status"`
	Limit  int    `url:"limit"`
}

type BugSet struct {
	Page  int   `json:"page"`
	Total int   `json:"total"`
	Limit int   `json:"limit"`
	Bugs  []Bug `json:"bugs"`
}

type Bug struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Status        string `json:"status"`
	ResolvedBuild string `json:"resolvedBuild"`
	OpenedBuild   string `json:"openedBuild"`
}

func GetBugSet(productId int, param BugParam, bugSet *BugSet) error {
	baseUrl := fmt.Sprintf(URL+"api.php/v1/products/%d/bugs", productId)
	offsetUrl, err := query.Values(param)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("GET", baseUrl+"?"+offsetUrl.Encode(), nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return err
	}
	// 设置请求头
	req.Header.Set("Token", Cache.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return err
	}
	defer resp.Body.Close()

	err1 := UnmarshalResponse(resp, bugSet)
	if err1 != nil {
		return err1
	}
	return nil
}

func GetBugAmount(productId int, bugAmount *int) error {
	var bugSet BugSet
	if err := GetBugSet(productId, BugParam{
		Status: "all",
		Limit:  1,
	}, &bugSet); err != nil {
		return err
	}
	*bugAmount = bugSet.Total
	return nil
}

func GetAllBug(productId int, allBug *[]Bug) error {
	var bugAmount int
	if err := GetBugAmount(productId, &bugAmount); err != nil {
		return err
	}
	var bugSet BugSet
	if err := GetBugSet(productId, BugParam{Status: "all", Limit: bugAmount}, &bugSet); err != nil {
		return err
	}
	*allBug = bugSet.Bugs
	return nil
}

/*
获取所有解决版本
*/
func GetAllResolvedBuild(productId int, allResolvedBuild *[]string) error {
	var allBug []Bug
	if err := GetAllBug(productId, &allBug); err != nil {
		return nil
	}
	set := mapset.NewSet()
	for _, bug := range allBug {
		if bug.ResolvedBuild != "" {
			set.Add(bug.ResolvedBuild)
		}
	}
	set.Each(func(val interface{}) bool {
		*allResolvedBuild = append(*allResolvedBuild, val.(string))
		return false
	})
	return nil
}

func GetBugs(productId int, resolvedBuild string, bugs *[]Bug) error {
	var allBug []Bug
	if err := GetAllBug(productId, &allBug); err != nil {
		return fmt.Errorf("GetAllBug -> %v", err)
	}
	for _, bug := range allBug {
		if bug.ResolvedBuild == resolvedBuild {
			*bugs = append(*bugs, bug)
		}
	}
	return nil
}
