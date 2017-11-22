package gaedb

import (
	"strconv"
	"bytes"
	"fmt"
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
	"github.com/strongo/log"
)

func key2str(key *datastore.Key) string {
	if key == nil {
		return "nil"
	}
	kind := key.Kind()
	if intID := key.IntID(); intID != 0 {
		return kind + ":int=" + strconv.FormatInt(intID, 10)
	} else if strID := key.StringID(); strID != "" {
		return kind + ":str=" + strID
	} else {
		return kind + ":new"
	}
}

func logKeys(c context.Context, f string, keys []*datastore.Key) {
	var buffer bytes.Buffer
	buffer.WriteString(f + "(\n")
	prevKey := "-"
	for _, key := range keys {
		ks := key2str(key)
		if ks == prevKey {
			log.Errorf(c, "Duplicate keys: "+ks)
		}
		buffer.WriteString(fmt.Sprintf("\t%v\n", ks))
		prevKey = ks
	}
	buffer.WriteString(")")
	log.Debugf(c, buffer.String())
}

