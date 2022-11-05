package notifier

import (
	"os"
	"testing"
)

func Test_getNTFYTimeoutMillisEnvVariable(t *testing.T) {
	expectedTimeoutMillis := 10000

	os.Setenv(ntfyTimeoutMillisEnvVariable, "10000")

	actualTimeoutMillis := getNTFYTimeoutMillisEnvVariable()

	if expectedTimeoutMillis != actualTimeoutMillis {
		t.Errorf("Timeout was incorrect want: %+v, but got: %+v", expectedTimeoutMillis, actualTimeoutMillis)
	}
	os.Unsetenv(ntfyTimeoutMillisEnvVariable)
}

func Test_getNTFYTimeoutMillisEnvVariable_negativeValue(t *testing.T) {
	expectedTimeoutMillis := 5000

	os.Setenv(ntfyTimeoutMillisEnvVariable, "-3000")

	actualTimeoutMillis := getNTFYTimeoutMillisEnvVariable()

	if expectedTimeoutMillis != actualTimeoutMillis {
		t.Errorf("Timeout was incorrect want: %+v, but got: %+v", expectedTimeoutMillis, actualTimeoutMillis)
	}
	os.Unsetenv(ntfyTimeoutMillisEnvVariable)
}

func Test_getNTFYTimeoutMillisEnvVariable_badValue(t *testing.T) {
	expectedTimeoutMillis := 5000

	os.Setenv(ntfyTimeoutMillisEnvVariable, "bad")

	actualTimeoutMillis := getNTFYTimeoutMillisEnvVariable()

	if expectedTimeoutMillis != actualTimeoutMillis {
		t.Errorf("Timeout was incorrect want: %+v, but got: %+v", expectedTimeoutMillis, actualTimeoutMillis)
	}
	os.Unsetenv(ntfyTimeoutMillisEnvVariable)
}

func Test_getNTFYTimeoutMillisEnvVariable_noValue(t *testing.T) {
	expectedTimeoutMillis := 5000

	actualTimeoutMillis := getNTFYTimeoutMillisEnvVariable()

	if expectedTimeoutMillis != actualTimeoutMillis {
		t.Errorf("Timeout was incorrect want: %+v, but got: %+v", expectedTimeoutMillis, actualTimeoutMillis)
	}
}

func Test_getNTFYDefaultPriorityEnvVariable(t *testing.T) {
	expectedDefaultPriority := 2

	os.Setenv(ntfyDefaultPriorityEnvVariable, "2")

	actualDefaultPriority := getNTFYDefaultPriorityEnvVariable()

	if expectedDefaultPriority != actualDefaultPriority {
		t.Errorf("Default priority was incorrect want: %+v, but got: %+v", expectedDefaultPriority, actualDefaultPriority)
	}
	os.Unsetenv(ntfyDefaultPriorityEnvVariable)
}

func Test_getNTFYDefaultPriorityEnvVariable_0(t *testing.T) {
	expectedDefaultPriority := 3

	os.Setenv(ntfyDefaultPriorityEnvVariable, "0")

	actualDefaultPriority := getNTFYDefaultPriorityEnvVariable()

	if expectedDefaultPriority != actualDefaultPriority {
		t.Errorf("Default priority was incorrect want: %+v, but got: %+v", expectedDefaultPriority, actualDefaultPriority)
	}
	os.Unsetenv(ntfyDefaultPriorityEnvVariable)
}

func Test_getNTFYDefaultPriorityEnvVariable_6(t *testing.T) {
	expectedDefaultPriority := 3

	os.Setenv(ntfyDefaultPriorityEnvVariable, "6")

	actualDefaultPriority := getNTFYDefaultPriorityEnvVariable()

	if expectedDefaultPriority != actualDefaultPriority {
		t.Errorf("Default priority was incorrect want: %+v, but got: %+v", expectedDefaultPriority, actualDefaultPriority)
	}
	os.Unsetenv(ntfyDefaultPriorityEnvVariable)
}

func Test_getNTFYDefaultPriorityEnvVariable_badValue(t *testing.T) {
	expectedDefaultPriority := 3

	os.Setenv(ntfyDefaultPriorityEnvVariable, "bad")

	actualDefaultPriority := getNTFYDefaultPriorityEnvVariable()

	if expectedDefaultPriority != actualDefaultPriority {
		t.Errorf("Default priority was incorrect want: %+v, but got: %+v", expectedDefaultPriority, actualDefaultPriority)
	}
	os.Unsetenv(ntfyDefaultPriorityEnvVariable)
}

func Test_getNTFYDefaultPriorityEnvVariable_noValue(t *testing.T) {
	expectedDefaultPriority := 3

	actualDefaultPriority := getNTFYDefaultPriorityEnvVariable()

	if expectedDefaultPriority != actualDefaultPriority {
		t.Errorf("Default priority was incorrect want: %+v, but got: %+v", expectedDefaultPriority, actualDefaultPriority)
	}
}
