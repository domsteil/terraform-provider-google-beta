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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeVpnTunnel_vpnTunnelBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeVpnTunnelDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeVpnTunnel_vpnTunnelBasicExample(context),
			},
			{
				ResourceName:            "google_compute_vpn_tunnel.tunnel1",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared_secret"},
			},
		},
	})
}

func testAccComputeVpnTunnel_vpnTunnelBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_vpn_tunnel" "tunnel1" {
  name          = "tunnel1%{random_suffix}"
  peer_ip       = "15.0.0.120"
  shared_secret = "a secret message"

  target_vpn_gateway = google_compute_vpn_gateway.target_gateway.id

  depends_on = [
    google_compute_forwarding_rule.fr_esp,
    google_compute_forwarding_rule.fr_udp500,
    google_compute_forwarding_rule.fr_udp4500,
  ]
}

resource "google_compute_vpn_gateway" "target_gateway" {
  name    = "vpn1%{random_suffix}"
  network = google_compute_network.network1.id
}

resource "google_compute_network" "network1" {
  name = "network1%{random_suffix}"
}

resource "google_compute_address" "vpn_static_ip" {
  name = "tf-test-vpn-static-ip%{random_suffix}"
}

resource "google_compute_forwarding_rule" "fr_esp" {
  name        = "tf-test-fr-esp%{random_suffix}"
  ip_protocol = "ESP"
  ip_address  = google_compute_address.vpn_static_ip.address
  target      = google_compute_vpn_gateway.target_gateway.id
}

resource "google_compute_forwarding_rule" "fr_udp500" {
  name        = "tf-test-fr-udp500%{random_suffix}"
  ip_protocol = "UDP"
  port_range  = "500"
  ip_address  = google_compute_address.vpn_static_ip.address
  target      = google_compute_vpn_gateway.target_gateway.id
}

resource "google_compute_forwarding_rule" "fr_udp4500" {
  name        = "tf-test-fr-udp4500%{random_suffix}"
  ip_protocol = "UDP"
  port_range  = "4500"
  ip_address  = google_compute_address.vpn_static_ip.address
  target      = google_compute_vpn_gateway.target_gateway.id
}

resource "google_compute_route" "route1" {
  name       = "route1%{random_suffix}"
  network    = google_compute_network.network1.name
  dest_range = "15.0.0.0/24"
  priority   = 1000

  next_hop_vpn_tunnel = google_compute_vpn_tunnel.tunnel1.id
}
`, context)
}

func TestAccComputeVpnTunnel_vpnTunnelBetaExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeVpnTunnelDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeVpnTunnel_vpnTunnelBetaExample(context),
			},
		},
	})
}

func testAccComputeVpnTunnel_vpnTunnelBetaExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_vpn_tunnel" "tunnel1" {
  provider      = google-beta
  name          = "tunnel1%{random_suffix}"
  peer_ip       = "15.0.0.120"
  shared_secret = "a secret message"

  target_vpn_gateway = google_compute_vpn_gateway.target_gateway.id

  depends_on = [
    google_compute_forwarding_rule.fr_esp,
    google_compute_forwarding_rule.fr_udp500,
    google_compute_forwarding_rule.fr_udp4500,
  ]

  labels = {
    foo = "bar"
  }
}

resource "google_compute_vpn_gateway" "target_gateway" {
  provider = google-beta
  name     = "vpn1%{random_suffix}"
  network  = google_compute_network.network1.id
}

resource "google_compute_network" "network1" {
  provider = google-beta
  name     = "network1%{random_suffix}"
}

resource "google_compute_address" "vpn_static_ip" {
  provider = google-beta
  name     = "tf-test-vpn-static-ip%{random_suffix}"
}

resource "google_compute_forwarding_rule" "fr_esp" {
  provider    = google-beta
  name        = "tf-test-fr-esp%{random_suffix}"
  ip_protocol = "ESP"
  ip_address  = google_compute_address.vpn_static_ip.address
  target      = google_compute_vpn_gateway.target_gateway.id
}

resource "google_compute_forwarding_rule" "fr_udp500" {
  provider    = google-beta
  name        = "tf-test-fr-udp500%{random_suffix}"
  ip_protocol = "UDP"
  port_range  = "500"
  ip_address  = google_compute_address.vpn_static_ip.address
  target      = google_compute_vpn_gateway.target_gateway.id
}

resource "google_compute_forwarding_rule" "fr_udp4500" {
  provider    = google-beta
  name        = "tf-test-fr-udp4500%{random_suffix}"
  ip_protocol = "UDP"
  port_range  = "4500"
  ip_address  = google_compute_address.vpn_static_ip.address
  target      = google_compute_vpn_gateway.target_gateway.id
}

resource "google_compute_route" "route1" {
  provider   = google-beta
  name       = "route1%{random_suffix}"
  network    = google_compute_network.network1.name
  dest_range = "15.0.0.0/24"
  priority   = 1000

  next_hop_vpn_tunnel = google_compute_vpn_tunnel.tunnel1.id
}

provider "google-beta" {
  region = "us-central1"
  zone   = "us-central1-a"
}
`, context)
}

func testAccCheckComputeVpnTunnelDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_vpn_tunnel" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("ComputeVpnTunnel still exists at %s", url)
			}
		}

		return nil
	}
}
