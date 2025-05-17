package service

import (
	"strconv"
	"sync"

	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/rpc/system/client/role"
	"github.com/bearllflee/scholar-track/rpc/system/internal/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

var CasbinServiceApp = new(CasbinService)

type CasbinService struct {
}

func (casbinService *CasbinService) UpdateApi(oldPath string, oldMethod string, newPath string, newMethod string) error {
	return global.DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(&gormadapter.CasbinRule{
		V1: newPath,
		V2: newMethod,
	}).Error
}

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

func (casbinService *CasbinService) AddPolicies(rules [][]string) error {
	e := casbinService.Casbin()
	_, err := e.AddPolicies(rules)
	return err
}

func (casbinService *CasbinService) UpdateCasbin(roleId uint64, policyInfos []*role.PolicyInfo) error {
	authorityId := strconv.Itoa(int(roleId))
	casbinService.ClearCasbin(0, authorityId)
	rules := make([][]string, 0)
	// 做权限去重
	duplicateMap := make(map[string]bool)
	for _, policyInfo := range policyInfos {
		key := authorityId + "_" + policyInfo.Path + "_" + policyInfo.Method
		if _, ok := duplicateMap[key]; !ok {
			duplicateMap[key] = true
			rules = append(rules, []string{authorityId, policyInfo.Path, policyInfo.Method})
		}
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return cerror.ErrHasSameApi
	}
	return nil
}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func (casbinService *CasbinService) Enforce(sub, obj, act string) (ok bool, err error) {
	e := casbinService.Casbin()
	ok, err = e.Enforce(sub, obj, act)
	return
}

func (casbinService *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		adapter, err := gormadapter.NewAdapterByDB(global.DB)
		if err != nil {
			logx.Error("casbin adapter error: ", err)
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub,p.sub) && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			logx.Error("casbin model error: ", err)
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, adapter)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		err = syncedCachedEnforcer.LoadPolicy()
		if err != nil {
			logx.Error("casbin load policy error: ", err)
			return
		}
	})
	return syncedCachedEnforcer
}

func (casbinService *CasbinService) FreshCasbin() error {
	e := casbinService.Casbin()
	err := e.LoadPolicy()
	return err
}
