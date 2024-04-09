

package favorite

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type Code4 int64

const (
	Code4_Success      Code4 = 1
	Code4_ParamInvalid Code4 = 2
	Code4_DBErr        Code4 = 3
)

func (p Code4) String() string {
	switch p {
	case Code4_Success:
		return "Success"
	case Code4_ParamInvalid:
		return "ParamInvalid"
	case Code4_DBErr:
		return "DBErr"
	}
	return "<UNSET>"
}

func Code4FromString(s string) (Code4, error) {
	switch s {
	case "Success":
		return Code4_Success, nil
	case "ParamInvalid":
		return Code4_ParamInvalid, nil
	case "DBErr":
		return Code4_DBErr, nil
	}
	return Code4(0), fmt.Errorf("not a valid Code4 string")
}

func Code4Ptr(v Code4) *Code4 { return &v }
func (p *Code4) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = Code4(result.Int64)
	return
}

func (p *Code4) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type User struct {
	UserID        int64  `thrift:"user_id,1" form:"user_id" json:"user_id" query:"user_id"`
	UserName      string `thrift:"user_name,2" form:"user_name" json:"user_name" query:"user_name"`
	FollowCount   int64  `thrift:"follow_count,3" form:"follow_count" json:"follow_count" query:"follow_count"`
	FollowerCount int64  `thrift:"follower_count,4" form:"follower_count" json:"follower_count" query:"follower_count"`
	FavoriteCount int64  `thrift:"favorite_count,5" form:"favorite_count" json:"favorite_count" query:"favorite_count"`
	Password      string `thrift:"password,6" form:"password" json:"password" query:"password"`
}

func NewUser() *User {
	return &User{}
}

func (p *User) GetUserID() (v int64) {
	return p.UserID
}

func (p *User) GetUserName() (v string) {
	return p.UserName
}

func (p *User) GetFollowCount() (v int64) {
	return p.FollowCount
}

func (p *User) GetFollowerCount() (v int64) {
	return p.FollowerCount
}

func (p *User) GetFavoriteCount() (v int64) {
	return p.FavoriteCount
}

func (p *User) GetPassword() (v string) {
	return p.Password
}

var fieldIDToName_User = map[int16]string{
	1: "user_id",
	2: "user_name",
	3: "follow_count",
	4: "follower_count",
	5: "favorite_count",
	6: "password",
}

func (p *User) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField6(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_User[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *User) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.UserID = v
	}
	return nil
}
func (p *User) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.UserName = v
	}
	return nil
}
func (p *User) ReadField3(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.FollowCount = v
	}
	return nil
}
func (p *User) ReadField4(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.FollowerCount = v
	}
	return nil
}
func (p *User) ReadField5(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.FavoriteCount = v
	}
	return nil
}
func (p *User) ReadField6(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Password = v
	}
	return nil
}

func (p *User) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("User"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
		if err = p.writeField6(oprot); err != nil {
			fieldId = 6
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *User) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user_id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.UserID); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *User) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user_name", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UserName); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *User) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("follow_count", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.FollowCount); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *User) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("follower_count", thrift.I64, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.FollowerCount); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *User) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("favorite_count", thrift.I64, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.FavoriteCount); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *User) writeField6(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("password", thrift.STRING, 6); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Password); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 end error: ", p), err)
}

func (p *User) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("User(%+v)", *p)

}

type Video struct {
	VideoId       int64  `thrift:"VideoId,1" form:"VideoId" json:"VideoId" query:"VideoId"`
	AuthorId      int64  `thrift:"AuthorId,2" form:"AuthorId" json:"AuthorId" query:"AuthorId"`
	PlayUrl       string `thrift:"PlayUrl,3" form:"PlayUrl" json:"PlayUrl" query:"PlayUrl"`
	CoverUrl      string `thrift:"CoverUrl,4" form:"CoverUrl" json:"CoverUrl" query:"CoverUrl"`
	FavoriteCount int64  `thrift:"FavoriteCount,5" form:"FavoriteCount" json:"FavoriteCount" query:"FavoriteCount"`
	CommentCount  int64  `thrift:"CommentCount,6" form:"CommentCount" json:"CommentCount" query:"CommentCount"`
	PublishTime   string `thrift:"PublishTime,7" form:"PublishTime" json:"PublishTime" query:"PublishTime"`
	Title         string `thrift:"Title,8" form:"Title" json:"Title" query:"Title"`
}

func NewVideo() *Video {
	return &Video{}
}

func (p *Video) GetVideoId() (v int64) {
	return p.VideoId
}

func (p *Video) GetAuthorId() (v int64) {
	return p.AuthorId
}

func (p *Video) GetPlayUrl() (v string) {
	return p.PlayUrl
}

func (p *Video) GetCoverUrl() (v string) {
	return p.CoverUrl
}

func (p *Video) GetFavoriteCount() (v int64) {
	return p.FavoriteCount
}

func (p *Video) GetCommentCount() (v int64) {
	return p.CommentCount
}

func (p *Video) GetPublishTime() (v string) {
	return p.PublishTime
}

func (p *Video) GetTitle() (v string) {
	return p.Title
}

var fieldIDToName_Video = map[int16]string{
	1: "VideoId",
	2: "AuthorId",
	3: "PlayUrl",
	4: "CoverUrl",
	5: "FavoriteCount",
	6: "CommentCount",
	7: "PublishTime",
	8: "Title",
}

func (p *Video) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField6(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 7:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField7(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 8:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField8(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Video[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Video) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.VideoId = v
	}
	return nil
}
func (p *Video) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.AuthorId = v
	}
	return nil
}
func (p *Video) ReadField3(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PlayUrl = v
	}
	return nil
}
func (p *Video) ReadField4(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.CoverUrl = v
	}
	return nil
}
func (p *Video) ReadField5(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.FavoriteCount = v
	}
	return nil
}
func (p *Video) ReadField6(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.CommentCount = v
	}
	return nil
}
func (p *Video) ReadField7(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.PublishTime = v
	}
	return nil
}
func (p *Video) ReadField8(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Title = v
	}
	return nil
}

func (p *Video) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Video"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
		if err = p.writeField6(oprot); err != nil {
			fieldId = 6
			goto WriteFieldError
		}
		if err = p.writeField7(oprot); err != nil {
			fieldId = 7
			goto WriteFieldError
		}
		if err = p.writeField8(oprot); err != nil {
			fieldId = 8
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Video) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("VideoId", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.VideoId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Video) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("AuthorId", thrift.I64, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.AuthorId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Video) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("PlayUrl", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PlayUrl); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Video) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("CoverUrl", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.CoverUrl); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *Video) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("FavoriteCount", thrift.I64, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.FavoriteCount); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *Video) writeField6(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("CommentCount", thrift.I64, 6); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.CommentCount); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 end error: ", p), err)
}

func (p *Video) writeField7(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("PublishTime", thrift.STRING, 7); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.PublishTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 7 end error: ", p), err)
}

func (p *Video) writeField8(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Title", thrift.STRING, 8); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Title); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 8 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 8 end error: ", p), err)
}

func (p *Video) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Video(%+v)", *p)

}

type Favorite struct {
	FavoriteId int64 `thrift:"FavoriteId,1" form:"FavoriteId" json:"FavoriteId" query:"FavoriteId"`
	UserId     int64 `thrift:"UserId,2" form:"UserId" json:"UserId" query:"UserId"`
	VideoId    int64 `thrift:"VideoId,3" form:"VideoId" json:"VideoId" query:"VideoId"`
	CommentId  int64 `thrift:"CommentId,4" form:"CommentId" json:"CommentId" query:"CommentId"`
}

func NewFavorite() *Favorite {
	return &Favorite{}
}

func (p *Favorite) GetFavoriteId() (v int64) {
	return p.FavoriteId
}

func (p *Favorite) GetUserId() (v int64) {
	return p.UserId
}

func (p *Favorite) GetVideoId() (v int64) {
	return p.VideoId
}

func (p *Favorite) GetCommentId() (v int64) {
	return p.CommentId
}

var fieldIDToName_Favorite = map[int16]string{
	1: "FavoriteId",
	2: "UserId",
	3: "VideoId",
	4: "CommentId",
}

func (p *Favorite) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Favorite[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Favorite) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.FavoriteId = v
	}
	return nil
}
func (p *Favorite) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.UserId = v
	}
	return nil
}
func (p *Favorite) ReadField3(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.VideoId = v
	}
	return nil
}
func (p *Favorite) ReadField4(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.CommentId = v
	}
	return nil
}

func (p *Favorite) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Favorite"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Favorite) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("FavoriteId", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.FavoriteId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *Favorite) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("UserId", thrift.I64, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.UserId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *Favorite) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("VideoId", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.VideoId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *Favorite) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("CommentId", thrift.I64, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.CommentId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *Favorite) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Favorite(%+v)", *p)

}

type FavoriteRequest struct {

	ActionType int64 `thrift:"ActionType,1" form:"action_type" form:"action_type" json:"action_type" vd:"($==0 ||$==1)"`
	CommentId  int64 `thrift:"CommentId,2" form:"comment_id" json:"comment_id"`
	VideoType  int64 `thrift:"VideoType,3" form:"video_type" form:"video_type" json:"video_type"`

}

func NewFavoriteRequest() *FavoriteRequest {
	return &FavoriteRequest{}
}

func (p *FavoriteRequest) GetActionType() (v int64) {
	return p.ActionType
}

func (p *FavoriteRequest) GetCommentId() (v int64) {
	return p.CommentId
}


func (p *FavoriteRequest) GetVideoType() (v int64) {
	return p.VideoType
}

var fieldIDToName_FavoriteRequest = map[int16]string{
	1: "ActionType",
	2: "CommentId",
	3: "VideoType",

}

func (p *FavoriteRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}

		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}

		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FavoriteRequest) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.ActionType = v
	}
	return nil
}
func (p *FavoriteRequest) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.CommentId = v
	}
	return nil
}

func (p *FavoriteRequest) ReadField3(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.VideoType = v
	}
	return nil
}


func (p *FavoriteRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("FavoriteRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FavoriteRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("ActionType", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.ActionType); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *FavoriteRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("CommentId", thrift.I64, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.CommentId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}


func (p *FavoriteRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("VideoType", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.VideoType); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}


func (p *FavoriteRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FavoriteRequest(%+v)", *p)

}

type FavoriteResponse struct {
	Code Code4  `thrift:"code,1" form:"code" json:"code" query:"code"`
	Msg  string `thrift:"msg,2" form:"msg" json:"msg" query:"msg"`
}

func NewFavoriteResponse() *FavoriteResponse {
	return &FavoriteResponse{}
}

func (p *FavoriteResponse) GetCode() (v Code4) {
	return p.Code
}

func (p *FavoriteResponse) GetMsg() (v string) {
	return p.Msg
}

var fieldIDToName_FavoriteResponse = map[int16]string{
	1: "code",
	2: "msg",
}

func (p *FavoriteResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FavoriteResponse) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Code = Code4(v)
	}
	return nil
}
func (p *FavoriteResponse) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Msg = v
	}
	return nil
}

func (p *FavoriteResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("FavoriteResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FavoriteResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("code", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(int32(p.Code)); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *FavoriteResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Msg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *FavoriteResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FavoriteResponse(%+v)", *p)

}

type ListFavoriteRequest struct {
}

func NewListFavoriteRequest() *ListFavoriteRequest {
	return &ListFavoriteRequest{}
}

var fieldIDToName_ListFavoriteRequest = map[int16]string{}

func (p *ListFavoriteRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ListFavoriteRequest) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("ListFavoriteRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ListFavoriteRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ListFavoriteRequest(%+v)", *p)

}

type ListFavoriteResponse struct {
	Code  Code4       `thrift:"code,1" form:"code" json:"code" query:"code"`
	Msg   string      `thrift:"msg,2" form:"msg" json:"msg" query:"msg"`
	Favs  []*Favorite `thrift:"favs,3" form:"favs" json:"favs" query:"favs"`
	Users []*User     `thrift:"users,4" form:"users" json:"users" query:"users"`
}

func NewListFavoriteResponse() *ListFavoriteResponse {
	return &ListFavoriteResponse{}
}

func (p *ListFavoriteResponse) GetCode() (v Code4) {
	return p.Code
}

func (p *ListFavoriteResponse) GetMsg() (v string) {
	return p.Msg
}

func (p *ListFavoriteResponse) GetFavs() (v []*Favorite) {
	return p.Favs
}

func (p *ListFavoriteResponse) GetUsers() (v []*User) {
	return p.Users
}

var fieldIDToName_ListFavoriteResponse = map[int16]string{
	1: "code",
	2: "msg",
	3: "favs",
	4: "users",
}

func (p *ListFavoriteResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ListFavoriteResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ListFavoriteResponse) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Code = Code4(v)
	}
	return nil
}
func (p *ListFavoriteResponse) ReadField2(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Msg = v
	}
	return nil
}
func (p *ListFavoriteResponse) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Favs = make([]*Favorite, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewFavorite()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.Favs = append(p.Favs, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}
func (p *ListFavoriteResponse) ReadField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Users = make([]*User, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewUser()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.Users = append(p.Users, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *ListFavoriteResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ListFavoriteResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ListFavoriteResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("code", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(int32(p.Code)); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ListFavoriteResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Msg); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *ListFavoriteResponse) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("favs", thrift.LIST, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Favs)); err != nil {
		return err
	}
	for _, v := range p.Favs {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *ListFavoriteResponse) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("users", thrift.LIST, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Users)); err != nil {
		return err
	}
	for _, v := range p.Users {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *ListFavoriteResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ListFavoriteResponse(%+v)", *p)

}

type FavoriteService interface {
	FavoriteService(ctx context.Context, req *FavoriteRequest) (r *FavoriteResponse, err error)

	ListFavorite(ctx context.Context, req *ListFavoriteRequest) (r *ListFavoriteResponse, err error)
}

type FavoriteServiceClient struct {
	c thrift.TClient
}

func NewFavoriteServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FavoriteServiceClient {
	return &FavoriteServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewFavoriteServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FavoriteServiceClient {
	return &FavoriteServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewFavoriteServiceClient(c thrift.TClient) *FavoriteServiceClient {
	return &FavoriteServiceClient{
		c: c,
	}
}

func (p *FavoriteServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *FavoriteServiceClient) FavoriteService(ctx context.Context, req *FavoriteRequest) (r *FavoriteResponse, err error) {
	var _args FavoriteServiceFavoriteServiceArgs
	_args.Req = req
	var _result FavoriteServiceFavoriteServiceResult
	if err = p.Client_().Call(ctx, "FavoriteService", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
func (p *FavoriteServiceClient) ListFavorite(ctx context.Context, req *ListFavoriteRequest) (r *ListFavoriteResponse, err error) {
	var _args FavoriteServiceListFavoriteArgs
	_args.Req = req
	var _result FavoriteServiceListFavoriteResult
	if err = p.Client_().Call(ctx, "ListFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type FavoriteServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      FavoriteService
}

func (p *FavoriteServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *FavoriteServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *FavoriteServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewFavoriteServiceProcessor(handler FavoriteService) *FavoriteServiceProcessor {
	self := &FavoriteServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("FavoriteService", &favoriteServiceProcessorFavoriteService{handler: handler})
	self.AddToProcessorMap("ListFavorite", &favoriteServiceProcessorListFavorite{handler: handler})
	return self
}
func (p *FavoriteServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type favoriteServiceProcessorFavoriteService struct {
	handler FavoriteService
}

func (p *favoriteServiceProcessorFavoriteService) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FavoriteServiceFavoriteServiceArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("FavoriteService", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := FavoriteServiceFavoriteServiceResult{}
	var retval *FavoriteResponse
	if retval, err2 = p.handler.FavoriteService(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing FavoriteService: "+err2.Error())
		oprot.WriteMessageBegin("FavoriteService", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("FavoriteService", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type favoriteServiceProcessorListFavorite struct {
	handler FavoriteService
}

func (p *favoriteServiceProcessorListFavorite) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FavoriteServiceListFavoriteArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ListFavorite", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := FavoriteServiceListFavoriteResult{}
	var retval *ListFavoriteResponse
	if retval, err2 = p.handler.ListFavorite(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ListFavorite: "+err2.Error())
		oprot.WriteMessageBegin("ListFavorite", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("ListFavorite", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type FavoriteServiceFavoriteServiceArgs struct {
	Req *FavoriteRequest `thrift:"req,1"`
}

func NewFavoriteServiceFavoriteServiceArgs() *FavoriteServiceFavoriteServiceArgs {
	return &FavoriteServiceFavoriteServiceArgs{}
}

var FavoriteServiceFavoriteServiceArgs_Req_DEFAULT *FavoriteRequest

func (p *FavoriteServiceFavoriteServiceArgs) GetReq() (v *FavoriteRequest) {
	if !p.IsSetReq() {
		return FavoriteServiceFavoriteServiceArgs_Req_DEFAULT
	}
	return p.Req
}

var fieldIDToName_FavoriteServiceFavoriteServiceArgs = map[int16]string{
	1: "req",
}

func (p *FavoriteServiceFavoriteServiceArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FavoriteServiceFavoriteServiceArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteServiceFavoriteServiceArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FavoriteServiceFavoriteServiceArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewFavoriteRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FavoriteServiceFavoriteServiceArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("FavoriteService_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FavoriteServiceFavoriteServiceArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *FavoriteServiceFavoriteServiceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FavoriteServiceFavoriteServiceArgs(%+v)", *p)

}

type FavoriteServiceFavoriteServiceResult struct {
	Success *FavoriteResponse `thrift:"success,0,optional"`
}

func NewFavoriteServiceFavoriteServiceResult() *FavoriteServiceFavoriteServiceResult {
	return &FavoriteServiceFavoriteServiceResult{}
}

var FavoriteServiceFavoriteServiceResult_Success_DEFAULT *FavoriteResponse

func (p *FavoriteServiceFavoriteServiceResult) GetSuccess() (v *FavoriteResponse) {
	if !p.IsSetSuccess() {
		return FavoriteServiceFavoriteServiceResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_FavoriteServiceFavoriteServiceResult = map[int16]string{
	0: "success",
}

func (p *FavoriteServiceFavoriteServiceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteServiceFavoriteServiceResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteServiceFavoriteServiceResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FavoriteServiceFavoriteServiceResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewFavoriteResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FavoriteServiceFavoriteServiceResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("FavoriteService_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FavoriteServiceFavoriteServiceResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *FavoriteServiceFavoriteServiceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FavoriteServiceFavoriteServiceResult(%+v)", *p)

}

type FavoriteServiceListFavoriteArgs struct {
	Req *ListFavoriteRequest `thrift:"req,1"`
}

func NewFavoriteServiceListFavoriteArgs() *FavoriteServiceListFavoriteArgs {
	return &FavoriteServiceListFavoriteArgs{}
}

var FavoriteServiceListFavoriteArgs_Req_DEFAULT *ListFavoriteRequest

func (p *FavoriteServiceListFavoriteArgs) GetReq() (v *ListFavoriteRequest) {
	if !p.IsSetReq() {
		return FavoriteServiceListFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

var fieldIDToName_FavoriteServiceListFavoriteArgs = map[int16]string{
	1: "req",
}

func (p *FavoriteServiceListFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FavoriteServiceListFavoriteArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteServiceListFavoriteArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FavoriteServiceListFavoriteArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewListFavoriteRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FavoriteServiceListFavoriteArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ListFavorite_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FavoriteServiceListFavoriteArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *FavoriteServiceListFavoriteArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FavoriteServiceListFavoriteArgs(%+v)", *p)

}

type FavoriteServiceListFavoriteResult struct {
	Success *ListFavoriteResponse `thrift:"success,0,optional"`
}

func NewFavoriteServiceListFavoriteResult() *FavoriteServiceListFavoriteResult {
	return &FavoriteServiceListFavoriteResult{}
}

var FavoriteServiceListFavoriteResult_Success_DEFAULT *ListFavoriteResponse

func (p *FavoriteServiceListFavoriteResult) GetSuccess() (v *ListFavoriteResponse) {
	if !p.IsSetSuccess() {
		return FavoriteServiceListFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_FavoriteServiceListFavoriteResult = map[int16]string{
	0: "success",
}

func (p *FavoriteServiceListFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteServiceListFavoriteResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_FavoriteServiceListFavoriteResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *FavoriteServiceListFavoriteResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewListFavoriteResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *FavoriteServiceListFavoriteResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ListFavorite_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *FavoriteServiceListFavoriteResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *FavoriteServiceListFavoriteResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FavoriteServiceListFavoriteResult(%+v)", *p)

}
