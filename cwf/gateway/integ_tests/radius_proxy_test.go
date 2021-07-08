// +build all authenticate

/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"fmt"
	cwfprotos "magma/cwf/cloud/go/protos"
	"magma/feg/cloud/go/protos"
	lteProtos "magma/lte/cloud/go/protos"
	"magma/lte/cloud/go/services/policydb/obsidian/models"
	"math"
	"testing"
	"time"

	"github.com/fiorix/go-diameter/v4/diam"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
)

// - Initialize 3 UEs and initiate Authentication. Assert that it is successful.
// - Disconnect all UEs.
func TestRadiusProxy(t *testing.T) {
	fmt.Println("\nRunning TestRadiusProxy...")
	tr := NewTestRunner(t)
	defer func() {
		// Clear hss, ocs, and pcrf
		assert.NoError(t, tr.CleanUp())
	}()

	ruleManager, err := NewRuleManager()
	assert.NoError(t, err)
	assert.NoError(t, usePCRFMockDriver())
	defer func() {
		// Clear hss, ocs, and pcrf
		assert.NoError(t, clearPCRFMockDriver())
		assert.NoError(t, ruleManager.RemoveInstalledRules())
		assert.NoError(t, tr.CleanUp())
	}()

	ues, err := tr.ConfigUEs(1)
	assert.NoError(t, err)
	imsi := ues[0].GetImsi()

	err = ruleManager.AddStaticPassAllToDB("usage-enforcement-static-pass-all", "mkey1", 0, models.PolicyRuleTrackingTypeONLYPCRF, 3)
	assert.NoError(t, err)
	tr.WaitForPoliciesToSync()

	usageMonitorInfo := getUsageInformation("mkey1", 1*MegaBytes)

	initRequest := protos.NewGxCCRequest(imsi, protos.CCRequestType_INITIAL)
	initAnswer := protos.NewGxCCAnswer(diam.Success).
		SetStaticRuleInstalls([]string{"usage-enforcement-static-pass-all"}, []string{}).
		SetUsageMonitorInfo(usageMonitorInfo)
	initExpectation := protos.NewGxCreditControlExpectation().Expect(initRequest).Return(initAnswer)

	// We expect an update request with some usage update (probably around 80-100% of the given quota)
	updateRequest1 := protos.NewGxCCRequest(imsi, protos.CCRequestType_UPDATE).
		SetUsageMonitorReport(usageMonitorInfo).
		SetUsageReportDelta(uint64(math.Round(0.2 * 1 * MegaBytes))).
		SetEventTrigger(int32(lteProtos.EventTrigger_USAGE_REPORT))
	updateAnswer1 := protos.NewGxCCAnswer(diam.Success).SetUsageMonitorInfo(usageMonitorInfo)
	updateExpectation1 := protos.NewGxCreditControlExpectation().Expect(updateRequest1).Return(updateAnswer1)
	expectations := []*protos.GxCreditControlExpectation{initExpectation, updateExpectation1}
	// On unexpected requests, just return the default update answer
	assert.NoError(t, setPCRFExpectations(expectations, updateAnswer1))

	tr.ProxyRadiusAndAssertSuccess(imsi)

	time.Sleep(2 * time.Second)

	// First wait until we see the original static-pass-all-ocs2 show up
	assert.Eventually(t,
		tr.WaitForEnforcementStatsForRule(imsi, "usage-enforcement-static-pass-all"), time.Minute, 2*time.Second)
	fmt.Println("CCR-I exchanged installed usage-enforcement-static-pass-all")

	req := &cwfprotos.GenTrafficRequest{Imsi: imsi, Volume: &wrappers.StringValue{Value: "900K"}}
	_, err = tr.GenULTraffic(req)
	assert.NoError(t, err)
	tr.WaitForEnforcementStatsToSync()
	req = &cwfprotos.GenTrafficRequest{Imsi: imsi, Volume: &wrappers.StringValue{Value: "200K"}}
	_, err = tr.GenULTraffic(req)
	assert.NoError(t, err)
	tr.WaitForEnforcementStatsToSync()

	// Assert that enforcement_stats rules are properly installed and the right
	// amount of data was passed through
	tr.AssertPolicyUsage(imsi, "usage-enforcement-static-pass-all", 1, uint64(math.Round(1.2*MegaBytes+Buffer)))

	// Assert that a CCR-I and at least one CCR-U were sent up to the PCRF
	tr.AssertAllGxExpectationsMetNoError()

	// When we initiate a UE disconnect, we expect a terminate request to go up
	terminateRequest := protos.NewGxCCRequest(imsi, protos.CCRequestType_TERMINATION)
	terminateAnswer := protos.NewGxCCAnswer(diam.Success)
	terminateExpectation := protos.NewGxCreditControlExpectation().Expect(terminateRequest).Return(terminateAnswer)
	expectations = []*protos.GxCreditControlExpectation{terminateExpectation}
	assert.NoError(t, setPCRFExpectations(expectations, nil))

	tr.DisconnectAndAssertSuccess(imsi)

}

/*
func TestGxUsageReportEnforcement(t *testing.T) {
	fmt.Println("\nRunning EndtoEnd-NON-EAP-Auth...")
	tr := NewTestRunner(t)
	ruleManager, err := NewRuleManager()
	assert.NoError(t, err)
	assert.NoError(t, usePCRFMockDriver())
	defer func() {
		// Clear hss, ocs, and pcrf
		assert.NoError(t, clearPCRFMockDriver())
		assert.NoError(t, ruleManager.RemoveInstalledRules())
		assert.NoError(t, tr.CleanUp())
	}()

	ues, err := tr.ConfigUEs(1)
	assert.NoError(t, err)
	imsi := ues[0].GetImsi()

	err = ruleManager.AddStaticPassAllToDB("usage-enforcement-static-pass-all", "mkey1", 0, models.PolicyRuleTrackingTypeONLYPCRF, 3)
	assert.NoError(t, err)
	tr.WaitForPoliciesToSync()

	usageMonitorInfo := getUsageInformation("mkey1", 1*MegaBytes)

	initRequest := protos.NewGxCCRequest(imsi, protos.CCRequestType_INITIAL)
	initAnswer := protos.NewGxCCAnswer(diam.Success).
		SetStaticRuleInstalls([]string{"usage-enforcement-static-pass-all"}, []string{}).
		SetUsageMonitorInfo(usageMonitorInfo)
	initExpectation := protos.NewGxCreditControlExpectation().Expect(initRequest).Return(initAnswer)

	// We expect an update request with some usage update (probably around 80-100% of the given quota)
	updateRequest1 := protos.NewGxCCRequest(imsi, protos.CCRequestType_UPDATE).
		SetUsageMonitorReport(usageMonitorInfo).
		SetUsageReportDelta(uint64(math.Round(0.2 * 1 * MegaBytes))).
		SetEventTrigger(int32(lteProtos.EventTrigger_USAGE_REPORT))
	updateAnswer1 := protos.NewGxCCAnswer(diam.Success).SetUsageMonitorInfo(usageMonitorInfo)
	updateExpectation1 := protos.NewGxCreditControlExpectation().Expect(updateRequest1).Return(updateAnswer1)
	expectations := []*protos.GxCreditControlExpectation{initExpectation, updateExpectation1}
	// On unexpected requests, just return the default update answer
	assert.NoError(t, setPCRFExpectations(expectations, updateAnswer1))

	tr.AuthenticateAndAssertSuccess(imsi)
	// First wait until we see the original static-pass-all-ocs2 show up
	assert.Eventually(t,
		tr.WaitForEnforcementStatsForRule(imsi, "usage-enforcement-static-pass-all"), time.Minute, 2*time.Second)
	fmt.Println("CCR-I exchanged installed usage-enforcement-static-pass-all")

	req := &cwfprotos.GenTrafficRequest{Imsi: imsi, Volume: &wrappers.StringValue{Value: "900K"}}
	_, err = tr.GenULTraffic(req)
	assert.NoError(t, err)
	tr.WaitForEnforcementStatsToSync()
	req = &cwfprotos.GenTrafficRequest{Imsi: imsi, Volume: &wrappers.StringValue{Value: "200K"}}
	_, err = tr.GenULTraffic(req)
	assert.NoError(t, err)
	tr.WaitForEnforcementStatsToSync()

	// Assert that enforcement_stats rules are properly installed and the right
	// amount of data was passed through
	tr.AssertPolicyUsage(imsi, "usage-enforcement-static-pass-all", 1, uint64(math.Round(1.2*MegaBytes+Buffer)))

	// Assert that a CCR-I and at least one CCR-U were sent up to the PCRF
	tr.AssertAllGxExpectationsMetNoError()

	// When we initiate a UE disconnect, we expect a terminate request to go up
	terminateRequest := protos.NewGxCCRequest(imsi, protos.CCRequestType_TERMINATION)
	terminateAnswer := protos.NewGxCCAnswer(diam.Success)
	terminateExpectation := protos.NewGxCreditControlExpectation().Expect(terminateRequest).Return(terminateAnswer)
	expectations = []*protos.GxCreditControlExpectation{terminateExpectation}
	assert.NoError(t, setPCRFExpectations(expectations, nil))

	tr.DisconnectAndAssertSuccess(imsi)
	tr.AssertEventuallyAllRulesRemovedAfterDisconnect(imsi)
	// Assert that we saw a Terminate request
	tr.AssertAllGxExpectationsMetNoError()
} */
