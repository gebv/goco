# AbstractModel

``` golang
type ModelAbstract struct {
    IsRemoved bool `json:"is_removed" sql:"type:boolean;default:false"`

    AtCreated time.Time   `json:"at_created" sql:"type:timestamp;default:null"`
    AtUpdated time.Time   `json:"at_updated" sql:"type:timestamp;default:null"`
    AtRemoved pq.NullTime `json:"at_removed,omitempty" sql:"type:timestamp;default:null"`
}

func (c *ModelAbstract) BeforeCreate() {
    c.AtCreated = time.Now()
}

func (c *ModelAbstract) AfterCreate() {

}

func (c *ModelAbstract) BeforeSave() {
    c.AtUpdated = time.Now()
}

func (c *ModelAbstract) BeforeDelete() {
    c.AtRemoved.Time = time.Now()
    c.AtRemoved.Valid = true
}

func (c *ModelAbstract) Maps() map[string]interface{} {
    return map[string]interface{}{
        "is_removed": &c.IsRemoved,

        "at_created": &c.AtCreated,
        "at_updated": &c.AtUpdated,
        "at_removed": &c.AtRemoved,
    }
}
```