package pingaccess

import (
	"encoding/json"
	"testing"

	pa "github.com/iwarapter/pingaccess-sdk-go/pingaccess"
)

func testHiddenFieldView() map[string]interface{} {
	return map[string]interface{}{
		"encrypted_value": "atat",
		"value":           "atat",
	}
}

func Test_weCanFlattenHiddenFieldView(t *testing.T) {
	initialHiddenFieldView := &pa.HiddenFieldView{
		Value:          String("atat"),
		EncryptedValue: String("atat"),
	}

	output := []map[string]interface{}{map[string]interface{}{ /*"encrypted_value": "atat" ,*/ "value": "atat"}}

	flattened := flattenHiddenFieldView(initialHiddenFieldView)

	equals(t, output, flattened)
}

func Test_expandHiddenFieldView(t *testing.T) {
	//expanded := flatmap.Expand(testHiddenFieldView(), "client_secret").([]interface{})
	expanded := []interface{}{testHiddenFieldView()}
	expandHiddenFieldView := expandHiddenFieldView(expanded)

	equals(t, "atat", *(*expandHiddenFieldView).Value)
	// equals(t, "atat", *(*expandHiddenFieldView).EncryptedValue)
}

func testOAuthClientCredentials() map[string]interface{} {
	return map[string]interface{}{
		"client_id": "atat",
		"client_secret": []interface{}{
			map[string]interface{}{
				"encrypted_value": "atat",
				"value":           "atat",
			},
		},
	}
}

func Test_weCanFlattenOAuthClientCredentials(t *testing.T) {
	initialOAuthClientCredentialsView := &pa.OAuthClientCredentialsView{
		ClientId: String("atat"),
		ClientSecret: &pa.HiddenFieldView{
			Value:          String("atat"),
			EncryptedValue: String("atat"),
		},
	}

	output := []map[string]interface{}{map[string]interface{}{"client_id": "atat", "client_secret": []map[string]interface{}{map[string]interface{}{"value": "atat" /*, "encrypted_value": "atat"*/}}}}

	flattened := flattenOAuthClientCredentialsView(initialOAuthClientCredentialsView)

	equals(t, output, flattened)
}

func Test_expandOAuthClientCredentials(t *testing.T) {
	//expanded := flatmap.Expand(testOAuthClientCredentials(), "client_credentials").([]interface{})
	expanded := []interface{}{testOAuthClientCredentials()}
	expandOAuthClientCredentialsView := expandOAuthClientCredentialsView(expanded)

	equals(t, "atat", *(*expandOAuthClientCredentialsView).ClientId)
	equals(t, "atat", *(*expandOAuthClientCredentialsView).ClientSecret.Value)
}

func testPolicyItem() map[string]interface{} {
	return map[string]interface{}{
		"id":   "1334",
		"type": "Rule",
	}
}

func Test_weCanFlattenPolicy(t *testing.T) {
	initialPolicyItem := []*pa.PolicyItem{
		&pa.PolicyItem{
			Id:   json.Number("1334"),
			Type: String("Rule"),
		},
	}

	output := []interface{}{map[string]interface{}{"id": "1334", "type": "Rule"}}

	flattened := flattenPolicyItem(initialPolicyItem)

	equals(t, output, flattened)
}

func Test_expandPolicyItem(t *testing.T) {
	//expanded := flatmap.Expand(testPolicyItem(), "policy.0.api").([]interface{})
	expanded := []interface{}{testPolicyItem()}
	expandPolicyItem := expandPolicyItem(expanded)

	equals(t, "1334", expandPolicyItem[0].Id.String())
	equals(t, "Rule", *(*expandPolicyItem[0]).Type)
}

func testPolicy() []interface{} {
	return []interface{}{map[string]interface{}{"api": []interface{}{map[string]interface{}{"id": "1334", "type": "Rule"}}}}
}

func Test_expandPolicy(t *testing.T) {
	// expanded := flatmap.Expand(testPolicy(), "policy").([]interface{})
	expandPolicyItem := expandPolicy(testPolicy())

	api := *(expandPolicyItem)["API"]
	equals(t, "1334", api[0].Id.String()) //[0].Id.String())
	// equals(t, "Rule", *(*expandPolicyItem)["api"])
}
