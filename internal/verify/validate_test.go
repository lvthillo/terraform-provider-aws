package verify

import (
	"regexp"
	"strings"
	"testing"
)

func TestValidTypeStringNullableBoolean(t *testing.T) {
	testCases := []struct {
		val         interface{}
		expectedErr *regexp.Regexp
	}{
		{
			val: "",
		},
		{
			val: "0",
		},
		{
			val: "1",
		},
		{
			val: "true",
		},
		{
			val: "false",
		},
		{
			val:         "invalid",
			expectedErr: regexp.MustCompile(`to be one of \["", false, true\]`),
		},
	}

	matchErr := func(errs []error, r *regexp.Regexp) bool {
		// err must match one provided
		for _, err := range errs {
			if r.MatchString(err.Error()) {
				return true
			}
		}

		return false
	}

	for i, tc := range testCases {
		_, errs := ValidTypeStringNullableBoolean(tc.val, "test_property")

		if len(errs) == 0 && tc.expectedErr == nil {
			continue
		}

		if len(errs) != 0 && tc.expectedErr == nil {
			t.Fatalf("expected test case %d to produce no errors, got %v", i, errs)
		}

		if !matchErr(errs, tc.expectedErr) {
			t.Fatalf("expected test case %d to produce error matching \"%s\", got %v", i, tc.expectedErr, errs)
		}
	}
}

func TestValidTypeStringNullableFloat(t *testing.T) {
	testCases := []struct {
		val         interface{}
		expectedErr *regexp.Regexp
	}{
		{
			val: "",
		},
		{
			val: "0",
		},
		{
			val: "1",
		},
		{
			val: "42.0",
		},
		{
			val:         "threeve",
			expectedErr: regexp.MustCompile(`cannot parse`),
		},
	}

	matchErr := func(errs []error, r *regexp.Regexp) bool {
		// err must match one provided
		for _, err := range errs {
			if r.MatchString(err.Error()) {
				return true
			}
		}

		return false
	}

	for i, tc := range testCases {
		_, errs := ValidTypeStringNullableFloat(tc.val, "test_property")

		if len(errs) == 0 && tc.expectedErr == nil {
			continue
		}

		if len(errs) != 0 && tc.expectedErr == nil {
			t.Fatalf("expected test case %d to produce no errors, got %v", i, errs)
		}

		if !matchErr(errs, tc.expectedErr) {
			t.Fatalf("expected test case %d to produce error matching \"%s\", got %v", i, tc.expectedErr, errs)
		}
	}
}

func TestValidAccountID(t *testing.T) {
	validNames := []string{
		"123456789012",
		"999999999999",
	}
	for _, v := range validNames {
		_, errors := ValidAccountID(v, "account_id")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid AWS Account ID: %q", v, errors)
		}
	}

	invalidNames := []string{
		"12345678901",   // too short
		"1234567890123", // too long
		"invalid",
		"x123456789012",
	}
	for _, v := range invalidNames {
		_, errors := ValidAccountID(v, "account_id")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid AWS Account ID", v)
		}
	}
}

func TestValidOrganizationID(t *testing.T) {
	validNames := []string{
		"o-aze123rty456",
		"o-123aze456rty780uio",
	}
	for _, v := range validNames {
		_, errors := ValidOrganizationID(v, "organization_id")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid AWS Organization ID: %q", v, errors)
		}
	}

	invalidNames := []string{
		"o-aze123rty", // too short
		"o-123aze456rty780uio123aze456rty780uio123", // too long
		"invalid",
		"123aze456rty780uio",
	}
	for _, v := range invalidNames {
		_, errors := ValidOrganizationID(v, "organization_id")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid AWS Organization ID", v)
		}
	}
}

func TestValidOrganizationalUnitID(t *testing.T) {
	validNames := []string{
		"ou-azer-12aefd983dz",
		"ou-1234-azerty1234567",
	}
	for _, v := range validNames {
		_, errors := ValidOrganizationalUnitID(v, "organizational_unit_id")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid AWS Organizational Unit ID: %q", v, errors)
		}
	}

	invalidNames := []string{
		"ou-aze-12a",               // too short
		"ou-1234aze-azerty1234567", // too long
		"invalid",
		"o-1234-azerty1234567",
	}
	for _, v := range invalidNames {
		_, errors := ValidOrganizationalUnitID(v, "organizational_unit_id")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid AWS Organizational Unit ID", v)
		}
	}
}

func TestValidARN(t *testing.T) {
	v := ""
	_, errors := ValidARN(v, "arn")
	if len(errors) != 0 {
		t.Fatalf("%q should not be validated as an ARN: %q", v, errors)
	}

	validNames := []string{
		"arn:aws:elasticbeanstalk:us-east-1:123456789012:environment/My App/MyEnvironment", // lintignore:AWSAT003,AWSAT005 // Beanstalk
		"arn:aws:iam::123456789012:user/David",                                             // lintignore:AWSAT005          // IAM User
		"arn:aws:iam::aws:policy/CloudWatchReadOnlyAccess",                                 // lintignore:AWSAT005          // Managed IAM policy
		"arn:aws:rds:eu-west-1:123456789012:db:mysql-db",                                   // lintignore:AWSAT003,AWSAT005 // RDS
		"arn:aws:s3:::my_corporate_bucket/exampleobject.png",                               // lintignore:AWSAT005          // S3 object
		"arn:aws:events:us-east-1:319201112229:rule/rule_name",                             // lintignore:AWSAT003,AWSAT005 // CloudWatch Rule
		"arn:aws:lambda:eu-west-1:319201112229:function:myCustomFunction",                  // lintignore:AWSAT003,AWSAT005 // Lambda function
		"arn:aws:lambda:eu-west-1:319201112229:function:myCustomFunction:Qualifier",        // lintignore:AWSAT003,AWSAT005 // Lambda func qualifier
		"arn:aws-cn:ec2:cn-north-1:123456789012:instance/i-12345678",                       // lintignore:AWSAT003,AWSAT005 // China EC2 ARN
		"arn:aws-cn:s3:::bucket/object",                                                    // lintignore:AWSAT005          // China S3 ARN
		"arn:aws-iso:ec2:us-iso-east-1:123456789012:instance/i-12345678",                   // lintignore:AWSAT003,AWSAT005 // C2S EC2 ARN
		"arn:aws-iso:s3:::bucket/object",                                                   // lintignore:AWSAT005          // C2S S3 ARN
		"arn:aws-iso-b:ec2:us-isob-east-1:123456789012:instance/i-12345678",                // lintignore:AWSAT003,AWSAT005 // SC2S EC2 ARN
		"arn:aws-iso-b:s3:::bucket/object",                                                 // lintignore:AWSAT005          // SC2S S3 ARN
		"arn:aws-us-gov:ec2:us-gov-west-1:123456789012:instance/i-12345678",                // lintignore:AWSAT003,AWSAT005 // GovCloud EC2 ARN
		"arn:aws-us-gov:s3:::bucket/object",                                                // lintignore:AWSAT005          // GovCloud S3 ARN
	}
	for _, v := range validNames {
		_, errors := ValidARN(v, "arn")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid ARN: %q", v, errors)
		}
	}

	invalidNames := []string{
		"arn",
		"123456789012",
		"arn:aws",
		"arn:aws:logs",            //lintignore:AWSAT005
		"arn:aws:logs:region:*:*", //lintignore:AWSAT005
	}
	for _, v := range invalidNames {
		_, errors := ValidARN(v, "arn")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid ARN", v)
		}
	}
}

func TestValidateCIDRBlock(t *testing.T) {
	for _, ts := range []struct {
		cidr  string
		valid bool
	}{
		{"10.2.2.0/24", true},
		{"10.2.2.0/1234", false},
		{"10.2.2.2/24", false},
		{"::/0", true},
		{"::0/0", true},
		{"2000::/15", true},
		{"2001::/15", false},
		{"", false},
	} {
		err := validateCIDRBlock(ts.cidr)
		if !ts.valid && err == nil {
			t.Fatalf("Input '%s' should error but didn't!", ts.cidr)
		}
		if ts.valid && err != nil {
			t.Fatalf("Got unexpected error for '%s' input: %s", ts.cidr, err)
		}
	}
}

func TestValidCIDRNetworkAddress(t *testing.T) {
	cases := []struct {
		CIDR              string
		ExpectedErrSubstr string
	}{
		{"notacidr", `is not a valid CIDR block`},
		{"10.0.1.0/16", `is not a valid CIDR block; did you mean`},
		{"10.0.1.0/24", ``},
		{"2001:db8::/122", ``},
		{"2001::/15", `is not a valid CIDR block; did you mean`},
	}

	for i, tc := range cases {
		_, errs := ValidCIDRNetworkAddress(tc.CIDR, "foo")
		if tc.ExpectedErrSubstr == "" {
			if len(errs) != 0 {
				t.Fatalf("%d/%d: Expected no error, got errs: %#v",
					i+1, len(cases), errs)
			}
		} else {
			if len(errs) != 1 {
				t.Fatalf("%d/%d: Expected 1 err containing %q, got %d errs",
					i+1, len(cases), tc.ExpectedErrSubstr, len(errs))
			}
			if !strings.Contains(errs[0].Error(), tc.ExpectedErrSubstr) {
				t.Fatalf("%d/%d: Expected err: %q, to include %q",
					i+1, len(cases), errs[0], tc.ExpectedErrSubstr)
			}
		}
	}
}

func TestValidIPv4CIDRBlock(t *testing.T) {
	for _, ts := range []struct {
		cidr  string
		valid bool
	}{
		{"10.2.2.0/24", true},
		{"10.2.2.0/1234", false},
		{"10/24", false},
		{"10.2.2.2/24", false},
		{"::/0", false},
		{"2000::/15", false},
		{"", false},
	} {
		err := ValidateIPv4CIDRBlock(ts.cidr)
		if !ts.valid && err == nil {
			t.Fatalf("Input '%s' should error but didn't!", ts.cidr)
		}
		if ts.valid && err != nil {
			t.Fatalf("Got unexpected error for '%s' input: %s", ts.cidr, err)
		}
	}
}

func TestValidIPv6CIDRBlock(t *testing.T) {
	for _, ts := range []struct {
		cidr  string
		valid bool
	}{
		{"10.2.2.0/24", false},
		{"10.2.2.0/1234", false},
		{"::/0", true},
		{"::0/0", true},
		{"2000::/15", true},
		{"2001::/15", false},
		{"2001:db8::/122", true},
		{"", false},
	} {
		err := ValidateIPv6CIDRBlock(ts.cidr)
		if !ts.valid && err == nil {
			t.Fatalf("Input '%s' should error but didn't!", ts.cidr)
		}
		if ts.valid && err != nil {
			t.Fatalf("Got unexpected error for '%s' input: %s", ts.cidr, err)
		}
	}
}

func TestValidIAMPolicyJSONString(t *testing.T) {
	type testCases struct {
		Value    string
		ErrCount int
	}

	invalidCases := []testCases{
		{
			Value:    `{0:"1"}`,
			ErrCount: 1,
		},
		{
			Value:    `{'abc':1}`,
			ErrCount: 1,
		},
		{
			Value:    `{"def":}`,
			ErrCount: 1,
		},
		{
			Value:    `{"xyz":[}}`,
			ErrCount: 1,
		},
		{
			Value:    ``,
			ErrCount: 1,
		},
		{
			Value:    `    {"xyz": "foo"}`,
			ErrCount: 1,
		},
	}

	for _, tc := range invalidCases {
		_, errors := ValidIAMPolicyJSON(tc.Value, "json")
		if len(errors) != tc.ErrCount {
			t.Fatalf("Expected %q to trigger a validation error.", tc.Value)
		}
	}

	validCases := []testCases{
		{
			Value:    `{}`,
			ErrCount: 0,
		},
		{
			Value:    `{"abc":["1","2"]}`,
			ErrCount: 0,
		},
	}

	for _, tc := range validCases {
		_, errors := ValidIAMPolicyJSON(tc.Value, "json")
		if len(errors) != tc.ErrCount {
			t.Fatalf("Expected %q not to trigger a validation error.", tc.Value)
		}
	}
}

func TestValidStringIsJSONOrYAML(t *testing.T) {
	type testCases struct {
		Value    string
		ErrCount int
	}

	invalidCases := []testCases{
		{
			Value:    `{"abc":"`,
			ErrCount: 1,
		},
		{
			Value:    "abc: [",
			ErrCount: 1,
		},
	}

	for _, tc := range invalidCases {
		_, errors := ValidStringIsJSONOrYAML(tc.Value, "template")
		if len(errors) != tc.ErrCount {
			t.Fatalf("Expected %q to trigger a validation error.", tc.Value)
		}
	}

	validCases := []testCases{
		{
			Value:    `{"abc":"1"}`,
			ErrCount: 0,
		},
		{
			Value:    `abc: 1`,
			ErrCount: 0,
		},
	}

	for _, tc := range validCases {
		_, errors := ValidStringIsJSONOrYAML(tc.Value, "template")
		if len(errors) != tc.ErrCount {
			t.Fatalf("Expected %q not to trigger a validation error.", tc.Value)
		}
	}
}

func TestValidOnceAWeekWindowFormat(t *testing.T) {
	cases := []struct {
		Value    string
		ErrCount int
	}{
		{
			// once a day window format
			Value:    "04:00-05:00",
			ErrCount: 1,
		},
		{
			// invalid day of week
			Value:    "san:04:00-san:05:00",
			ErrCount: 1,
		},
		{
			// invalid hour
			Value:    "sun:24:00-san:25:00",
			ErrCount: 1,
		},
		{
			// invalid min
			Value:    "sun:04:00-sun:04:60",
			ErrCount: 1,
		},
		{
			// valid format
			Value:    "sun:04:00-sun:05:00",
			ErrCount: 0,
		},
		{
			// "Sun" can also be used
			Value:    "Sun:04:00-Sun:05:00",
			ErrCount: 0,
		},
		{
			// valid format
			Value:    "",
			ErrCount: 0,
		},
	}

	for _, tc := range cases {
		_, errors := ValidOnceAWeekWindowFormat(tc.Value, "maintenance_window")

		if len(errors) != tc.ErrCount {
			t.Fatalf("Expected %d validation errors, But got %d errors for \"%s\"", tc.ErrCount, len(errors), tc.Value)
		}
	}
}

func TestValidOnceADayWindowFormat(t *testing.T) {
	cases := []struct {
		Value    string
		ErrCount int
	}{
		{
			// once a week window format
			Value:    "sun:04:00-sun:05:00",
			ErrCount: 1,
		},
		{
			// invalid hour
			Value:    "24:00-25:00",
			ErrCount: 1,
		},
		{
			// invalid min
			Value:    "04:00-04:60",
			ErrCount: 1,
		},
		{
			// valid format
			Value:    "04:00-05:00",
			ErrCount: 0,
		},
		{
			// valid format
			Value:    "",
			ErrCount: 0,
		},
	}

	for _, tc := range cases {
		_, errors := ValidOnceADayWindowFormat(tc.Value, "backup_window")

		if len(errors) != tc.ErrCount {
			t.Fatalf("Expected %d validation errors, But got %d errors for \"%s\"", tc.ErrCount, len(errors), tc.Value)
		}
	}
}

func TestValidLaunchTemplateName(t *testing.T) {
	validNames := []string{
		"fooBAR123",
		"(./_)",
	}
	for _, v := range validNames {
		_, errors := ValidLaunchTemplateName(v, "name")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid Launch Template name: %q", v, errors)
		}
	}

	invalidNames := []string{
		"tf",
		strings.Repeat("W", 126), // > 125
		"invalid*",
		"invalid\name",
		"inavalid&",
		"invalid+",
		"invalid!",
		"invalid:",
		"invalid;",
	}
	for _, v := range invalidNames {
		_, errors := ValidLaunchTemplateName(v, "name")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid Launch Template name: %q", v, errors)
		}
	}

	invalidNamePrefixes := []string{
		strings.Repeat("W", 100), // > 99
	}
	for _, v := range invalidNamePrefixes {
		_, errors := ValidLaunchTemplateName(v, "name_prefix")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid Launch Template name prefix: %q", v, errors)
		}
	}
}

func TestValidLaunchTemplateID(t *testing.T) {
	validIds := []string{
		"lt-foobar123456",
	}
	for _, v := range validIds {
		_, errors := ValidLaunchTemplateID(v, "id")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid Launch Template id: %q", v, errors)
		}
	}

	invalidIds := []string{
		strings.Repeat("W", 256),
		"invalid-foobar123456",
		"lt_foobar123456",
	}
	for _, v := range invalidIds {
		_, errors := ValidLaunchTemplateID(v, "id")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid Launch Template id: %q", v, errors)
		}
	}
}

func TestValidUTCTimestamp(t *testing.T) {
	validT := []string{
		"2006-01-02T15:04:05Z",
	}

	invalidT := []string{
		"2015-03-07 23:45:00",
		"27-03-2019 23:45:00",
		"Mon, 02 Jan 2006 15:04:05 -0700",
	}

	for _, f := range validT {
		_, errors := ValidUTCTimestamp(f, "utc_timestamp")
		if len(errors) > 0 {
			t.Fatalf("expected the time %q to be in valid format, got error %q", f, errors)
		}
	}

	for _, f := range invalidT {
		_, errors := ValidUTCTimestamp(f, "utc_timestamp")
		if len(errors) == 0 {
			t.Fatalf("expected the time %q to fail validation", f)
		}
	}
}

func TestValidateTypeStringIsDateOrInt(t *testing.T) {
	validT := []string{
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05-07:00",
		"1234",
		"0",
	}

	for _, f := range validT {
		_, errors := ValidStringDateOrPositiveInt(f, "parameter")
		if len(errors) > 0 {
			t.Fatalf("expected the value %q to be either RFC 3339 or positive integer, got error %q", f, errors)
		}
	}

	invalidT := []string{
		"2018-03-01T00:00:00", // No time zone
		"ABC",
		"-789",
	}

	for _, f := range invalidT {
		_, errors := ValidStringDateOrPositiveInt(f, "parameter")
		if len(errors) == 0 {
			t.Fatalf("expected the value %q to fail validation", f)
		}
	}
}
