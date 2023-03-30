package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccTagBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckRequiredEnvVars(t) },
		Providers: testAccProviders,
		CheckDestroy: func(s *terraform.State) error {
			return testAccCheckDestroy(s, "contrail_tag", "tag")
		},
		Steps: []resource.TestStep{
			{
				Config: testAccTag,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExists("contrail_tag.tag_test", "tag"),
				),
			},
		},
	})
}

const testAccTag = `
resource "contrail_tag" "tag_test" {
    name = "test_tag"
    tag_value = "value"
    tag_type_name = "type"
	display_name = "type=value"
}`
