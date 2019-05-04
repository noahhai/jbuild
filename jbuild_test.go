package jbuild_test

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"

	"jbuild"
	"github.com/stretchr/testify/assert"
)

func TestAddMap(t *testing.T) {
	// arrange
	j := jbuild.Jmap{"j": "v2"}

	// act
	j.AddMap(jbuild.Jmap{"v2": "v5"}, "j")
	j.AddMap(jbuild.Jmap{"l": "v2"}, "k", "v1")
	j.AddMap(jbuild.Jmap{"m": "v3"}, "k", "v1")

	//assert
	expectedJson := `
{  
   "j":{  
      "v2":"v5"
   },
   "k":{  
      "v1":{  
         "l":"v2",
         "m":"v3"
      }
   }
}`

	m, err := json.Marshal(j)
	assert.Nil(t, err)
	assert.Equal(t, expectedJson, strings.TrimSpace(string(m)))
}

func TestAdd(t *testing.T) {
	// arrange
	j := jbuild.Jmap{"j": "v2"}

	// act
	j.Add("v2", "k", "v1", "l")
	j.Add("v3", "k", "v1", "m")

	//assert
	expectedJson := `
{  
   "j":"v2",
   "k":{  
      "v1":{  
         "l":"v2",
         "m":"v3"
      }
   }
}`

	m, err := json.Marshal(j)
	assert.Nil(t, err)
	assert.Equal(t, replaceAllNonWhitespace(expectedJson), replaceAllNonWhitespace(string(m)))
}

var mergeCases = []bool{
	false,
	true,
}

func TestMerge(t *testing.T) {
	for _, c := range mergeCases {
		t.Run(fmt.Sprintf("errorOnConflict=%t", c), func(t *testing.T) {
			// arrange
			j1 := jbuild.Jmap{"j": "v2"}
			j1.Add("v2", "k", "v1", "l")
			j1.Add("v3", "k", "v1", "m")
			j1.Add(jbuild.Jmap{"old":"val1"}, "k", "v1", "p")


			j2 := jbuild.Jmap{"j": "new-v2"}
			j2.Add("new-v2", "k", "v1", "l")
			j2.Add("v4", "k", "v1", "n")
			j2.Add(jbuild.Jmap{"new":"val2"}, "k", "v1", "p")

			errOnConflict := c

			// act
			err := j1.Merge(j2, &jbuild.MergeOptions{ErrorOnKeyConflict: errOnConflict})

			//assert
			assert.Equal(t, errOnConflict, err != nil)
			if errOnConflict {
				return
			}

			expectedJson := `
{  
   "j":"new-v2",
   "k":{  
      "v1":{  
         "l":"new-v2",
         "m":"v3",
         "n":"v4",
         "p":{
            "new":"val2",
            "old":"val1"
         }
      }
   }
}`

			m, err := json.Marshal(j1)
			assert.Nil(t, err)
			assert.Equal(t, replaceAllNonWhitespace(expectedJson), replaceAllNonWhitespace(string(m)))
		})
	}

}

func replaceAllNonWhitespace(s string) string {
	reg, err := regexp.Compile(`\s+`)
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, "")
}