// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccMonitoringNotificationChannel_notificationChannelBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitoringNotificationChannelDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringNotificationChannel_notificationChannelBasicExample(context),
			},
			{
				ResourceName:            "google_monitoring_notification_channel.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sensitive_labels"},
			},
		},
	})
}

func testAccMonitoringNotificationChannel_notificationChannelBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_monitoring_notification_channel" "basic" {
  display_name = "Test Notification Channel%{random_suffix}"
  type         = "email"
  labels = {
    email_address = "fake_email@blahblah.com"
  }
}
`, context)
}

func testAccCheckMonitoringNotificationChannelDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_monitoring_notification_channel" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{MonitoringBasePath}}v3/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil, isMonitoringConcurrentEditError)
			if err == nil {
				return fmt.Errorf("MonitoringNotificationChannel still exists at %s", url)
			}
		}

		return nil
	}
}
