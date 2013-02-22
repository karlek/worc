package save

import "os"
import "bytes"
import "encoding/gob"

import "github.com/karlek/worc/creature"
import "github.com/karlek/worc/terrain"
import "github.com/karlek/worc/menu"
import "github.com/mewkiz/pkg/errorsutil"

type Save struct {
	Path   string
	exists bool
}

func init() {
	gob.Register(new(creature.Creature))
	gob.Register(new(terrain.Area))
	gob.Register(new(menu.AreaScreen))
}

func New(path string) (sav *Save, err error) {
	sav = &Save{Path: path}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		sav.exists = false
	} else {
		sav.exists = true
	}

	return sav, nil
}

func (sav *Save) Save(blobs []interface{}) (err error) {
	var sto []interface{}
	for _, blob := range blobs {
		sto = append(sto, blob)
	}

	var buffer bytes.Buffer
	err = gob.NewEncoder(&buffer).Encode(&sto)
	if err != nil {
		return errorsutil.ErrorfColor("encode error: %#v", err)
	}

	f, err := os.Create(sav.Path)
	if err != nil {
		return err
	}
	_, err = f.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (sav *Save) Load() (blobs []interface{}, err error) {
	f, err := os.Open(sav.Path)
	if err != nil {
		return nil, err
	}

	sto := new([]interface{})
	err = gob.NewDecoder(f).Decode(sto)
	if err != nil {
		return nil, errorsutil.ErrorfColor("decode error: %#v", err)
	}

	return *sto, nil
}

// Returns true if a save file exists
func (sav *Save) Exists() bool {
	return sav.exists
}
