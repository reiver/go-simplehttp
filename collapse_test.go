package simplehttp


import (
	"testing"

	"reflect"
)


func TestCollapse(t *testing.T) {

	tests := []struct{
		Cascade     []interface{}
		Expected map[string]interface{}

	}{
		{
			Cascade:     []interface{}{},
			Expected: map[string]interface{}{},
		},



		{
			Cascade: []interface{}{
				map[string]string{
				},
			},
			Expected: map[string]interface{}{
			},
		},
		{
			Cascade: []interface{}{
				map[string]string{
				},
				map[string]string{
				},
			},
			Expected: map[string]interface{}{
			},
		},
		{
			Cascade: []interface{}{
				map[string]string{
				},
				map[string]string{
				},
				map[string]string{
				},
			},
			Expected: map[string]interface{}{
			},
		},



		{
			Cascade: []interface{}{
				struct{}{},
			},
			Expected: map[string]interface{}{
			},
		},
		{
			Cascade: []interface{}{
				struct{}{},
				struct{}{},
			},
			Expected: map[string]interface{}{
			},
		},
		{
			Cascade: []interface{}{
				struct{}{},
				struct{}{},
				struct{}{},
			},
			Expected: map[string]interface{}{
			},
		},



		{
			Cascade: []interface{}{
				map[string]string{
				},
				struct{}{},
			},
			Expected: map[string]interface{}{
			},
		},
		{
			Cascade: []interface{}{
				struct{}{},
				map[string]string{
				},
			},
			Expected: map[string]interface{}{
			},
		},



		{
			Cascade: []interface{}{
				map[string]string{
					"apple":"one",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
			},
		},
		{
			Cascade: []interface{}{
				map[string]string{
					"apple":"one",
				},
				map[string]string{
					"banana":"two",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
			},
		},
		{
			Cascade: []interface{}{
				map[string]string{
					"apple":"one",
				},
				map[string]string{
					"banana":"two",
				},
				map[string]string{
					"cherry":"three",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
			},
		},



		{
			Cascade: []interface{}{
			struct{
					Apple  string `json:"apple"`
				}{
					Apple:"one",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
			},
		},
		{
			Cascade: []interface{}{
				struct{
					Apple  string `json:"apple"`
				}{
					Apple:"one",
				},
				struct{
					Banana string `json:"banana"`
				}{
					Banana:"two",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
			},
		},
		{
			Cascade: []interface{}{
				struct{
					Apple  string `json:"apple"`
				}{
					Apple:"one",
				},
				struct{
					Banana string `json:"banana"`
				}{
					Banana:"two",
				},
				struct{
					Cherry string `json:"cherry"`
				}{
					Cherry:"three",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
			},
		},



		{
			Cascade: []interface{}{
				map[string]string{
					"apple":"one",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
			},
		},
		{
			Cascade: []interface{}{
				map[string]string{
					"apple":"one",
					"banana":"two",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
			},
		},
		{
			Cascade: []interface{}{
				map[string]string{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
			},
		},



		{
			Cascade: []interface{}{
				struct{
					Apple string `json:"apple"`
				}{
					Apple:"one",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
			},
		},
		{
			Cascade: []interface{}{
				struct{
					Apple  string `json:"apple"`
					Banana string `json:"banana"`
				}{
					Apple:"one",
					Banana:"two",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
			},
		},
		{
			Cascade: []interface{}{
				struct{
					Apple  string `json:"apple"`
					Banana string `json:"banana"`
					Cherry string `json:"cherry"`
				}{
					Apple:"one",
					Banana:"two",
					Cherry:"three",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
			},
		},



		{
			Cascade: []interface{}{
				struct{
					Apple  string `json:"apple"`
				}{
					Apple:"one",
				},
				map[string]interface{}{
					"banana":"two",
				},
				struct{
					Cherry string `json:"cherry"`
				}{
					Cherry:"three",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
			},
		},



		{
			Cascade: []interface{}{
				struct{
					Apple  string `json:"apple"`
				}{
					Apple:"one",
				},
				map[string]interface{}{
					"banana":"two",
				},
				struct{
					Cherry string `json:"cherry"`
				}{
					Cherry:"three",
				},
				map[string]interface{}{
					"fig":"left",
				},
				struct{
					Fig string `json:"fig"`
				}{
					Fig:"right",
				},
			},
			Expected: map[string]interface{}{
					"apple":"one",
					"banana":"two",
					"cherry":"three",
					"fig":"right",
			},
		},



		{
			Cascade: []interface{}{
				map[string]interface{}{
					"aardvark":"one",
					"bat":map[string]interface{}{
						"cat":3,
						"dog":"4",
					},
				},
			},
			Expected: map[string]interface{}{
					"aardvark":"one",
					"bat":map[string]interface{}{
						"cat":3,
						"dog":"4",
					},
			},
		},
		{
			Cascade: []interface{}{
				map[string]interface{}{
					"aardvark":"one",
					"bat":map[string]interface{}{
						"cat":3,
						"dog":"4",
					},
				},
				struct{
					Aardvark int `json:"aardvark"`
				}{
					Aardvark:123,
				},
			},
			Expected: map[string]interface{}{
					"aardvark":123,
					"bat":map[string]interface{}{
						"cat":3,
						"dog":"4",
					},
			},
		},
		{
			Cascade: []interface{}{
				map[string]interface{}{
					"aardvark":"one",
					"bat":map[string]interface{}{
						"cat":3,
						"dog":"4",
					},
				},
				struct{
					Aardvark int `json:"aardvark"`
				}{
					Aardvark:123,
				},
				&struct{
					Aardvark bool `json:"aardvark"`
				}{
					Aardvark:true,
				},
			},
			Expected: map[string]interface{}{
					"aardvark":true,
					"bat":map[string]interface{}{
						"cat":3,
						"dog":"4",
					},
			},
		},
	}


	for testNumber, test := range tests {
		collapsed := collapse(test.Cascade...)


		if expected, actual := test.Expected, collapsed; !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, expected %#v to be %#v, but wasn't.", testNumber, expected, actual)
		}
	}
}
