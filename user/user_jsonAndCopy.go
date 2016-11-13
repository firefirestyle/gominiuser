package user

import (
	"encoding/json"
	"time"

	"github.com/firefirestyle/go.miniprop"
	"golang.org/x/net/context"
)

// ----
// json and copy
// ----
func (userObj *User) SetUserFromsJson(ctx context.Context, source string) error {
	v := make(map[string]interface{})
	e := json.Unmarshal([]byte(source), &v)
	if e != nil {
		return e
	}
	//
	userObj.SetUserFromsMap(ctx, v)
	return nil
}

func (userObj *User) CopyWithoutUserNameAndSign(ctx context.Context, copyObj *User) {
	itemInfo := userObj.ToMapAll()
	itemInfo[TypeUserName] = copyObj.GetUserName()
	itemInfo[TypeSign] = copyObj.gaeObject.Sign
	copyObj.SetUserFromsMap(ctx, itemInfo)
}

func (userObj *User) SetUserFromsMap(ctx context.Context, v map[string]interface{}) {
	propObj := miniprop.NewMiniPropFromMap(v)
	userObj.gaeObject.RootGroup = propObj.GetString(TypeRootGroup, "")
	userObj.gaeObject.DisplayName = propObj.GetString(TypeDisplayName, "")
	userObj.gaeObject.UserName = propObj.GetString(TypeUserName, "")
	userObj.gaeObject.Created = propObj.GetTime(TypeCreated, time.Now()) //srcCreated
	userObj.gaeObject.Updated = propObj.GetTime(TypeUpdated, time.Now()) //time.Unix(0, int64(v[TypeLogined].(float64))) //srcLogin
	userObj.gaeObject.State = propObj.GetString(TypeState, "")
	userObj.gaeObject.PublicInfo = propObj.GetString(TypePublicInfo, "")
	userObj.gaeObject.PrivateInfo = propObj.GetString(TypePrivateInfo, "")
	userObj.gaeObject.PointValues = propObj.GetPropFloatList("", TypePointValues, []float64{0})
	userObj.gaeObject.PointNames = propObj.GetPropStringList("", TypePointNames, []string{"Main"})
	userObj.gaeObject.IconUrl = propObj.GetString(TypeIconUrl, "")
	userObj.gaeObject.Sign = propObj.GetString(TypeSign, "")
	userObj.gaeObject.Tags = propObj.GetString(TypeTags, "")
	userObj.gaeObject.Cont = propObj.GetString(TypeCont, "")
}

func (obj *User) ToMapPublic() map[string]interface{} {

	return map[string]interface{}{
		TypeRootGroup:   obj.gaeObject.RootGroup,
		TypeDisplayName: obj.gaeObject.DisplayName,        //
		TypeUserName:    obj.gaeObject.UserName,           //
		TypeCreated:     obj.gaeObject.Created.UnixNano(), //
		TypeUpdated:     obj.gaeObject.Updated.UnixNano(), //
		TypeState:       obj.gaeObject.State,              //
		TypePointNames:  obj.gaeObject.PointNames,         //
		TypePointValues: obj.gaeObject.PointValues,        //
		TypeTagNames:    obj.gaeObject.TagNames,           //
		TypeTagValues:   obj.gaeObject.TagValues,          //
		TypeIconUrl:     obj.gaeObject.IconUrl,            //
		TypePublicInfo:  obj.gaeObject.PublicInfo,
		TypeSign:        obj.gaeObject.Sign,
		TypeTags:        obj.gaeObject.Tags,
		TypeCont:        obj.gaeObject.Cont,
	}
}

func (obj *User) ToMapAll() map[string]interface{} {
	v := obj.ToMapPublic()
	v[TypePrivateInfo] = obj.gaeObject.PrivateInfo
	return v
}

func (obj *User) ToJson() []byte {
	return miniprop.NewMiniPropFromMap(obj.ToMapAll()).ToJson()
}

func (obj *User) ToJsonPublic() []byte {
	return miniprop.NewMiniPropFromMap(obj.ToMapPublic()).ToJson()
}
