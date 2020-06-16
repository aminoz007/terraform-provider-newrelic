// Code generated by typegen; DO NOT EDIT.

package alerts

/* AlertsIncidentPreference - Determines how incidents are created for critical violations of the conditions contained in the policy. */
type AlertsIncidentPreference string

var AlertsIncidentPreferenceTypes = struct {
	/* PER_CONDITION - A condition will create a condition-level incident when it violates its critical threshold.
	Other violating conditions will create their own incidents. */
	PER_CONDITION AlertsIncidentPreference
	/* PER_CONDITION_AND_TARGET - Each target of each condition will create an entity-level incident upon critical violation.
	Other violating targets will create their own incidents (even on the same condition). */
	PER_CONDITION_AND_TARGET AlertsIncidentPreference
	/* PER_POLICY - A condition will create a policy-level incident when it violates its critical threshold.
	Other violating conditions will be grouped into this incident. */
	PER_POLICY AlertsIncidentPreference
}{
	PER_CONDITION:            "PER_CONDITION",
	PER_CONDITION_AND_TARGET: "PER_CONDITION_AND_TARGET",
	PER_POLICY:               "PER_POLICY",
}

/* AlertsMutingRuleConditionGroupInput - A group of MutingRuleConditions combined by an operator. */
type AlertsMutingRuleConditionGroupInput struct {
	/* Conditions - The individual MutingRuleConditions within the group. */
	Conditions []AlertsMutingRuleConditionInput `json:"conditions"`
	/* Operator - The operator used to combine all the MutingRuleConditions within the group. */
	Operator AlertsMutingRuleConditionGroupOperator `json:"operator"`
}

/* AlertsMutingRuleConditionGroupOperator - An operator used to combine MutingRuleConditions within a MutingRuleConditionGroup. */
type AlertsMutingRuleConditionGroupOperator string

var AlertsMutingRuleConditionGroupOperatorTypes = struct {
	/* AND - Match conditions by AND */
	AND AlertsMutingRuleConditionGroupOperator
	/* OR - Match conditions by OR */
	OR AlertsMutingRuleConditionGroupOperator
}{
	AND: "AND",
	OR:  "OR",
}

/* AlertsMutingRuleConditionInput - A condition which describes how to target a New Relic Alerts Violation. */
type AlertsMutingRuleConditionInput struct {
	/* Attribute - The attribute on a violation. Expects one of:

	* **accountId** - The account id
	* **conditionId** - The alert condition id
	* **policyId** - The alert policy id
	* **policyName** - The alert policy name
	* **conditionName** - The alert condition name
	* **conditionType** - The alert condition type, such as `metric`
	* **conditionRunbookUrl** - The alert condition's runbook url
	* **product** - The target product (e.g., `SYNTHETICS`)
	* **targetId** - The ID of the alerts target
	* **targetName** - The name of the alerts target
	* **nrqlEventType** - The NRQL event type
	* **tag** - Arbitrary tags associated with some entity (e.g., FACET from a NRQL query)
	* **nrqlQuery** - The NRQL query string */
	Attribute string `json:"attribute"`
	/* Operator - The operator used to compare the attribute's value with the supplied value(s). */
	Operator AlertsMutingRuleConditionOperator `json:"operator"`
	/* Values - The value(s) to compare against the attribute's value. */
	Values []string `json:"values"`
}

/* AlertsMutingRuleConditionOperator - The list of operators to be used in a MutingRuleCondition. Each operator is limited to one value in the `values` list unless otherwise specified. */
type AlertsMutingRuleConditionOperator string

var AlertsMutingRuleConditionOperatorTypes = struct {
	/* ANY - Where attribute is any. */
	ANY AlertsMutingRuleConditionOperator
	/* CONTAINS - Where attribute contains value. */
	CONTAINS AlertsMutingRuleConditionOperator
	/* ENDS_WITH - Where attribute ends with value. */
	ENDS_WITH AlertsMutingRuleConditionOperator
	/* EQUALS - Where attribute equals value. */
	EQUALS AlertsMutingRuleConditionOperator
	/* IN - Where attribute in values. (Limit 500) */
	IN AlertsMutingRuleConditionOperator
	/* IS_BLANK - Where attribute is blank. */
	IS_BLANK AlertsMutingRuleConditionOperator
	/* IS_NOT_BLANK - Where attribute is not blank. */
	IS_NOT_BLANK AlertsMutingRuleConditionOperator
	/* NOT_CONTAINS - Where attribute does not contain value. */
	NOT_CONTAINS AlertsMutingRuleConditionOperator
	/* NOT_ENDS_WITH - Where attribute does not end with value. */
	NOT_ENDS_WITH AlertsMutingRuleConditionOperator
	/* NOT_EQUALS - Where attribute does not equal value. */
	NOT_EQUALS AlertsMutingRuleConditionOperator
	/* NOT_IN - Where attribute not in values. (Limit 500) */
	NOT_IN AlertsMutingRuleConditionOperator
	/* NOT_STARTS_WITH - Where attribute does not start with value. */
	NOT_STARTS_WITH AlertsMutingRuleConditionOperator
	/* STARTS_WITH - Where attribute starts with value. */
	STARTS_WITH AlertsMutingRuleConditionOperator
}{
	ANY:             "ANY",
	CONTAINS:        "CONTAINS",
	ENDS_WITH:       "ENDS_WITH",
	EQUALS:          "EQUALS",
	IN:              "IN",
	IS_BLANK:        "IS_BLANK",
	IS_NOT_BLANK:    "IS_NOT_BLANK",
	NOT_CONTAINS:    "NOT_CONTAINS",
	NOT_ENDS_WITH:   "NOT_ENDS_WITH",
	NOT_EQUALS:      "NOT_EQUALS",
	NOT_IN:          "NOT_IN",
	NOT_STARTS_WITH: "NOT_STARTS_WITH",
	STARTS_WITH:     "STARTS_WITH",
}

/* AlertsMutingRuleInput - Input for creating MutingRules for New Relic Alerts Violations. */
type AlertsMutingRuleInput struct {
	/* Condition - The condition that defines which violations to target. */
	Condition AlertsMutingRuleConditionGroupInput `json:"condition"`
	/* Description - The description of the MutingRule. */
	Description string `json:"description"`
	/* Enabled - Whether the MutingRule is enabled */
	Enabled bool `json:"enabled"`
	/* Name - The name of the MutingRule. */
	Name string `json:"name"`
}

/* AlertsPoliciesSearchCriteriaInput - Search criteria for returning specific policies. */
type AlertsPoliciesSearchCriteriaInput struct {
	/* IDs - The list of policy ids to return. */
	IDs []string `json:"ids"`
}

/* AlertsPoliciesSearchResultSet - Collection of policies with pagination information. */
type AlertsPoliciesSearchResultSet struct {
	/* NextCursor - Cursor pointing to the end of the current page of policy records. Null if final page. */
	NextCursor string `json:"nextCursor"`
	/* Policies - Set of policies returned for the supplied cursor and criteria. */
	Policies []AlertsPolicy `json:"policies"`
	/* TotalCount - Total number of policy records for the given search criteria. */
	TotalCount int `json:"totalCount"`
}

/* AlertsPolicy - Container for conditions with associated notifications channels. */
type AlertsPolicy struct {
	/* AccountID - Account ID of the policy. */
	AccountID int `json:"accountId"`
	/* ID - Primary key for policies. */
	ID string `json:"id"`
	/* IncidentPreference - Determines how incidents are created for critical violations of the conditions contained in the policy. */
	IncidentPreference AlertsIncidentPreference `json:"incidentPreference"`
	/* Name - Description of the policy. */
	Name string `json:"name"`
}

/* AlertsPolicyInput - Container for conditions with associated notifications channels. */
type AlertsPolicyInput struct {
	/* IncidentPreference - Determines how incidents are created for critical violations of the conditions contained in the policy. */
	IncidentPreference AlertsIncidentPreference `json:"incidentPreference"`
	/* Name - Description of the policy. */
	Name string `json:"name"`
}

/* AlertsPolicyUpdateInput - Policy fields to be updated. */
type AlertsPolicyUpdateInput struct {
	/* IncidentPreference - Determines how incidents are created for critical violations of the conditions contained in the policy. */
	IncidentPreference AlertsIncidentPreference `json:"incidentPreference"`
	/* Name - Description of the policy. */
	Name string `json:"name"`
}
