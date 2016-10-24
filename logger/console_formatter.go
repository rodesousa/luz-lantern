// Copyright © 2016 Roberto De Sousa (https://github.com/rodesousa) / Patrick Tavares (https://github.com/ptavares)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ConsoleFormatter is the simpliest logger possible. Format the message as a fmt.Println()
package logger

import (
	"bytes"
	"fmt"
	"github.com/logrus"
)

type ConsoleFormatter struct{}

func (f ConsoleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	f.appendKeyValue(b, "time", entry.Time.Format(logrus.DefaultTimestampFormat))
	fmt.Fprintf(b, "[%s] ", entry.Level)

	if entry.Message != "" {
		fmt.Fprintf(b, "%s ", entry.Message)
	}
	for k := range entry.Data {
		f.appendKeyValue(b, k, entry.Data[k])
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *ConsoleFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	b.WriteString(key)
	b.WriteByte('=')

	switch value := value.(type) {
	case string:
		fmt.Fprintf(b, "%s", value)
	case error:
		errmsg := value.Error()
		fmt.Fprintf(b, "%s", errmsg)
	default:
		fmt.Fprint(b, value)
	}
	b.WriteByte(' ')
}
