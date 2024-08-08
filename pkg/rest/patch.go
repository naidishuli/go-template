package rest

import (
    "encoding/json"
    "reflect"

    "bets/utils"
    "gorm.io/gorm"
)

// PatchArgs will be used for optional keys in the request
type PatchArgs struct {
    Values map[string]any
    keys   []string
}

func (p *PatchArgs) AddKeys(keys ...string) {
    p.keys = append(p.keys, keys...)
}

func (p *PatchArgs) ExistValueKey(key string) bool {
    _, ok := p.Values[key]
    return ok
}

func (p *PatchArgs) ExistKey(key string) bool {
    for _, k := range p.keys {
        if key == k {
            return true
        }
    }
    return false
}

func (p *PatchArgs) GetKeys() []string {
    return p.keys
}

func (p *PatchArgs) GetValues() map[string]any {
    return p.Values
}

func (p *PatchArgs) SelectKeys(db *gorm.DB) *gorm.DB {
    if len(p.keys) == 0 {
        return db.Select("")
    }

    return db.Select(p.keys[0], p.keys[1:])
}

func (p *PatchArgs) Unmarshall(data []byte, dest any) error {
    err := json.Unmarshal(data, &p.Values)
    if err != nil {
        return err
    }

    p.assignKeys(dest, nil)

    return nil
}

func (p *PatchArgs) assignKeys(dest any, pTag *string) {
    v := reflect.ValueOf(dest)

    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }

    values := p.Values

    for i := 0; i < v.NumField(); i++ {
        field := v.Type().Field(i)
        if !field.IsExported() {
            continue
        }

        kTag := field.Tag.Get("key")
        jTag := field.Tag.Get("json")
        if jTag == "" {
            jTag = field.Tag.Get("form")
        }

        if kTag == "*" {
            kTag = jTag
        }

        if kTag == "-" || kTag == "" {
            continue
        }

        if kTag == "embedded" &&
            (field.Type.Kind() == reflect.Struct || field.Type.Elem().Kind() == reflect.Struct) {
            p.assignKeys(v.Field(i).Interface(), utils.PtrTo(jTag))
            continue
        }

        if pTag != nil {
            m, ok := values[*pTag]
            if ok {
                values = m.(map[string]any)
            }
        }

        _, ok := values[jTag]
        if ok {
            p.keys = append(p.keys, kTag)
        }
    }
}

type PatchManyAssociationsArgs struct {
    ID    uint
    Items []uint `json:"items"`
}
